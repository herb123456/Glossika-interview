package services

import (
	"Glossika_interview/config"
	"Glossika_interview/database/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

// key 可以帶在環境變數中
var jwtKey = []byte(config.GetString("jwt.secret"))

type UserService struct {
	DB *gorm.DB
}

func (us *UserService) VerifyEmail(email, code string) (*models.User, error) {
	// query the email and code from the database
	var user models.User
	res := us.DB.Where("email = ? AND verification_code = ?", email, code).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	// update the user's verified status
	user.Verified = true
	user.VerificationAt = time.Now()
	res = us.DB.Save(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

func (us *UserService) Login(email, password string) (*models.User, error) {
	// query the email and password from the database
	var user models.User
	res := us.DB.Where("email = ?", email).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	// compare the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return &user, nil
}

func (us *UserService) GenerateToken(u *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// 设置声明
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = u.ID
	claims["email"] = u.Email
	claims["verified"] = u.Verified
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
