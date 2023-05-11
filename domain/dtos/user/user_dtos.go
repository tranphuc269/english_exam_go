package dtos

import (
	"english_exam_go/infrastructure/data/entities"
	auth_utils "english_exam_go/utils/auth"
	"english_exam_go/utils/resource"
	"gorm.io/gorm"
	"time"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Id          uint   `json:"id"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type RegisterAccountRequest struct {
	Name        string            `json:"name"`
	Email       string            `json:"email"`
	Password    string            `json:"password"`
	DateOfBirth string            `json:"date_of_birth"`
	PhoneNumber string            `json:"phone_number"`
	Address     string            `json:"address"`
	Role        resource.UserRole `json:"role,omitempty" binding:"required,userRoleEnum"`
}

type UpdateAccountRequest struct {
	Name        string `json:"name"`
	Password    string `json:"password"`
	DateOfBirth string `json:"date_of_birth"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

type UserResponse struct {
	gorm.Model
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Role        int       `json:"role"`
}

func (rar *RegisterAccountRequest) RegisterAccountToUserEnt() (*entities.User, error) {
	date, err := time.Parse("2006-01-02", rar.DateOfBirth)
	if err != nil {
		return nil, err
	}
	hashPassword, err := auth_utils.HashPassword(rar.Password)
	if err != nil {
		return nil, err
	}
	return &entities.User{
		Model:       gorm.Model{},
		Name:        rar.Name,
		Email:       rar.Email,
		Password:    hashPassword,
		PhoneNumber: rar.PhoneNumber,
		Address:     rar.Address,
		DateOfBirth: date,
		Role:        rar.Role,
	}, nil
}

func UserEntToResponse(ue *entities.User) *UserResponse {
	return &UserResponse{
		Model:       ue.Model,
		Name:        ue.Name,
		Email:       ue.Email,
		PhoneNumber: ue.PhoneNumber,
		Address:     ue.Address,
		DateOfBirth: ue.DateOfBirth,
		Role:        ue.Role.ToInt(),
	}
}
