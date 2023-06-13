package controller

import "hacktiv8_final/user"

type AuthService interface {
	Login(request user.LoginUsersRequest) (string, error)
	Register(request user.CreateUsersRequest) error
	FindUserById(userId uint) (uint, error)
	ValidateToken(token string) (interface{}, error)
}
