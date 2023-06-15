package repository

import "gorm.io/gorm"

type SocialMediaRepo struct {
	Db *gorm.DB
}

func (s *SocialMediaRepo) SaveSocialMedia(SocialMedia SocialMedia) {
	err := s.Db.Create(&SocialMedia).Error
	if err != nil {
		panic(err)
	}
}

func (s *SocialMediaRepo) UpdateSocialMedia(socialMediaId uint, SocialMedia SocialMedia) {
	err := s.Db.Model(&SocialMedia).Where("id = ?", socialMediaId).Updates(SocialMedia).Error
	if err != nil {
		panic(err)
	}
}

func (s *SocialMediaRepo) DeleteSocialMedia(id uint) {
	var sosialMedia SocialMedia
	err := s.Db.Where("id = ?", id).Delete(&sosialMedia).Error
	if err != nil {
		panic(err)
	}
}

func (s *SocialMediaRepo) FindByIdSocialMedia(SocialMediaId uint) (SocialMedia, error) {
	var socialMediaById SocialMedia
	err := s.Db.Preload("User").First(&socialMediaById, SocialMediaId).Error
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
