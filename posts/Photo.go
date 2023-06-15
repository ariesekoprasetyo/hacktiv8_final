package posts

import (
	"errors"
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
	Title    string                     `json:"title"`
	Caption  string                     `json:"caption"`
	PhotoUrl string                     `json:"photoUrl"`
	UserID   uint                       `json:"userID"`
	Username string                     `json:"username"`
	Comment  []repository.CommentRespon `json:"comment"`
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

func (p *PhotoService) UpdatePhotoById(photoId uint, request UpdatePhotoRequest, userId uint) error {
	if err := p.Validate.Struct(request); err != nil {
		return err
	}
	photoData, err := p.PhotoRepo.FindByIdPhoto(photoId)
	if err != nil {
		return err
	}
	if photoData.UserID != userId {
		return errors.New("forbiden")
	}
	photoData.Title = request.Title
	photoData.Caption = request.Caption
	photoData.PhotoUrl = request.PhotoURL
	p.PhotoRepo.UpdatePhoto(photoId, photoData)
	return nil

}

func (p *PhotoService) DeletePhotoById(photoId uint, userId uint) error {
	result, err := p.PhotoRepo.FindByIdPhoto(photoId)
	if err != nil {
		return err
	}
	if result.ID != userId {
		return errors.New("forbiden")
	}
	p.PhotoRepo.DeletePhoto(photoId)
	return nil
}

func (p *PhotoService) FindByIdPhoto(photoId uint) (ResponPhoto, error) {
	var finalResult ResponPhoto
	photoByIdResult, err := p.PhotoRepo.FindByIdPhoto(photoId)
	if err != nil {
		return finalResult, err
	}
	finalResult.Title = photoByIdResult.Title
	finalResult.Caption = photoByIdResult.Caption
	finalResult.PhotoUrl = photoByIdResult.PhotoUrl
	finalResult.UserID = photoByIdResult.UserID
	finalResult.Username = photoByIdResult.User.Username
	finalResult.Comment = photoByIdResult.Comment

	return finalResult, nil
}

func (p *PhotoService) FindAllPhoto() []ResponPhoto {
	var finalResult []ResponPhoto
	result := p.PhotoRepo.FindAllPhoto()
	for _, value := range result {
		finalResult = append(finalResult, ResponPhoto{
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
