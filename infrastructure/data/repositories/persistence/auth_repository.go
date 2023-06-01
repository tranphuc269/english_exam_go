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
	GetUsers(context.Context, int, int, int, string, string) ([]*entities.User, int)
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

func (ar AuthRepositoryImpl) GetUsers(ctx context.Context, role int, offset int, limit int, name string, code string) ([]*entities.User, int) {
	//TODO implement me
	db := repositories.GetConn()
	var userEntities []*entities.User
	var userEntities2 []*entities.User
	query := db.Offset(offset).Limit(limit).Where("role = ?", role).Order("created_at")
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	if code != "" {
		query = query.Where("code LIKE ?", "%"+code+"%")
	}
	_ = query.Find(&userEntities)

	query2 := db.Where("role = ?", role).Order("created_at")
	if name != "" {
		query2 = query.Where("name LIKE ?", "%"+name+"%")
	}

	if code != "" {
		query2 = query.Where("code LIKE ?", "%"+code+"%")
	}
	_ = query2.Find(&userEntities2)
	return userEntities, len(userEntities2)
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
	err := db.Model(&entities.User{}).Where("email = ?", email).Order("created_at").First(&userEnt)
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
