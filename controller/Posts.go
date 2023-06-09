package controller

import (
	"hacktiv8_final/repository"
)

type Posts interface {
	CreateComment(CreateCommentRequest)
	UpdateComment(UpdateCommentRequest)
	DeleteComment(commentId int)
	FindByIdComment(commentId int) (repository.Comment, error)
	FindAllComment() []repository.Comment
}
