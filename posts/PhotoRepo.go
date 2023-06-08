package posts

import "hacktiv8_final/repository"

type RepositoryPhoto interface {
	Save(Photo repository.Photo)
	Update(Photo repository.Photo)
	Delete(PhotoId int)
	FindById(PhotoId int) (repository.Photo, error)
	FindAll() []repository.Photo
}
