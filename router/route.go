package router

import (
	"Glossika_interview/controller"
	"github.com/gin-gonic/gin"
)

func Route(g *gin.Engine) {
	g.POST("/register", controller.UsersController{}.Register)
	g.POST("/verify-email", controller.UsersController{}.VerifyEmail)
}
