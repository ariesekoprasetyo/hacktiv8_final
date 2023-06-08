package repository

import (
	"gorm.io/gorm"
	"hacktiv8_final/posts"
)

type Repository struct {
	Db   *gorm.DB
	repo posts.Repository
}

func (r *Repository) Save(data interface{}) {
	result := r.Db.Create(&data)
	panic(result.Error)
}

func (r *Repository) FindAll(data interface{}) interface{} {
	result := r.Db.Find(&data)
	panic(result.Error)
	return data
}

func (r *Repository) FindById(id int, data interface{}) (interface{}, error) {
	result := r.Db.Find(&data, id)
	if result != nil {
		return data, nil
	} else {
		return data, result.Error
	}

}

func (r *Repository) Delete(id int, data interface{}) {
	result := r.Db.Where("id = ?", id).Delete(&data)
	panic(result.Error)
}

func (r *Repository) Update(data interface{}) {

}
