package dtos

import (
	"english_exam_go/infrastructure/data/entities"
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
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	DateOfBirth string `json:"date_of_birth"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

func (rar *RegisterAccountRequest) RegisterAccountToUserEnt() (*entities.UserEnt, error) {
	date, err := time.Parse("2006-01-02", rar.DateOfBirth)
	if err != nil {
		return nil, err
	}
	return &entities.UserEnt{
		Model:       gorm.Model{},
		Name:        rar.Name,
		Email:       rar.Email,
		Password:    rar.Password,
		PhoneNumber: rar.PhoneNumber,
		Address:     rar.Address,
		DateOfBirth: date,
	}, nil
}
