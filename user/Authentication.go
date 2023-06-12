package user

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"hacktiv8_final/controller"
	"hacktiv8_final/repository"
	"log"
)

type Authz struct {
	AuthRepo RepositoryUsers
	Validate *validator.Validate
}

func (a *Authz) Login(request controller.LoginUsersRequest) (string, error) {
	log.Println(request)
	return "", nil
}

func (a *Authz) Register(request controller.CreateUsersRequest) error {
	err := a.Validate.Struct(request)
	if err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	newUser := repository.User{
		Username: request.Username,
		Email:    request.Email,
		Password: string(hashedPassword),
		Age:      request.Age,
	}
	a.AuthRepo.SaveUser(newUser)
	return nil
}

//func (a *Authz) Register(bodyReq createuser) {
//	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(CreateUserRequest.Password), bcrypt.DefaultCost)
//	if err != nil {
//		panic(err)
//	}
//	newUser := repository.User{
//		Username: CreateUserRequest.Username,
//		Email:    CreateUserRequest.Email,
//		Password: string(hashedPassword),
//		Age:      30,
//	}
//	a.Repo.Save(newUser)
//}
