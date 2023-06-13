package controller

import (
	"hacktiv8_final/posts"
	"hacktiv8_final/repository"
)

type CommentService interface {
	CreateComment(request posts.CreateCommentRequest) error
	UpdateComment(request posts.UpdateCommentRequest) error
	DeleteComment(commentId uint) error
	FindByIdComment(commentId uint) (repository.Comment, error)
	FindAllComment() []repository.Comment
}
