package controller

import (
	"hacktiv8_final/user"
)

type SocialMediaService interface {
	CreateSocialMedia(request user.SocialMediaRequest) error
	UpdateSocialMedia(socialMediaId uint, request user.SocialMediaUpdate) error
	DeleteSocialMedia(socialMediaId uint) error
	FindByIdSocialMedia(socialMediaId uint) (user.ResponSocialMedia, error)
	FindAllSocialMedia() []user.ResponSocialMedia
}
