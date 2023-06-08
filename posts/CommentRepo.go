package posts

import "hacktiv8_final/repository"

type RepositoryComment interface {
	Save(Comment repository.Comment)
	Update(Comment repository.Comment)
	Delete(CommentId int)
	FindById(CommentId int) (repository.Comment, error)
	FindAll() []repository.Comment
	FindByUserName(username string) (repository.Comment, error)
}
