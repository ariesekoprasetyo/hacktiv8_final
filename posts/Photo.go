package posts

import (
	"github.com/go-playground/validator/v10"
	"hacktiv8_final/controller"
	"hacktiv8_final/repository"
)

type PhotoService struct {
	PhotoRepo RepositoryPhoto
	Validate  *validator.Validate
}

func (p *PhotoService) CreatePhoto(request controller.CreatePhotoRequest) error {
	if err := p.Validate.Struct(request); err != nil {
		return err
	}
	photo := repository.Photo{
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoUrl: request.PhotoURL,
		UserId:   request.UserID,
	}
	p.PhotoRepo.SavePhoto(photo)
	return nil
}

func (p *PhotoService) UpdatePhoto(request controller.UpdatePhotoRequest) error {
	if err := p.Validate.Struct(request); err != nil {
		return err
	}
	photoData, err := p.PhotoRepo.FindByIdPhoto(request.ID)
	if err != nil {
		return err
	}
	photoData.Title = request.Title
	photoData.Caption = request.Caption
	photoData.PhotoUrl = request.PhotoURL
	p.PhotoRepo.UpdatePhoto(photoData)
	return nil

}

func (p *PhotoService) DeletePhoto(photoId uint) error {
	_, err := p.PhotoRepo.FindByIdPhoto(photoId)
	if err != nil {
		return err
	}
	p.PhotoRepo.DeletePhoto(photoId)
	return nil
}

func (p *PhotoService) FindByIdPhoto(photoId uint) (repository.Photo, error) {
	photoByIdResult, err := p.PhotoRepo.FindByIdPhoto(photoId)
	if err != nil {
		return photoByIdResult, err
	}
	return photoByIdResult, nil
}

func (p *PhotoService) FindAllPhoto() []repository.Photo {
	var photos []repository.Photo
	result := p.PhotoRepo.FindAllPhoto()
	for _, value := range result {
		photo := repository.Photo{
			Title:    value.Title,
			Caption:  value.Caption,
			PhotoUrl: value.PhotoUrl,
			UserId:   value.UserId,
			User:     value.User,
			Comment:  value.Comment,
		}
		photos = append(photos, photo)
	}
	return photos
}
