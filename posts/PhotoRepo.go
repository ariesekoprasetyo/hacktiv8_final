package posts

import "hacktiv8_final/repository"

type RepositoryPhoto interface {
	SavePhoto(Photo repository.Photo)
	UpdatePhoto(photoId uint, Photo repository.Photo)
	DeletePhoto(PhotoId uint)
	FindByIdPhoto(PhotoId uint) (repository.Photo, error)
	FindAllPhoto() []repository.Photo
}
