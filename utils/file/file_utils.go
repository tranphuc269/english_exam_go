package file

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"time"
)

type Form struct {
	Files []*multipart.FileHeader `form:"files" binding:"required"`
}

const (
	maxPartSize = int64(5 * 1024 * 1024)
	maxRetries  = 3
)

var (
	awsAccessKeyID     = "AKIATHGWX7MGQITVVMUP"
	awsSecretAccessKey = "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"
	awsBucketRegion    = "ap-southeast-1"
	awsBucketName      = "kita"
)

func UploadFile(c *gin.Context) {
	//TODO implement me
	var form Form
	_ = c.ShouldBind(&form)
	// Validate inputs
	successPaths := UploadFileToS3(form.Files)
	c.JSON(http.StatusOK, gin.H{
		"paths": successPaths,
	})
}

func UploadFileToS3(files []*multipart.FileHeader) []string {
	valid, message := validateUploadFiles(files)

	if !valid {
		fmt.Println(message)
		return []string{}
	}

	// get credentials
	creds := credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, "")
	_, err := creds.Get()
	if err != nil {
		fmt.Printf("bad credentials: %s", err)
		// Response to client
		return []string{}
	}

	cfg := aws.NewConfig().WithRegion(awsBucketRegion).WithCredentials(creds)
	ss, err := session.NewSession(cfg)

	if err != nil {
		fmt.Printf("Failed creating new session: %s", err)
		// Response to client
		return []string{}
	}

	svc := s3.New(ss, cfg)

	now := time.Now()
	nowRFC3339 := now.Format(time.RFC3339)

	successPaths := make([]string, len(files))
	pathNumber := 0
	for _, formFile := range files {
		binaryFile, err := readFile(formFile)

		if err != nil {
			// Response to client
			return []string{}
		}

		path := "/media/" + nowRFC3339 + "-" + formFile.Filename
		fileType := http.DetectContentType(binaryFile)

		input := &s3.CreateMultipartUploadInput{
			Bucket:      aws.String(awsBucketName),
			Key:         aws.String(path),
			ContentType: aws.String(fileType),
			ACL:         aws.String("public-read"),
		}

		resp, err := svc.CreateMultipartUpload(input)
		if err != nil {
			fmt.Println("Error " + err.Error())
			// Response to client
			return []string{}
		}
		fmt.Println("Created multipart upload request")

		var curr, partLength int64
		var remaining = formFile.Size
		var completedParts []*s3.CompletedPart
		partNumber := 1
		for curr = 0; remaining != 0; curr += partLength {
			if remaining < maxPartSize {
				partLength = remaining
			} else {
				partLength = maxPartSize
			}
			// Upload binaries part
			completedPart, _ := uploadPart(svc, resp, binaryFile[curr:curr+partLength], partNumber)

			// If upload this part fail
			// Make an abort upload error and exit
			if err != nil {
				fmt.Println(err.Error())
				err := abortMultipartUpload(svc, resp)
				if err != nil {
					fmt.Println(err.Error())
				}
				// Response to client
				return []string{}
			}
			// else append completed part to a whole
			remaining -= partLength
			partNumber++
			completedParts = append(completedParts, completedPart)
		}

		completeResponse, err := completeMultipartUpload(svc, resp, completedParts)
		fmt.Println(completeResponse)
		if err != nil {
			fmt.Println(err.Error())
			// Response to client
			return []string{}
		}

		fmt.Printf("Successfully uploaded file: %s\n", completeResponse.String())

		// Dereference the pointer
		successPaths[pathNumber] = *completeResponse.Location
		pathNumber++

		// Save to disk
		//`formFile` has io.reader method
		//err := c.SaveUploadedFile(formFile, formFile.Filename)
		//if err != nil {
		//	log.Fatalf("Failed saving file %v to disk", formFile.Filename)
		//}
	}
	return successPaths
}

func validateUploadFiles(files []*multipart.FileHeader) (bool, string) {
	for _, formFile := range files {
		size := formFile.Size
		//contentType := formFile.Header.Get("Content-Type")

		if size > maxPartSize {
			return false, "File too large"
		}

		//if contentType != "image/jpeg" && contentType != "image/png" {
		//	return false, "Filetype is not supported"
		//}
	}
	return true, "ok"
}

func readFile(file *multipart.FileHeader) ([]byte, error) {
	// Get raw file bytes - no reader method
	openedFile, _ := file.Open()

	binaryFile, err := ioutil.ReadAll(openedFile)

	if err != nil {
		return nil, err
	}

	defer func(openedFile multipart.File) {
		err := openedFile.Close()
		if err != nil {
			log.Fatalf("Failed closing file %v", file.Filename)
		}
	}(openedFile)
	return binaryFile, nil
}

func completeMultipartUpload(svc *s3.S3, resp *s3.CreateMultipartUploadOutput, completedParts []*s3.CompletedPart) (*s3.CompleteMultipartUploadOutput, error) {
	completeInput := &s3.CompleteMultipartUploadInput{
		Bucket:   resp.Bucket,
		Key:      resp.Key,
		UploadId: resp.UploadId,
		MultipartUpload: &s3.CompletedMultipartUpload{
			Parts: completedParts,
		},
	}
	return svc.CompleteMultipartUpload(completeInput)
}

func uploadPart(svc *s3.S3, resp *s3.CreateMultipartUploadOutput, fileBytes []byte, partNumber int) (*s3.CompletedPart, error) {
	tryNum := 1
	partInput := &s3.UploadPartInput{
		Body:          bytes.NewReader(fileBytes),
		Bucket:        resp.Bucket,
		Key:           resp.Key,
		PartNumber:    aws.Int64(int64(partNumber)),
		UploadId:      resp.UploadId,
		ContentLength: aws.Int64(int64(len(fileBytes))),
	}
	for tryNum <= maxRetries {
		uploadResult, err := svc.UploadPart(partInput)
		if err != nil {
			if tryNum == maxRetries {
				if aerr, ok := err.(awserr.Error); ok {
					return nil, aerr
				}
				return nil, err
			}
			fmt.Printf("Retrying to upload part #%v\n", partNumber)
			tryNum++
		} else {
			fmt.Printf("Uploaded part #%v\n", partNumber)
			return &s3.CompletedPart{
				ETag:       uploadResult.ETag,
				PartNumber: aws.Int64(int64(partNumber)),
			}, nil
		}
	}
	return nil, nil
}

func abortMultipartUpload(svc *s3.S3, resp *s3.CreateMultipartUploadOutput) error {
	fmt.Println("Aborting multipart upload for UploadId#" + *resp.UploadId)
	abortInput := &s3.AbortMultipartUploadInput{
		Bucket:   resp.Bucket,
		Key:      resp.Key,
		UploadId: resp.UploadId,
	}
	_, err := svc.AbortMultipartUpload(abortInput)
	return err
}
