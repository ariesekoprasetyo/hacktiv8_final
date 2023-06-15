package controller

import (
	"hacktiv8_final/posts"
)

type CommentService interface {
	CreateComment(request posts.CreateCommentRequest) error
	UpdateComment(commentId uint, request posts.UpdateCommentRequest) error
	DeleteCommentById(commentId uint, userId uint) error
	FindByIdComment(commentId uint) (posts.ResponCommnet, error)
	FindAllComment() []posts.ResponCommnet
}
