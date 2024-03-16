package services

import (
	"Glossika_interview/database/models"
	"gorm.io/gorm"
	"time"
)

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
