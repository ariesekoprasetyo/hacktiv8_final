package user

import "hacktiv8_final/repository"

type RepositoryUsers interface {
	SaveUser(Users repository.User)
	UpdateUser(Users repository.User)
	DeleteUser(UsersId uint)
	FindByIdUser(UserId uint) (repository.User, error)
	FindAllUser() []repository.User
	FindByUsername(Username string) (repository.User, error)
}
