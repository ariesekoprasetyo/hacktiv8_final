package controller

import "hacktiv8_final/repository"

type PhotoService interface {
	CreatePhoto(request CreatePhotoRequest)
	UpdatePhoto(request UpdatePhotoRequest)
	DeletePhoto(photoId uint)
	FindByIdPhoto(photoId uint) (repository.Photo, error)
	FindAllPhoto() []repository.Photo
}
