package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func DeserializeUser(controller *AuthController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.TrimPrefix(authorizationHeader, "Bearer ")

		if len(fields) != 0 {
			token = strings.Replace(fields, "Bearer ", "", 1)
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		sub, err := controller.Service.ValidateToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail 1", "message": err.Error()})
			return
		}

		id, errId := strconv.ParseUint(fmt.Sprint(sub), 10, 64)
		if errId != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail 2", "message": err.Error()})
			return
		}
		result, err := controller.Service.FindUserById(uint(id))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail 3", "message": "the user belonging to this token no logger exists"})
			return
		}

		ctx.Set("userId", result)
		ctx.Next()

	}
}
