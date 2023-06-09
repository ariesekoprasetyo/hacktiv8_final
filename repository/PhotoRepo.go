package repository

import "gorm.io/gorm"

type PhotoRepo struct {
	Db *gorm.DB
}

func (p PhotoRepo) SavePhoto(Photo Photo) {
	//TODO implement me
	panic("implement me")
}

func (p PhotoRepo) UpdatePhoto(Photo Photo) {
	//TODO implement me
	panic("implement me")
}

func (p PhotoRepo) DeletePhoto(PhotoId int) {
	//TODO implement me
	panic("implement me")
}

func (p PhotoRepo) FindByIdPhoto(PhotoId int) (Photo, error) {
	//TODO implement me
	panic("implement me")
}

func (p PhotoRepo) FindAllPhoto() []Photo {
	//TODO implement me
	panic("implement me")
}
