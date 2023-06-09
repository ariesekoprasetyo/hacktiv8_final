package posts

import (
	"hacktiv8_final/controller"
	"hacktiv8_final/repository"
)

type CommentController struct {
	CommentRepo RepositoryComment
}

// UpdateComment
func (c *CommentController) UpdateComment(request controller.UpdateCommentRequest) {
	//TODO implement me
	panic("implement me")
}

// DeleteComment
func (c *CommentController) DeleteComment(commentId int) {
	//TODO implement me
	panic("implement me")
}

func (c *CommentController) FindByIdComment(commentId int) (repository.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CommentController) FindAllComment() []repository.Comment {
	//TODO implement me
	panic("implement me")
}

func (c *CommentController) CreateComment(request controller.CreateCommentRequest) {
	comment := repository.Comment{
		UserId:  request.UserId,
		PhotoID: request.UserId,
		Message: request.Message,
	}
	c.CommentRepo.SaveComment(comment)
}
