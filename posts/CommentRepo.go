package posts

import "hacktiv8_final/repository"

type RepositoryComment interface {
	SaveComment(Comment repository.Comment)
	UpdateComment(Comment repository.Comment)
	DeleteComment(CommentId int)
	FindByIdComment(CommentId int) (repository.Comment, error)
	FindAllComment() []repository.Comment
}
