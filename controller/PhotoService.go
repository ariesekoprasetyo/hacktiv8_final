package controller

import (
	"hacktiv8_final/posts"
)

type PhotoService interface {
	CreatePhoto(request posts.CreatePhotoRequest) error
	UpdatePhoto(photoId uint, request posts.UpdatePhotoRequest) error
	DeletePhoto(photoId uint) error
	FindByIdPhoto(photoId uint) (posts.ResponPhoto, error)
	FindAllPhoto() []posts.ResponPhoto
}
