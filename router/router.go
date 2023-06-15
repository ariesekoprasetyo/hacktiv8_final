package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"hacktiv8_final/controller"
	"net/http"
	"time"
)

func CORSMiddleware(c *gin.Context) {
	cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Authorization", "Origin"},
		ExposeHeaders:    []string{"Content-Type", "Content-Length"},
		AllowCredentials: true,
		AllowWebSockets:  true,
		MaxAge:           12 * time.Hour,
	})
	c.Next()
}

func NewRouter(commentController *controller.CommentController, photoController *controller.PhotoController, authController *controller.AuthController, socialMediaController *controller.SocialMediaController) *gin.Engine {
	router := gin.Default()
	router.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, "welcome home")
	})
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	groupRouter := router.Group("api/v1/")
	groupRouter.Use(CORSMiddleware)
	{
		comment := groupRouter.Group("/posts/comment")
		comment.Use(controller.DeserializeUser(authController))
		{
			comment.POST("", commentController.CreateComment)
			comment.GET("", commentController.FindAllComment)
			comment.GET("/:id", commentController.FindByIdComment)
			comment.DELETE("/:id", commentController.DeleteComment)
			comment.PUT("/:id", commentController.UpdateComment)
		}
		photo := groupRouter.Group("/posts/photo")
		photo.Use(controller.DeserializeUser(authController))
		{
			photo.POST("", photoController.CreatePhoto)
			photo.GET("", photoController.FindAllPhoto)
			photo.GET("/:id", photoController.FindByIdPhoto)
			photo.DELETE("/:id", photoController.DeletePhoto)
			photo.PUT("/:id", photoController.UpdatePhoto)
		}
		socialmedia := groupRouter.Group("/users/socialmedia")
		socialmedia.Use(controller.DeserializeUser(authController))
		{
			socialmedia.POST("", socialMediaController.CreateSocialMedia)
			socialmedia.GET("", socialMediaController.FindAllSocialMedia)
			socialmedia.GET("/:id", socialMediaController.FindByIdSocialMedia)
			socialmedia.DELETE("/:id", socialMediaController.DeleteSocialMedia)
			socialmedia.PUT("/:id", socialMediaController.UpdateSocialMedia)
		}
		auth := groupRouter.Group("/users/authentication")
		{
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)
		}

	}
	return router
}
