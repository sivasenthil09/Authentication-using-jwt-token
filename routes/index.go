package routes

import (
	"authentication/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, controller controllers.UserController) {

	router.POST("/signup", controller.Signup)
	router.GET("/login", controller.Login)
	router.GET("/home", controller.Home)
	// router.PATCH("/refresh", controller.refresh)
}

func Default(router *gin.Engine) {
	router.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Server is Healthy"})
	})
}
