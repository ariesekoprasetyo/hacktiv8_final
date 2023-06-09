package posts

import (
	"hacktiv8_final/controller"
	"hacktiv8_final/repository"
)

type CommentController struct {
	CommentRepo RepositoryComment
}

func (c *CommentController) CreateComment(request controller.CreateCommentRequest) {
	comment := repository.Comment{
		UserId:  request.UserId,
		PhotoID: request.UserId,
		Message: request.Message,
	}
	c.CommentRepo.Save(comment)
}
