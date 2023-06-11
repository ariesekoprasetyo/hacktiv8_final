package repository

import (
	"gorm.io/gorm"
)

type PhotoRepo struct {
	Db *gorm.DB
}

func (p PhotoRepo) SavePhoto(photo Photo) {
	result := p.Db.Create(&photo)
	panic(result.Error)
}

func (p PhotoRepo) UpdatePhoto(photo Photo) {
	result := p.Db.Model(&photo).Updates(photo)
	panic(result.Error)
}

func (p PhotoRepo) DeletePhoto(PhotoId uint) {
	var photo Photo
	result := p.Db.Where("id = ?", PhotoId).Delete(photo)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (p PhotoRepo) FindByIdPhoto(PhotoId uint) (Photo, error) {
	var photoById Photo
	result := p.Db.Find(&photoById, PhotoId)
	if result != nil {
		return photoById, nil
	}
	return photoById, result.Error
}

func (p PhotoRepo) FindAllPhoto() []Photo {
	var allPhoto []Photo
	result := p.Db.Preload("Comments").Preload("Users").Find(&allPhoto)
	panic(result.Error)
	return allPhoto
}
