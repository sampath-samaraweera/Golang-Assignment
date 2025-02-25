package models

import "github.com/golang-jwt/jwt/v5"

type User struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
type UserToken struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

type UserLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UpdateUser struct {
	NewUserName string `json:"new_user_name"`
}