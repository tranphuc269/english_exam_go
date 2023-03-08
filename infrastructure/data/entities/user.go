package entities

import (
	"english_exam_go/utils/resource"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type UserEnt struct {
	gorm.Model
	Name        string            `json:"name"`
	Email       string            `json:"email"`
	Password    string            `json:"password"`
	PhoneNumber string            `json:"phone_number"`
	Address     string            `json:"address"`
	DateOfBirth time.Time         `json:"date_of_birth"`
	Role        resource.UserRole `json:"role"`
}

func (ue *UserEnt) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(ue.Password), []byte(providedPassword))
	return err
}
