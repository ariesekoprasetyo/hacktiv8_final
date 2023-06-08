package user

import (
	"golang.org/x/crypto/bcrypt"
	"hacktiv8_final/posts"
	"hacktiv8_final/repository"
)

type createUsersRequest struct {
	Username string `validate:"required,min=2,max=100" json:"username"`
	Email    string `validate:"required,min=2,max=100" json:"email"`
	Password string `validate:"required,min=2,max=100" json:"password"`
}

type Authz struct {
	Repo posts.Repository
}

func (a *Authz) Register(CreateUserRequest createUsersRequest) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(CreateUserRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	newUser := repository.User{
		Username: CreateUserRequest.Username,
		Email:    CreateUserRequest.Email,
		Password: string(hashedPassword),
		Age:      30,
	}
	a.Repo.Save(newUser)
}
