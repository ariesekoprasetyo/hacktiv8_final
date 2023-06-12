package controller

import "hacktiv8_final/repository"

type PhotoService interface {
	CreatePhoto(request CreatePhotoRequest) error
	UpdatePhoto(request UpdatePhotoRequest) error
	DeletePhoto(photoId uint) error
	FindByIdPhoto(photoId uint) (repository.Photo, error)
	FindAllPhoto() []repository.Photo
}
