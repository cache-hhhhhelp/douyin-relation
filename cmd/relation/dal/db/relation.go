package db

import (
	"context"
	"douyin/pkg/constants"
	"gorm.io/gorm"
)

type Relation struct {
	gorm.Model
	FromUser int64 `json:from_user`
	ToUSer   int64 `json:to_user`
}

func (r *Relation) TableName() string {
	return constants.FollowTableName
}

func CreateRelation(ctx context.Context, relations *Relation) error {
	return DB.WithContext(ctx).Create(relations).Error
}

func MGetFollows(ctx context.Context, userId int) ([]int, error) {
	var res []int
	err := DB.WithContext(ctx).Where("from_user = ?", userId).Select("to_user").Find(&res).Error
	return res, err
}

func MGetFollowers(ctx context.Context, userId int) ([]int, error) {
	var res []int
	err := DB.WithContext(ctx).Where("to_user = ?", userId).Select("from_user").Find(&res).Error
	return res, err
}

func DeleteRelation(ctx context.Context, from_user, to_user int) error {
	return DB.WithContext(ctx).Where("from_user = ? and to_user = ? ", from_user, to_user).Delete(&Relation{}).Error
}
