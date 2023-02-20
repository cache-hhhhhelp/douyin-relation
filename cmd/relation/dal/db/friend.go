package db

import (
	"context"
	"douyin/pkg/constants"
	"gorm.io/gorm"
)

type Friend struct {
	gorm.Model
	FromUser int64 `json:from_user`
	ToUSer   int64 `json:to_user`
}

func (r *Friend) TableName() string {
	return constants.FRIENDTABLENAME
}

func MGetFriend(ctx context.Context, userId int) ([]int, error) {
	var res []int
	err := DB.WithContext(ctx).Where("from_user = ?", userId).Select("to_user").Find(&res).Error
	return res, err
}
