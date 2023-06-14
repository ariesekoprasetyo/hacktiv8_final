package posts

import (
	"github.com/go-playground/validator/v10"
	"hacktiv8_final/repository"
)

type CreatePhotoRequest struct {
	Title    string `validate:"required" json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `validate:"required" json:"photo_url"`
	UserID   uint   `json:"user_id"`
}
type UpdatePhotoRequest struct {
	Title    string `validate:"required" json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `validate:"required" json:"photo_url"`
}
type PhotoService struct {
	PhotoRepo RepositoryPhoto
	Validate  *validator.Validate
}

type ResponPhoto struct {
	ID       uint
	Title    string
	Caption  string
	PhotoUrl string
	UserID   uint
	Username string
	Comment  []repository.Comment
}

func (p *PhotoService) CreatePhoto(request CreatePhotoRequest) error {
	if err := p.Validate.Struct(request); err != nil {
		return err
	}
	photo := repository.Photo{
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoUrl: request.PhotoURL,
		UserID:   request.UserID,
	}
	p.PhotoRepo.SavePhoto(photo)
	return nil
}

func (p *PhotoService) UpdatePhoto(photoId uint, request UpdatePhotoRequest) error {
	if err := p.Validate.Struct(request); err != nil {
		return err
	}
	photoData, err := p.PhotoRepo.FindByIdPhoto(photoId)
	if err != nil {
		return err
	}
	photoData.Title = request.Title
	photoData.Caption = request.Caption
	photoData.PhotoUrl = request.PhotoURL
	p.PhotoRepo.UpdatePhoto(photoId, photoData)
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

func (p *PhotoService) FindAllPhoto() []ResponPhoto {
	var finalResult []ResponPhoto
	result := p.PhotoRepo.FindAllPhoto()
	for _, value := range result {
		finalResult = append(finalResult, ResponPhoto{
			ID:       value.ID,
			Title:    value.Title,
			Caption:  value.Caption,
			PhotoUrl: value.PhotoUrl,
			UserID:   value.UserID,
			Username: value.User.Username,
			Comment:  value.Comment,
		})
	}
	return finalResult
}
