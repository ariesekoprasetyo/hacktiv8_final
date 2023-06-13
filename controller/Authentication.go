package controller

import (
	"github.com/gin-gonic/gin"
	"hacktiv8_final/user"
	"log"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type AuthController struct {
	Service AuthService
}

func (ac *AuthController) Login(c *gin.Context) {
	bodyReqLogin := user.LoginUsersRequest{}
	err := c.ShouldBind(&bodyReqLogin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	token, err := ac.Service.Login(bodyReqLogin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":     "Ok",
		"message":    "Successfully log in",
		"token_type": "Bearer",
		"token":      token,
	})
}

func (ac *AuthController) Register(c *gin.Context) {
	bodyReqRegister := user.CreateUsersRequest{}
	err := c.ShouldBind(&bodyReqRegister)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = ac.Service.Register(bodyReqRegister)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	log.Print("Masuk Sini")
	c.JSON(http.StatusOK, gin.H{
		"Status":  "Ok",
		"Message": "Successfully Created user !",
		"Data":    nil,
	})
}
