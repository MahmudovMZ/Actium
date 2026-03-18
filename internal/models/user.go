package models

import (
	"Actium_Todo/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID        int64  `json:"user_id"`
	UserName  string `json:"user_name"`
	Password  string `json:"user_pass"`
	CreatedAt string `json:"joined_at"`
}

var Users = make(map[int]User)

var Secret = config.JWTSecret // signature to check the token

type Claims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}
