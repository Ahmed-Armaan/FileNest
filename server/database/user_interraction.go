package database

import (
	"github.com/Ahmed-Armaan/FileNest/database/helper"
	"gorm.io/gorm"
)

func InsertUser(userName string, googleID string, email string, profileImage string) (*User, error) {
	//	if _, err := GetUserByGoogleID(googleID); err == nil {
	//		return nil
	//	} else if err != gorm.ErrRecordNotFound {
	//		return err
	//	}

	user := &User{
		UserName:     userName,
		GoogleID:     googleID,
		Email:        email,
		ProfileImage: profileImage,
	}

	//if err := DB.Create(&user).Error; err != nil {
	//	return false, err, uuid.UUID{}
	//}
	//return true, nil, user.ID

	// insert new user and create a root node, using transaction to achieve atomicity
	if err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			if helper.ResolvePostgresError(err) == helper.ErrUniqueViolation {
				return nil
			}
			return err
		}

		if err := insertRootNode(tx, user.ID); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return user, err
	}

	return user, nil
}

func GetUserByGoogleID(googleID string) (*User, error) {
	var user User

	err := DB.Where("google_id = ?", googleID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
