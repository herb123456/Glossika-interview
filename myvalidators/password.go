package myvalidators

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func passwordValidator(fl validator.FieldLevel) bool {
	passwd := fl.Field().String()

	// 驗證密碼長度
	if len(passwd) < 6 || len(passwd) > 16 {
		return false
	}

	// 驗證是否包含至少一個大寫字母、一個小寫字母和一個特殊字符
	var (
		uppercase = regexp.MustCompile(`[A-Z]`)
		lowercase = regexp.MustCompile(`[a-z]`)
		special   = regexp.MustCompile(`[()\[\]{}<>+\-*/?,.:;"'_\\|~!@#$%^&=]`)
	)

	return uppercase.MatchString(passwd) &&
		lowercase.MatchString(passwd) &&
		special.MatchString(passwd)
}
