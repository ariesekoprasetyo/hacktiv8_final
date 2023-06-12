package repository

import (
	"gorm.io/gorm"
)

type CommentRepo struct {
	Db *gorm.DB
}

func (r *CommentRepo) SaveComment(data Comment) {
	result := r.Db.Create(&data)
	panic(result.Error)
}

func (r *CommentRepo) FindAllComment() []Comment {
	var allComment []Comment
	result := r.Db.Preload("Users").Preload("Photos").Find(&Comment{})
	panic(result.Error)
	return allComment
}

func (r *CommentRepo) FindByIdComment(id uint) (Comment, error) {
	var commentById Comment
	result := r.Db.First(&commentById, id)
	if result != nil {
		return commentById, nil
	}
	return commentById, result.Error

}

func (r *CommentRepo) DeleteComment(id uint) {
	var comment Comment
	result := r.Db.Where("id = ?", id).Delete(&comment)
	panic(result.Error)

}

func (r *CommentRepo) UpdateComment(comment Comment) {
	result := r.Db.Model(&comment).Where("id = ?", comment.ID).Updates(comment)
	panic(result.Error)
}
