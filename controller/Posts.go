package controller

import "github.com/gin-gonic/gin"

type CreateCommentRequest struct {
	UserId  uint   `json:"user_id" binding:"required"`
	PhotoID uint   `json:"photo_id" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type UpdateCommentRequest struct {
	Message string `json:"message" binding:"required"`
}

type PostController struct {
	PostsControl Posts
}

func (controller *PostController) CreateComment(c *gin.Context) {
	bodyReqComment := CreateCommentRequest{}
	err := c.ShouldBind(bodyReqComment)
	if err != nil {
		panic(err)
	}
	controller.PostsControl.CreateComment(bodyReqComment)
}
