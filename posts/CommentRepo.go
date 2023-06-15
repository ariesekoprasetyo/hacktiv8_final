package posts

import (
	"hacktiv8_final/repository"
)

type RepositoryComment interface {
	SaveComment(Comment repository.Comment)
	UpdateComment(commentId uint, Comment repository.Comment)
	DeleteComment(CommentId uint)
	FindByIdComment(CommentId uint) (repository.Comment, error)
	FindAllComment() []repository.Comment
}
