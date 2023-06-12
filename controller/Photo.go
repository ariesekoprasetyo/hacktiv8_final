package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CreatePhotoRequest struct {
	Title    string `validate:"required" json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `validate:"required" json:"photo_url"`
	UserID   uint   `json:"user_id"`
}
type UpdatePhotoRequest struct {
	ID       uint   `json:"id"`
	Title    string `validate:"required" json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `validate:"required" json:"photo_url"`
}

type PhotoController struct {
	Service PhotoService
}

func (pc *PhotoController) CreatePhoto(c *gin.Context) {
	bodyReqPhoto := CreatePhotoRequest{}
	err := c.ShouldBind(&bodyReqPhoto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	finalBodyReqPhoto := CreatePhotoRequest{
		Title:    bodyReqPhoto.Title,
		Caption:  bodyReqPhoto.Caption,
		PhotoURL: bodyReqPhoto.PhotoURL,
		UserID:   uint(c.Keys["userId"].(int)),
	}
	err = pc.Service.CreatePhoto(finalBodyReqPhoto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Success Add",
	})
}

func (pc *PhotoController) UpdatePhoto(c *gin.Context) {
	bodyReqUpdatePhoto := UpdatePhotoRequest{}
	err := c.ShouldBind(&bodyReqUpdatePhoto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
}

func (pc *PhotoController) DeletePhoto(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	err = pc.Service.DeletePhoto(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Success",
	})
}

func (pc *PhotoController) FindByIdPhoto(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	photo, err := pc.Service.FindByIdPhoto(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
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
