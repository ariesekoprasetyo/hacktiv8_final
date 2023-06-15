package controller

import (
	"github.com/gin-gonic/gin"
	"hacktiv8_final/user"
	"net/http"
	"strconv"
)

type SocialMediaController struct {
	Service SocialMediaService
}

func (sm *SocialMediaController) CreateSocialMedia(c *gin.Context) {
	bodyReqSocialMedia := user.SocialMediaRequest{}
	err := c.ShouldBind(&bodyReqSocialMedia)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	finalBodyReqSocialMedia := user.SocialMediaRequest{
		Name:      bodyReqSocialMedia.Name,
		SosmedUrl: bodyReqSocialMedia.SosmedUrl,
		UserId:    uint(c.Keys["userId"].(uint)),
	}
	err = sm.Service.CreateSocialMedia(finalBodyReqSocialMedia)
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

func (sm *SocialMediaController) UpdateSocialMediaById(c *gin.Context) {
	bodyReqUpdateSocialMedia := user.SocialMediaUpdate{}
	err := c.ShouldBind(bodyReqUpdateSocialMedia)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
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
	err = sm.Service.UpdateSocialMediaById(uint(id), bodyReqUpdateSocialMedia, uint(c.Keys["userId"].(uint)))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Success",
	})
}

func (sm *SocialMediaController) DeleteSocialMediaById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = sm.Service.DeleteSocialMediaById(uint(id), uint(c.Keys["userId"].(uint)))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Success",
	})
}

func (sm *SocialMediaController) FindByIdSocialMedia(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	socialmedia, err := sm.Service.FindByIdSocialMedia(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": socialmedia,
	})
}

func (sm *SocialMediaController) FindAllSocialMedia(c *gin.Context) {
	result := sm.Service.FindAllSocialMedia()
	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
