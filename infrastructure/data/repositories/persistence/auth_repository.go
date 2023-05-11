package persistence

import (
	"context"
	"english_exam_go/infrastructure/data/entities"
	"english_exam_go/infrastructure/data/repositories"
	"fmt"
)

type IAuthRepository interface {
	CreateUser(context.Context, *entities.User) (*entities.User, error)
	FindUserByEmail(context.Context, string) (*entities.User, error)
	FindById(context.Context, int) (*entities.User, error)
	GetUsers(context.Context, int) []*entities.User
	UpdateUser(context.Context, *entities.User) error
}

type AuthRepositoryImpl struct {
}

func (ar AuthRepositoryImpl) FindById(ctx context.Context, ID int) (*entities.User, error) {
	//TODO implement me
	db := repositories.GetConn()
	userEnt := entities.User{}
	err := db.Model(&entities.User{}).First(&userEnt, ID)
	if err.Error != nil {
		return nil, &repositories.NotFoundError{
			Msg:           repositories.DefaultNotFoundMsg,
			ErrMsg:        fmt.Sprintf("[infrastructure.data.repositories.persistence.FindUserByID] failed to find userEnt from rdb. id : %d", ID),
			OriginalError: err.Error,
		}
	}

	return &userEnt, nil
}

func (ar AuthRepositoryImpl) UpdateUser(ctx context.Context, user *entities.User) error {
	//TODO implement me
	db := repositories.GetConn()
	result := db.Save(&user)
	return result.Error
}

func (ar AuthRepositoryImpl) GetUsers(ctx context.Context, role int) []*entities.User {
	//TODO implement me
	db := repositories.GetConn()
	var userEntities []*entities.User
	_ = db.Where("role = ?", role).Order("created_at").Find(&userEntities)
	return userEntities
}

func (ar AuthRepositoryImpl) CreateUser(ctx context.Context, ent *entities.User) (*entities.User, error) {
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

func (ar AuthRepositoryImpl) FindUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	db := repositories.GetConn()
	userEnt := entities.User{}
	err := db.Model(&entities.User{}).First(&userEnt)
	if err.Error != nil {
		return nil, &repositories.NotFoundError{
			Msg:           repositories.DefaultNotFoundMsg,
			ErrMsg:        fmt.Sprintf("[infrastructure.data.repositories.persistence.FindUserByEmail] failed to find userEnt from rdb. email : %s", email),
			OriginalError: err.Error,
		}
	}

	return &userEnt, nil
}

func CreateAuthRepository() IAuthRepository {
	return &AuthRepositoryImpl{}
}
