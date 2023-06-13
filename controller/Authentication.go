package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateUsersRequest struct {
	Username string `validate:"required" json:"username"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
	Age      int    `validate:"required,gte=8" json:"age"`
}

type LoginUsersRequest struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}

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
	bodyReqLogin := LoginUsersRequest{}
	err := c.ShouldBind(&bodyReqLogin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	token, err := ac.Service.Login(bodyReqLogin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":     "Ok",
		"message":    "Successfully log in",
		"token_type": "Bearer",
		"token":      token,
	})
}

func (ac *AuthController) Register(c *gin.Context) {
	bodyReqRegister := CreateUsersRequest{}
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
	c.JSON(http.StatusOK, gin.H{
		"Status":  "Ok",
		"Message": "Successfully Created user !",
		"Data":    nil,
	})
}
