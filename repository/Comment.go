package repository

import (
	"gorm.io/gorm"
)

type CommentRepo struct {
	Db *gorm.DB
}

func (r *CommentRepo) SaveComment(data Comment) {
	err := r.Db.Create(&data).Error
	if err != nil {
		panic(err)
	}
}

func (r *CommentRepo) FindAllComment() []Comment {
	var allComment []Comment
	err := r.Db.Preload("User").Preload("Photo").Find(&allComment).Error
	if err != nil {
		panic(err)
	}
	return allComment
}

func (r *CommentRepo) FindByIdComment(id uint) (Comment, error) {
	var commentById Comment
	err := r.Db.Preload("User").Preload("Photo").First(&commentById, id).Error
	if err != nil {
		return commentById, err
	}
	return commentById, nil

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
