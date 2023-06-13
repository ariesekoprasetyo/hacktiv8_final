package repository

import (
	"gorm.io/gorm"
	"log"
)

type PhotoRepo struct {
	Db *gorm.DB
}

func (p PhotoRepo) SavePhoto(photo Photo) {
	err := p.Db.Create(&photo).Error
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func (p PhotoRepo) UpdatePhoto(photoId uint, photo Photo) {
	err := p.Db.Model(&photo).Where("id = ?", photoId).Updates(photo).Error
	if err != nil {
		panic(err)
	}
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
	err := p.Db.Preload("User").First(&photoById, PhotoId).Error
	if err != nil {
		return photoById, err
	}
	return photoById, nil
}

func (p PhotoRepo) FindAllPhoto() []Photo {
	var allPhoto []Photo
	err := p.Db.Preload("User").Preload("Comment").Find(&allPhoto).Error
	if err != nil {
		panic(err)
	}
	return allPhoto
}
