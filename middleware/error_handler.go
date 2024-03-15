package middleware

import (
	"Glossika_interview/response"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next() // 調用後續的處理函數

	// 如果有錯誤發生，則處理它們
	if len(c.Errors) > 0 {
		response.Error(c, c.Errors.Last())
	}
}
