package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateUsersRequest struct {
	Username string `validate:"required" json:"username"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required,email" json:"password"`
	Age      int    `validate:"required,min=8" json:"age"`
}

type LoginUsersRequest struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
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
			"message": err,
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
			"message": err,
		})
	}
	err = ac.Service.Register(bodyReqRegister)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "Ok",
		"message": "Successfully Created user !",
		"data":    nil,
	})
}
