package controller

import (
	"hacktiv8_final/posts"
	"hacktiv8_final/repository"
)

type PhotoService interface {
	CreatePhoto(request posts.CreatePhotoRequest) error
	UpdatePhoto(request posts.UpdatePhotoRequest) error
	DeletePhoto(photoId uint) error
	FindByIdPhoto(photoId uint) (repository.Photo, error)
	FindAllPhoto() []repository.Photo
}
