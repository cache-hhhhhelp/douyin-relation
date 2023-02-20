package db

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string
	FollowCount   int
	FollowerCount int
}

func (r *User) TableName() string {
	return "User"
}

func GetUser(ctx context.Context, userId int) (User, error) {
	var res User
	err := DB.WithContext(ctx).Where("id = ?", userId).Find(&res).Error
	return res, err
}
