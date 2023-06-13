package controller

import (
	"github.com/gin-gonic/gin"
	"hacktiv8_final/posts"
	"net/http"
	"strconv"
)

type CommentController struct {
	Service CommentService
}

func (cc *CommentController) CreateComment(c *gin.Context) {
	bodyReqComment := posts.CreateCommentRequest{}
	err := c.ShouldBind(&bodyReqComment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	finalBodyReqComment := posts.CreateCommentRequest{
		UserId:  uint(c.Keys["userId"].(int)),
		PhotoID: bodyReqComment.PhotoID,
		Message: bodyReqComment.Message,
	}
	err = cc.Service.CreateComment(finalBodyReqComment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Add",
	})
}

func (cc *CommentController) UpdateComment(c *gin.Context) {
	bodyReqUpdateComment := posts.UpdateCommentRequest{}
	err := c.ShouldBind(&bodyReqUpdateComment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	err = cc.Service.UpdateComment(bodyReqUpdateComment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Success",
	})

}

func (cc *CommentController) DeleteComment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	err = cc.Service.DeleteComment(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Success",
	})
}

func (cc *CommentController) FindByIdComment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	comment, err := cc.Service.FindByIdComment(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": comment,
	})
}

func (cc *CommentController) FindAllComment(c *gin.Context) {
	result := cc.Service.FindAllComment()
	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
