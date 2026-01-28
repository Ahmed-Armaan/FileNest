package database

import (
	"gorm.io/gorm"
)

var UserDbColums = struct {
	ID           string
	GoogleID     string
	UserName     string
	Email        string
	ProfileImage string
}{
	ID:           "id",
	GoogleID:     "google_id",
	UserName:     "user_name",
	Email:        "email",
	ProfileImage: "profile_image",
}

func GetUserIdByGoogleIdSubQuery(googleId string, columns ...string) *gorm.DB {
	return DB.Model(&User{}).Select("id").Where("google_id = ?", googleId)
}

func GetUserDataByGoogleId(googleId string, columns ...string) (*User, error) {
	user := User{}
	if err := DB.Model(&User{}).
		Select(columns).
		Where("google_id = ?", googleId).
		Take(&user).Error; err != nil {
		return &user, err
	}

	return &user, nil
}
