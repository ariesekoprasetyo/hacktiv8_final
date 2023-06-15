package user

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"hacktiv8_final/repository"
)

type CreateUsersRequest struct {
	Username string `validate:"required" json:"username"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=6" json:"password"`
	Age      int    `validate:"required,min=8" json:"age"`
}

type LoginUsersRequest struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}

type Authz struct {
	AuthRepo  RepositoryUsers
	Validate  *validator.Validate
	JwtSecret string
}

func (a *Authz) Login(request LoginUsersRequest) (string, error) {
	//Find username in database
	result, err := a.AuthRepo.FindByUsername(request.Username)
	if err != nil {
		return "", errors.New("incorrect credential")
	}
	verifyError := VerifyPassword(result.Password, request.Password)
	if verifyError != nil {
		return "", errors.New("invalid username or password")
	}
	//Generate Token
	token, errToken := GenerateToken(result.ID, a.JwtSecret)
	if errToken != nil {
		return "", errToken
	}
	tokenFinal := "Bearer " + token
	return tokenFinal, nil
}

func (a *Authz) Register(request CreateUsersRequest) error {
	err := a.Validate.Struct(request)
	if err != nil {
		return err
	}
	hashedPassword, _ := HashPassword(request.Password)
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

func (a *Authz) ValidateToken(token string) (interface{}, error) {
	return ValidateToken(token, a.JwtSecret)
}
