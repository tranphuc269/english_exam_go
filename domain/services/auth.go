package services

import (
	"context"
	dtos "english_exam_go/domain/dtos/user"
	auth_utils "english_exam_go/utils/auth"
	"fmt"
)

type IAuthService interface {
	Login(ctx context.Context, request dtos.LoginRequest) (string, error)
}

type AuthServiceImpl struct {
}

func (a AuthServiceImpl) Login(ctx context.Context, request dtos.LoginRequest) (string, error) {
	//TODO implement me
	jwt, err := auth_utils.GenerateJWT(request.Email, "Tran Van Phuc")
	return jwt, err

}

func CreateAuthService() IAuthService {
	fmt.Println("Create Auth Service")
	return &AuthServiceImpl{}
}
