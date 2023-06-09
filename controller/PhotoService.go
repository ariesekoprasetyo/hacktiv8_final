package controller

import (
	"hacktiv8_final/posts"
)

type PhotoService interface {
	CreatePhoto(request posts.CreatePhotoRequest) error
	UpdatePhotoById(photoId uint, request posts.UpdatePhotoRequest, userId uint) error
	DeletePhotoById(photoId uint, userId uint) error
	FindByIdPhoto(photoId uint) (posts.ResponPhoto, error)
	FindAllPhoto() []posts.ResponPhoto
}
