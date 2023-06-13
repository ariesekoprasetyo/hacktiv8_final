package repository

import "gorm.io/gorm"

type SocialMediaRepo struct {
	Db *gorm.DB
}

func (s *SocialMediaRepo) SaveSocialMedia(SocialMedia SocialMedia) {
	result := s.Db.Create(&SocialMedia)
	panic(result.Error)
}

func (s *SocialMediaRepo) UpdateSocialMedia(SocialMedia SocialMedia) {
	result := s.Db.Model(&SocialMedia).Where("id = ?", SocialMedia.ID).Updates(SocialMedia)
	panic(result.Error)
}

func (s *SocialMediaRepo) DeleteSocialMedia(SocialMediaId uint) {
	var sosialMedia SocialMedia
	result := s.Db.Where("id = ?", SocialMediaId).Delete(&sosialMedia)
	panic(result.Error)
}

func (s *SocialMediaRepo) FindByIdSocialMedia(SocialMediaId uint) (SocialMedia, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SocialMediaRepo) FindAllSocialMedia() []SocialMedia {
	//TODO implement me
	panic("implement me")
}
