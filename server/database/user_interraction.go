package database

import (
	"gorm.io/gorm"
)

func InsertUser(userName string, googleID string, email string, profileImage string) error {
	if _, err := GetUserByGoogleID(googleID); err == nil {
		return nil
	} else if err != gorm.ErrRecordNotFound {
		return err
	}

	user := &User{
		UserName:     userName,
		GoogleID:     googleID,
		Email:        email,
		ProfileImage: profileImage,
	}

	if err := DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByGoogleID(googleID string) (*User, error) {
	var user User

	err := DB.Where("google_id = ?", googleID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
