package user

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"hacktiv8_final/controller"
	"hacktiv8_final/repository"
)

type Authz struct {
	AuthRepo RepositoryUsers
	Validate *validator.Validate
}

func (a *Authz) Login(request controller.LoginUsersRequest) (string, error) {
	//Find username in database
	result, err := a.AuthRepo.FindByUsername(request.Username)
	if err != nil {
		return "", errors.New("incorrect credential")
	}
	verify_error := controller.VerifyPassword(result.Password, request.Password)
	if verify_error != nil {
		return "", errors.New("invalid username or password")
	}
	//Generate Token
	token, errToken := controller.GenerateToken(result.ID)
	if errToken != nil {
		return "", errToken
	}
	return token, nil
}

func (a *Authz) Register(request controller.CreateUsersRequest) error {
	err := a.Validate.Struct(request)
	if err != nil {
		return err
	}
	hashedPassword, err := controller.HashPassword(request.Password)
	if err != nil {
		return err
	}
	newUser := repository.User{
		Username: request.Username,
		Email:    request.Email,
		Password: hashedPassword,
		Age:      request.Age,
	}
	a.AuthRepo.SaveUser(newUser)
	return nil
}

func (a *Authz) FindUserById(userId uint) (uint, error) {
	result, err := a.AuthRepo.FindByIdUser(userId)
	if err != nil {
		return 0, err
	}
	return result.ID, err
}
