package controller

import (
	"Glossika_interview/database/models"
	"Glossika_interview/myerror"
	"Glossika_interview/response"
	"Glossika_interview/services"
	"Glossika_interview/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type UsersController struct{}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,passwd"`
}

func (u UsersController) Register(c *gin.Context) {
	var req RegisterRequest
	// validate the request
	if err := c.ShouldBindJSON(&req); err != nil {
		e := myerror.NewValidationError(err)
		response.Error(c, e)

		return
	}
	db := c.MustGet("db").(*gorm.DB)

	// check if the email is already registered
	var user models.User
	res := db.Where("email = ?", req.Email).First(&user)
	if user.ID != 0 {
		e := myerror.NewUserError(myerror.TYPE_USER_EXIST)
		response.Error(c, e)

		return
	}

	// hash the password
	bytesPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 7)
	if err != nil {
		e := myerror.NewGeneralError(err.Error())
		response.Error(c, e)

		return
	}
	hashedPasswd := string(bytesPassword)

	// generate a random 4 digits verification code
	verificationCode := fmt.Sprintf("%0*d", 4, utils.RandomDigits(4))

	// create a user
	user = models.User{
		Email:                  req.Email,
		Password:               hashedPasswd,
		VerificationCode:       verificationCode,
		VerificationCodeExpiry: time.Now().Add(20 * time.Minute),
	}
	res = db.Create(&user)
	if res.Error != nil {
		e := myerror.NewGeneralError(res.Error.Error())
		response.Error(c, e)

		return
	}

	// send a verification email
	go services.SendEmail(user.Email, "Verification Code", "Your verification code is "+verificationCode)

	response.Success(c, user)
}
