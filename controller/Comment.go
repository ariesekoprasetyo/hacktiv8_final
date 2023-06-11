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
	ID      uint   `json:"id" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type CommentController struct {
	Service CommentService
}

func (cc *CommentController) CreateComment(c *gin.Context) {
	bodyReqComment := CreateCommentRequest{}
	err := c.ShouldBind(&bodyReqComment)
	if err != nil {
		panic(err)
	}
	cc.Service.CreateComment(bodyReqComment)
}

func (cc *CommentController) UpdateComment(c *gin.Context) {
	bodyReqUpdateComment := UpdateCommentRequest{}
	err := c.ShouldBind(&bodyReqUpdateComment)
	if err != nil {
		panic(err)
	}
	cc.Service.UpdateComment(bodyReqUpdateComment)
}

func (cc *CommentController) DeleteComment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	cc.Service.DeleteComment(uint(id))

}

func (cc *CommentController) FindByIdComment(c *gin.Context) {
	//id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	//if err != nil {
	//	panic(err)
	//}
	//comment, err := cc.Service.FindByIdComment(uint(id))
	//if err != nil {
	//	return
	//}
	//gin.H{"coco":comment.UserId}
}

func (cc *CommentController) FindAllComment(c *gin.Context) {
	cc.Service.FindAllComment()
}
