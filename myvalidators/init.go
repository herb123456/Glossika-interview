package myvalidators

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log/slog"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 註冊自定義驗證函數
		err := v.RegisterValidation("passwd", passwordValidator)
		if err != nil {
			slog.Error("RegisterValidation error: \n" + err.Error())
			panic(err)
		}
	}
}
