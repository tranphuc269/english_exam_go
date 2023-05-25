package services

import (
	"context"
	dtos "english_exam_go/domain/dtos/user"
	"english_exam_go/infrastructure/data/repositories/persistence"
	auth_utils "english_exam_go/utils/auth"
	"fmt"
	"time"
)

type IAuthService interface {
	Login(ctx context.Context, request dtos.LoginRequest) (*dtos.AuthResponse, error)
	Register(context.Context, dtos.RegisterAccountRequest) (*dtos.AuthResponse, error)
	Me(context.Context, string) (*dtos.UserResponse, error)
	FindById(context.Context, int) (*dtos.UserResponse, error)
	Update(context.Context, dtos.UpdateAccountRequest, string) error
	Teachers(context.Context, int, int, string, string) []*dtos.UserResponse
	Students(context.Context, int, int, string) []*dtos.UserResponse
}

type AuthServiceImpl struct {
	ar persistence.IAuthRepository
	//fs FileSer
}

func (as AuthServiceImpl) FindById(ctx context.Context, ID int) (*dtos.UserResponse, error) {
	//TODO implement me
	userEnt, _ := as.ar.FindById(ctx, ID)
	return dtos.UserEntToResponse(userEnt), nil
}

func (as AuthServiceImpl) Update(ctx context.Context, request dtos.UpdateAccountRequest, email string) error {
	//TODO implement me
	currentUser, _ := as.ar.FindUserByEmail(ctx, email)
	if request.Name != "" {
		currentUser.Name = request.Name
	}
	if request.Password != "" {
		hashPassword, _ := auth_utils.HashPassword(request.Password)
		currentUser.Password = hashPassword
	}

	if request.DateOfBirth != "" {
		date, _ := time.Parse("2006-01-02", request.DateOfBirth)
		currentUser.DateOfBirth = date
	}

	if request.Avatar != "" {
		currentUser.Avatar = request.Avatar
	}
	if request.Code != "" {
		currentUser.Code = request.Code
	}
	if request.PhoneNumber != "" {
		currentUser.PhoneNumber = request.PhoneNumber
	}
	if request.Address != "" {
		currentUser.Address = request.Address
	}
	err := as.ar.UpdateUser(ctx, currentUser)
	if err != nil {
		return err
	}
	return nil
}

func (as AuthServiceImpl) Teachers(ctx context.Context, offset int, limit int, name string, code string) []*dtos.UserResponse {
	//TODO implement me
	var result []*dtos.UserResponse
	userEnts := as.ar.GetUsers(ctx, 3, offset, limit, name, code)
	for _, ent := range userEnts {
		result = append(result, dtos.UserEntToResponse(ent))
	}
	return result
}

func (as AuthServiceImpl) Students(ctx context.Context, offset int, limit int, code string) []*dtos.UserResponse {
	//TODO implement me
	var result []*dtos.UserResponse
	userEnts := as.ar.GetUsers(ctx, 2, offset, limit, "", code)
	for _, ent := range userEnts {
		result = append(result, dtos.UserEntToResponse(ent))
	}
	return result
}

func (as AuthServiceImpl) Me(ctx context.Context, email string) (*dtos.UserResponse, error) {
	//TODO implement me
	userEnt, _ := as.ar.FindUserByEmail(ctx, email)
	return dtos.UserEntToResponse(userEnt), nil
}

func (as AuthServiceImpl) Register(ctx context.Context, request dtos.RegisterAccountRequest) (*dtos.AuthResponse, error) {
	//TODO implement me
	//_, err := as.Me(ctx, request.Email)
	//if err != nil {
	//	return nil, err
	//}
	userEnt, err := request.RegisterAccountToUserEnt()
	if err != nil {
		return nil, err
	}
	userEnt, err = as.ar.CreateUser(ctx, userEnt)
	if err != nil {
		return nil, err
	}
	jwt, err := auth_utils.GenerateJWT(int(userEnt.ID), userEnt.Email, userEnt.Name, request.Role.ToString())
	if err != nil {
		return nil, err
	}
	var authRes dtos.AuthResponse
	authRes = dtos.AuthResponse{
		Id:          userEnt.ID,
		AccessToken: jwt,
		ExpiresIn:   1 * 60 * 60,
	}
	return &authRes, err
}

func (as AuthServiceImpl) Login(ctx context.Context, request dtos.LoginRequest) (*dtos.AuthResponse, error) {
	userEnt, err := as.ar.FindUserByEmail(ctx, request.Email)
	err = userEnt.CheckPassword(request.Password)
	if err != nil {
		fmt.Println("Mật khẩu không chính xác")
	}
	if err != nil {
		return nil, err
	}

	jwt, err := auth_utils.GenerateJWT(int(userEnt.ID), request.Email, userEnt.Name, userEnt.Role.ToString())
	if err != nil {
		return nil, err
	}
	var authRes dtos.AuthResponse
	authRes = dtos.AuthResponse{
		Id:          5,
		AccessToken: jwt,
		ExpiresIn:   1 * 60 * 60,
	}
	return &authRes, err

}

func CreateAuthService(ar persistence.IAuthRepository) IAuthService {
	return &AuthServiceImpl{
		ar: ar,
	}
}
