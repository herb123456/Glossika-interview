package controller

import (
	"Glossika_interview/myerror"
	"Glossika_interview/response"
	"Glossika_interview/services"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type VerifyEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required,len=4"`
}

func (u UsersController) VerifyEmail(c *gin.Context) {
	var req VerifyEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		e := myerror.NewValidationError(err)
		response.Error(c, e)
		return
	}

	userService := services.UserService{DB: c.MustGet("db").(*gorm.DB)}
	user, err := userService.VerifyEmail(req.Email, req.Code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, myerror.NewUserError(myerror.TYPE_EMAIL_CODE_WRONG))
			return
		} else {
			response.Error(c, myerror.NewGeneralError(err.Error()))
			return
		}
	}
	if !user.Verified {
		response.Error(c, myerror.NewUserError(myerror.TYPE_EMAIL_CODE_WRONG))
		return
	}

	// 返回成功
	response.Success(c, nil)
}
