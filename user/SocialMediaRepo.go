package user

import "hacktiv8_final/repository"

type RepositorySocialMedia interface {
	SaveSocialMedia(SocialMedia repository.SocialMedia)
	UpdateSocialMedia(SocialMedia repository.SocialMedia)
	DeleteSocialMedia(SocialMediaId uint)
	FindByIdSocialMedia(SocialMediaId uint) (repository.SocialMedia, error)
	FindAllSocialMedia() []repository.SocialMedia
}
