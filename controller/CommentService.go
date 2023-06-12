package controller

import (
	"hacktiv8_final/repository"
)

type CommentService interface {
	CreateComment(request CreateCommentRequest) error
	UpdateComment(request UpdateCommentRequest) error
	DeleteComment(commentId uint) error
	FindByIdComment(commentId uint) (repository.Comment, error)
	FindAllComment() []repository.Comment
}
