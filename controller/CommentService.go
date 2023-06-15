package controller

import (
	"hacktiv8_final/posts"
)

type CommentService interface {
	CreateComment(request posts.CreateCommentRequest) error
	UpdateCommentById(commentId uint, request posts.UpdateCommentRequest, userId uint) error
	DeleteCommentById(commentId uint, userId uint) error
	FindByIdComment(commentId uint) (posts.ResponCommnet, error)
	FindAllComment() []posts.ResponCommnet
}
