package controller

import (
	"hacktiv8_final/repository"
)

type PostsService interface {
	CreateComment(CreateCommentRequest)
	UpdateComment(UpdateCommentRequest)
	DeleteComment(commentId uint)
	FindByIdComment(commentId uint) (repository.Comment, error)
	FindAllComment() []repository.Comment
}
