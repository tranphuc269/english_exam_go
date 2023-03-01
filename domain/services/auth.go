package services

import (
	"context"
	dtos "english_exam_go/domain/dtos/user"
	"english_exam_go/infrastructure/data/repositories/persistence"
	auth_utils "english_exam_go/utils/auth"
)

type IAuthService interface {
	Login(ctx context.Context, request dtos.LoginRequest) (*dtos.AuthResponse, error)
	Register(context.Context, dtos.RegisterAccountRequest) (*dtos.AuthResponse, error)
}

type AuthServiceImpl struct {
	ar persistence.IAuthRepository
}

func (as AuthServiceImpl) Register(ctx context.Context, request dtos.RegisterAccountRequest) (*dtos.AuthResponse, error) {
	//TODO implement me
	userEnt, err := request.RegisterAccountToUserEnt()
	if err != nil {
		return nil, err
	}
	userEnt, err = as.ar.CreateUser(ctx, userEnt)
	if err != nil {
		return nil, err
	}
	jwt, err := auth_utils.GenerateJWT(userEnt.Email, userEnt.Name)
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
	//TODO implement me
	jwt, err := auth_utils.GenerateJWT(request.Email, "Tran Van Phuc")
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
