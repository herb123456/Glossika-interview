package response

import (
	"Glossika_interview/config"
	"Glossika_interview/myerror"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
)

// Response 是一个用于API响应的通用结构体
type Response struct {
	Code      int         `json:"code"` // 状态码
	ErrorCode string      `json:"errorCode,omitempty"`
	Message   string      `json:"message"` // 消息
	Data      interface{} `json:"data"`    // 返回数据
	Debug     interface{} `json:"debug,omitempty"`
}

// Success 用于返回成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}

// Error 用于返回错误响应
func Error(c *gin.Context, e error) {
	dres := Response{
		Code:    500,
		Message: "Internal Server Error",
		Data:    nil,
		Debug:   nil,
	}

	isDebugConfig := config.Get("app.debug")
	isDebug, ok := isDebugConfig.(bool)
	if !ok {
		isDebug = false
	}
	if isDebug {
		dres.Debug = debug.Stack()
	}

	// 取最後一個錯誤進行處理
	var appErr *myerror.AppError
	ok = errors.As(e, &appErr)
	if ok {
		// 使用自定義的錯誤訊息回傳給客戶端
		dres.Code = appErr.Code
		dres.ErrorCode = appErr.ErrorCode
		dres.Message = appErr.Message
		if isDebug {
			dres.Debug = appErr.DebugMessage
		}
	}

	c.JSON(http.StatusOK, dres)
	return
}
