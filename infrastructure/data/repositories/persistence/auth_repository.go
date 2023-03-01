package persistence

import (
	"context"
	"english_exam_go/infrastructure/data/entities"
	"english_exam_go/infrastructure/data/repositories"
	"fmt"
)

type IAuthRepository interface {
	CreateUser(context.Context, *entities.UserEnt) (*entities.UserEnt, error)
	FindUserById(context.Context, uint) (*entities.UserEnt, error)
}

type AuthRepositoryImpl struct {
}

func (ar AuthRepositoryImpl) CreateUser(ctx context.Context, ent *entities.UserEnt) (*entities.UserEnt, error) {
	db := repositories.GetConn()

	result := db.Create(ent)
	if result.Error != nil {
		return nil, &repositories.RdbRuntimeError{
			ErrMsg:        fmt.Sprintf("[infrastructure.data.repositories.persistence.CreateUser] fail to insert User to Database"),
			OriginalError: result.Error,
		}
	}

	return ent, nil
}

func (ar AuthRepositoryImpl) FindUserById(ctx context.Context, ID uint) (*entities.UserEnt, error) {

	db := repositories.GetConn()

	userEnt := &entities.UserEnt{}

	err := db.Take(&userEnt, "id=?", ID)

	if err != nil {
		return nil, &repositories.NotFoundError{
			Msg:           repositories.DefaultNotFoundMsg,
			ErrMsg:        fmt.Sprintf("[infrastructure.data.repositories.persistence.FindUserById] failed to find userEnt from rdb. ID : %d", ID),
			OriginalError: err.Error,
		}
	}
	return userEnt, nil
}

func CreateAuthRepository() IAuthRepository {
	return &AuthRepositoryImpl{}
}
