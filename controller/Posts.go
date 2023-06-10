package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type CreateCommentRequest struct {
	UserId  uint   `json:"user_id" binding:"required"`
	PhotoID uint   `json:"photo_id" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type UpdateCommentRequest struct {
	Message string `json:"message" binding:"required"`
}

type PostsController struct {
	PostsControl PostsService
}

func (controller *PostsController) CreateComment(c *gin.Context) {
	bodyReqComment := CreateCommentRequest{}
	err := c.ShouldBind(&bodyReqComment)
	if err != nil {
		panic(err)
	}
	controller.PostsControl.CreateComment(bodyReqComment)
}

func (controller *PostsController) UpdateComment(c *gin.Context) {
	bodyReqUpdateComment := UpdateCommentRequest{}
	err := c.ShouldBind(&bodyReqUpdateComment)
	if err != nil {
		panic(err)
	}
	controller.PostsControl.UpdateComment(bodyReqUpdateComment)
}

func (controller *PostsController) DeleteComment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	controller.PostsControl.DeleteComment(uint(id))

}

func (controller *PostsController) FindByIdComment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	controller.PostsControl.FindByIdComment(uint(id))
}

func (controller *PostsController) FindAllComment(c *gin.Context) {
	controller.PostsControl.FindAllComment()
}
