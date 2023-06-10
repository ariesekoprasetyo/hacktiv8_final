package repository

import (
	"gorm.io/gorm"
)

type CommentRepo struct {
	Db *gorm.DB
}

func (r *CommentRepo) SaveComment(data Comment) {
	result := r.Db.Debug().Create(&data)
	panic(result.Error)
}

func (r *CommentRepo) FindAllComment() []Comment {
	var allComment []Comment
	result := r.Db.Find(&Comment{})
	panic(result.Error)
	return allComment
}

func (r *CommentRepo) FindByIdComment(id uint) (Comment, error) {
	result := r.Db.Find(&Comment{}, id)
	if result != nil {
		return Comment{}, nil
	} else {
		return Comment{}, result.Error
	}

}

func (r *CommentRepo) DeleteComment(id uint) {
	result := r.Db.Where("id = ?", id).Delete(&Comment{})
	panic(result.Error)
}

func (r *CommentRepo) UpdateComment(comment Comment) {

}
