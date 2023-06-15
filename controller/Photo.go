package controller

import (
	"github.com/gin-gonic/gin"
	"hacktiv8_final/posts"
	"net/http"
	"strconv"
)

type PhotoController struct {
	Service PhotoService
}

func (pc *PhotoController) CreatePhoto(c *gin.Context) {
	bodyReqPhoto := posts.CreatePhotoRequest{}

	err := c.ShouldBind(&bodyReqPhoto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	finalBodyReqPhoto := posts.CreatePhotoRequest{
		Title:    bodyReqPhoto.Title,
		Caption:  bodyReqPhoto.Caption,
		PhotoURL: bodyReqPhoto.PhotoURL,
		UserID:   uint(c.Keys["userId"].(uint)),
	}
	err = pc.Service.CreatePhoto(finalBodyReqPhoto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Success Add",
	})
}

func (pc *PhotoController) UpdatePhoto(c *gin.Context) {
	bodyReqUpdatePhoto := posts.UpdatePhotoRequest{}
	err := c.ShouldBind(&bodyReqUpdatePhoto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = pc.Service.UpdatePhoto(uint(id), bodyReqUpdatePhoto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Success Update Photo",
	})
}

func (pc *PhotoController) DeletePhoto(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	err = pc.Service.DeletePhotoById(uint(id), uint(c.Keys["userId"].(uint)))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Success",
	})
}

func (pc *PhotoController) FindByIdPhoto(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	photo, err := pc.Service.FindByIdPhoto(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": photo,
	})
}

func (pc *PhotoController) FindAllPhoto(c *gin.Context) {
	result := pc.Service.FindAllPhoto()
	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
