package entities

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserEnt struct {
	gorm.Model
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

func (ue *UserEnt) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(ue.Password), []byte(providedPassword))
	return err
}
