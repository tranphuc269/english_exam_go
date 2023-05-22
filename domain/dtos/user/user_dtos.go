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
	Name        string            `form:"name"`
	Email       string            `form:"email"`
	Password    string            `form:"password"`
	Code        string            `form:"code"`
	DateOfBirth string            `form:"date_of_birth"`
	Avatar      string            `form:"avatar"`
	PhoneNumber string            `form:"phone_number"`
	Address     string            `form:"address"`
	Role        resource.UserRole `form:"role,omitempty" binding:"required,userRoleEnum"`
}

type UpdateAccountRequest struct {
	Name        string `form:"name"`
	Password    string `form:"password"`
	DateOfBirth string `form:"date_of_birth"`
	Avatar      string `form:"avatar"`
	Code        string `form:"code"`
	PhoneNumber string `form:"phone_number"`
	Address     string `form:"address"`
}

type UserResponse struct {
	gorm.Model
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Code        string    `json:"code"`
	Avatar      string    `json:"avatar"`
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
		Code:        rar.Code,
		Avatar:      rar.Avatar,
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
		Code:        ue.Code,
		Avatar:      ue.Avatar,
	}
}
