package controller

import (
	"hacktiv8_final/repository"
)

type CommentService interface {
	CreateComment(request CreateCommentRequest)
	UpdateComment(request UpdateCommentRequest)
	DeleteComment(commentId uint)
	FindByIdComment(commentId uint) (repository.Comment, error)
	FindAllComment() []repository.Comment
}
