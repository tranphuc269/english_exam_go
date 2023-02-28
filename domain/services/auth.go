package services

import (
	"context"
	dtos "english_exam_go/domain/dtos/user"
)

type IAuthService interface {
	Login(ctx context.Context, request dtos.LoginRequest) (string, error)
}

type AuthServiceImpl struct {
}

func (a AuthServiceImpl) Login(ctx context.Context, request dtos.LoginRequest) (string, error) {
	//TODO implement me
	panic("implement me")
}

func CreateAuthService() IAuthService {
	return &AuthServiceImpl{}
}
