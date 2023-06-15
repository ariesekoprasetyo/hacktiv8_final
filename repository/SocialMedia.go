package repository

import "gorm.io/gorm"

type SocialMediaRepo struct {
	Db *gorm.DB
}

func (s *SocialMediaRepo) SaveSocialMedia(SocialMedia SocialMedia) {
	result := s.Db.Create(&SocialMedia)
	panic(result.Error)
}

func (s *SocialMediaRepo) UpdateSocialMedia(socialMediaId uint, SocialMedia SocialMedia) {
	result := s.Db.Model(&SocialMedia).Where("id = ?", socialMediaId).Updates(SocialMedia)
	panic(result.Error)
}

func (s *SocialMediaRepo) DeleteSocialMedia(SocialMediaId uint) {
	var sosialMedia SocialMedia
	result := s.Db.Where("id = ?", SocialMediaId).Delete(&sosialMedia)
	panic(result.Error)
}

func (s *SocialMediaRepo) FindByIdSocialMedia(SocialMediaId uint) (SocialMedia, error) {
	var socialMediaById SocialMedia
	err := s.Db.Preload("User").First(socialMediaById, SocialMediaId).Error
	if err != nil {
		return socialMediaById, err
	}
	return socialMediaById, nil
}

func (s *SocialMediaRepo) FindAllSocialMedia() []SocialMedia {
	var socialmedia []SocialMedia
	err := s.Db.Preload("User").Find(&socialmedia).Error
	if err != nil {
		panic(err)
	}
	return socialmedia
}
