package repository

import (
	"errors"
	"gorm.io/gorm"
)

type UsersRepo struct {
	Db *gorm.DB
}

func (u *UsersRepo) SaveUser(Users User) {
	result := u.Db.Create(&Users)
	panic(result.Error)
}

func (u *UsersRepo) UpdateUser(Users User) {
	result := u.Db.Model(&Users).Where("id = ?", Users.ID).Updates(Users)
	panic(result.Error)
}

func (u *UsersRepo) DeleteUser(UsersId uint) {
	var User User
	result := u.Db.Where("id = ?", UsersId).Delete(&User)
	panic(result.Error)
}

func (u *UsersRepo) FindByIdUser(UserId uint) (User, error) {
	var User User
	result := u.Db.First(&User, UserId)
	if result != nil {
		return User, nil
	}
	return User, result.Error
}

func (u *UsersRepo) FindAllUser() []User {
	var Users []User
	result := u.Db.Find(&User{})
	panic(result.Error)
	return Users
}

func (u *UsersRepo) FindByUsername(Username string) (User, error) {
	var Users User
	result := u.Db.First(&Users, "username = ?", Username)
	if result != nil {
		return Users, errors.New("invalid username or password")
	}
	return Users, nil
}
