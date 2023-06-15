package controller

import (
	"hacktiv8_final/user"
)

type SocialMediaService interface {
	CreateSocialMedia(request user.SocialMediaRequest) error
	UpdateSocialMediaById(socialMediaId uint, request user.SocialMediaUpdate, userId uint) error
	DeleteSocialMediaById(socialMediaId uint, userId uint) error
	FindByIdSocialMedia(socialMediaId uint) (user.ResponSocialMedia, error)
	FindAllSocialMedia() []user.ResponSocialMedia
}
