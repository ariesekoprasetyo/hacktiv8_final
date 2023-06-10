package posts

import (
	"hacktiv8_final/controller"
	"hacktiv8_final/repository"
)

type PhotoService struct {
	PhotoRepo RepositoryPhoto
}

func (p PhotoService) CreatePhoto(request controller.CreatePhotoRequest) {
	//TODO implement me
	panic("implement me")
}

func (p PhotoService) UpdatePhoto(request controller.UpdatePhotoRequest) {
	//TODO implement me
	panic("implement me")
}

func (p PhotoService) SavePhoto(Photo repository.Photo) {
	//TODO implement me
	panic("implement me")
}
func (p PhotoService) DeletePhoto(PhotoId uint) {
	//TODO implement me
	panic("implement me")
}

func (p PhotoService) FindByIdPhoto(PhotoId uint) (repository.Photo, error) {
	//TODO implement me
	panic("implement me")
}

func (p PhotoService) FindAllPhoto() []repository.Photo {
	//TODO implement me
	panic("implement me")
}
