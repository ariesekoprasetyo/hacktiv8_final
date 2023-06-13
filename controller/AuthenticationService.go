package controller

type AuthService interface {
	Login(request LoginUsersRequest) (string, error)
	Register(request CreateUsersRequest) error
	FindUserById(userId uint) (uint, error)
}
