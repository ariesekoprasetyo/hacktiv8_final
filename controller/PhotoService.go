package controller

import (
	"hacktiv8_final/posts"
	"hacktiv8_final/repository"
)

type PhotoService interface {
	CreatePhoto(request posts.CreatePhotoRequest) error
	UpdatePhoto(photoId uint, request posts.UpdatePhotoRequest) error
	DeletePhoto(photoId uint) error
	FindByIdPhoto(photoId uint) (repository.Photo, error)
	FindAllPhoto() []posts.ResponPhoto
}
