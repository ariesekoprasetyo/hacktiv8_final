package user

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"hacktiv8_final/repository"
)

type SocialMediaRequest struct {
	Name      string `validate:"required" json:"name"`
	SosmedUrl string `validate:"required" json:"sosmedUrl"`
	UserId    uint   `json:"user_id"`
}

type SocialMediaUpdate struct {
	Name string `validate:"required" json:"name"`
}

type ResponSocialMedia struct {
	Name      string                `json:"name"`
	SosmedUrl string                `json:"sosmedUrl"`
	UserID    uint                  `json:"userID"`
	User      repository.UserRespon `json:"user"`
}

type SocialMediaService struct {
	SocialMediaRepo RepositorySocialMedia
	Validate        *validator.Validate
}

func (s *SocialMediaService) CreateSocialMedia(request SocialMediaRequest) error {
	if err := s.Validate.Struct(request); err != nil {
		return err
	}
	socialmedia := repository.SocialMedia{
		Name:      request.Name,
		SosmedUrl: request.SosmedUrl,
		UserID:    request.UserId,
	}
	s.SocialMediaRepo.SaveSocialMedia(socialmedia)
	return nil
}

func (s *SocialMediaService) UpdateSocialMediaById(socialMediaId uint, request SocialMediaUpdate, userId uint) error {
	err := s.Validate.Struct(request)
	if err != nil {
		return err
	}
	socialediaData, err := s.SocialMediaRepo.FindByIdSocialMedia(socialMediaId)
	if err != nil {
		return err
	}
	if socialediaData.UserID != userId {
		return errors.New("forbidden")
	}
	socialediaData.Name = request.Name
	s.SocialMediaRepo.UpdateSocialMedia(socialMediaId, socialediaData)
	return nil
}

func (s *SocialMediaService) DeleteSocialMediaById(socialMediaId uint, userId uint) error {
	result, err := s.SocialMediaRepo.FindByIdSocialMedia(socialMediaId)
	if err != nil {
		return err
	}
	if result.UserID != userId {
		return errors.New("forbidden")
	}
	s.SocialMediaRepo.DeleteSocialMedia(socialMediaId)
	return nil
}

func (s *SocialMediaService) FindByIdSocialMedia(socialMediaId uint) (ResponSocialMedia, error) {
	var finalResult ResponSocialMedia
	socialmediaresult, err := s.SocialMediaRepo.FindByIdSocialMedia(socialMediaId)
	if err != nil {
		return finalResult, err
	}
	finalResult.Name = socialmediaresult.Name
	finalResult.SosmedUrl = socialmediaresult.SosmedUrl
	finalResult.UserID = socialmediaresult.UserID
	finalResult.User = socialmediaresult.User
	return finalResult, nil
}

func (s *SocialMediaService) FindAllSocialMedia() []ResponSocialMedia {
	var finalResult []ResponSocialMedia
	result := s.SocialMediaRepo.FindAllSocialMedia()
	for _, value := range result {
		finalResult = append(finalResult, ResponSocialMedia{
			Name:      value.Name,
			SosmedUrl: value.SosmedUrl,
			UserID:    value.UserID,
			User:      value.User,
		})
	}
	return finalResult
}
