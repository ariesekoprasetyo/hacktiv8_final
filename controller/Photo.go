package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreatePhotoRequest struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
}
type UpdatePhotoRequest struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
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
