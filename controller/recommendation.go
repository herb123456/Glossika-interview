package controller

import (
	"Glossika_interview/database/models"
	"github.com/gin-gonic/gin"
)

type RecommendationController struct{}

func (rc RecommendationController) GetRecommendation(c *gin.Context) {
	user := c.MustGet("user").(models.User)
}
