package repository

import (
	"gorm.io/gorm"
	"hacktiv8_final/posts"
)

type Repository struct {
	Db          *gorm.DB
	PhotoRepo   posts.RepositoryPhoto
	CommentRepo posts.RepositoryComment
}

func (r *Repository) Save(data Comment) {
	result := r.Db.Create(&data)
	panic(result.Error)
}

func (r *Repository) FindAll(data Comment) Comment {
	result := r.Db.Find(&data)
	panic(result.Error)
	return data
}

func (r *Repository) FindById(id int) (Comment, error) {
	result := r.Db.Find(&Comment{}, id)
	if result != nil {
		return Comment{}, nil
	} else {
		return Comment{}, result.Error
	}

}

func (r *Repository) Delete(id int) {
	result := r.Db.Where("id = ?", id).Delete(&Comment{})
	panic(result.Error)
}

func (r *Repository) Update(data interface{}) {

}
