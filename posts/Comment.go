package posts

import (
	"hacktiv8_final/controller"
	"hacktiv8_final/repository"
)

type CommentService struct {
	CommentRepo RepositoryComment
}

func (c *CommentService) UpdateComment(request controller.UpdateCommentRequest) {
	//TODO implement me
	panic("implement me")
}

func (c *CommentService) DeleteComment(commentId uint) {
	//TODO implement me
	panic("implement me")
}

func (c *CommentService) FindByIdComment(commentId uint) (repository.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CommentService) FindAllComment() []repository.Comment {
	//TODO implement me
	panic("implement me")
}

func (c *CommentService) CreateComment(request controller.CreateCommentRequest) {
	comment := repository.Comment{
		UserId:  request.UserId,
		PhotoID: request.UserId,
		Message: request.Message,
	}
	c.CommentRepo.SaveComment(comment)
}
