package controller

import (
	"Glossika_interview/myerror"
	"Glossika_interview/response"
	"Glossika_interview/services"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u UsersController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		e := myerror.NewValidationError(err)
		response.Error(c, e)
		return
	}

	userService := services.UserService{DB: c.MustGet("db").(*gorm.DB)}
	user, err := userService.Login(req.Email, req.Password)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, myerror.NewUserError(myerror.TYPE_EMAIL_PASSWORD_WRONG))
			return
		} else {
			response.Error(c, myerror.NewGeneralError(err.Error()))
			return
		}
	}
	if !user.Verified {
		response.Error(c, myerror.NewUserError(myerror.TYPE_EMAIL_NOT_VERIFIED))
		return
	}

	// generate jwt token
	token, err := userService.GenerateToken(user)
	if err != nil {
		response.Error(c, myerror.NewGeneralError(err.Error()))
		return
	}

	// 返回成功
	response.Success(c, gin.H{"token": token})
}
