package utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var jwtKey = []byte("tranphuc")

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GenerateJWT(email string, username string, role string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email:    email,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return tokenString, err
}
func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}

func ParseToken(signedToken string) (jwtClaim *JWTClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	//if err != nil {
	//	return
	//}
	claims, _ := token.Claims.(*JWTClaim)
	fmt.Printf("auth : %s", claims.Username)
	return claims, err
}

func HashPassword(password string) (response string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
