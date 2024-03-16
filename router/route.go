package router

import (
	"Glossika_interview/controller"
	"Glossika_interview/middleware"
	"github.com/gin-gonic/gin"
)

func Route(g *gin.Engine) {
	g.POST("/register", controller.UsersController{}.Register)
	g.POST("/verify-email", controller.UsersController{}.VerifyEmail)
	g.POST("/login", controller.UsersController{}.Login)

	g.GET("/recommendation", middleware.JWTAuthMiddleware(), controller.RecommendationController{}.GetRecommendation)
}
