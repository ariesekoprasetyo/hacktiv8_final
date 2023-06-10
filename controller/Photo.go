package controller

import "github.com/gin-gonic/gin"

type CreatePhotoRequest struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
}
type UpdatePhotoRequest struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
}

type PhotoController struct {
	Service PhotoService
}

func (pc *PhotoController) CreatePhoto(c *gin.Context) {
	bodyReqPhoto := CreatePhotoRequest{}
	err := c.ShouldBind(&bodyReqPhoto)
	if err != nil {
		panic(err)
	}
	pc.Service.CreatePhoto(bodyReqPhoto)
}
