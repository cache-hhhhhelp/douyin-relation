package main

import (
	"context"
	"douyin/cmd/relation/dal/db"
	relation "douyin/kitex_gen/relation"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

func getUserId(token *string) int64 {
	// TODO
	return 0
}
func getUser(ctx context.Context, useId int) (rUser *relation.User) {
	user, _ := db.GetUser(ctx, useId)
	var id int64
	id = int64(user.ID)
	var followCount int64
	var followerCount int64
	followCount = int64(user.FollowCount)
	followerCount = int64(user.FollowerCount)
	var isFollow bool
	isFollow = true
	rUser = &relation.User{
		Id:            &id,
		Name:          &(user.Name),
		FollowCount:   &followCount,
		FollowerCount: &followerCount,
		IsFollow:      &isFollow,
	}
	return
}

// Follow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) Follow(ctx context.Context, req *relation.DouyinRelationActionRequest) (resp *relation.DouyinRelationActionResponse, err error) {
	resp = new(relation.DouyinRelationActionResponse)
	relationModel := &db.Relation{
		FromUser: getUserId(req.Token),
		ToUSer:   *req.ToUserId,
	}
	if *req.ActionType == 1 {
		err = db.CreateRelation(ctx, relationModel)
	} else if *req.ActionType == 2 {
		err = db.DeleteRelation(ctx, int(getUserId(req.Token)), int(*req.ToUserId))
	}
	var code int32 = 0
	var msg string = "Success"
	if err != nil {
		code = 1
		msg = "Error"
	}
	resp.StatusMsg = &msg
	resp.StatusCode = &code
	return
}

// ListFollow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) ListFollow(ctx context.Context, req *relation.DouyinRelationFollowListRequest) (resp *relation.DouyinRelationFollowListResponse, err error) {
	var followIds []int
	followIds, err = db.MGetFollows(ctx, int(*req.UserId))
	var code int32 = 0
	var msg string = "Success"
	if err != nil {
		code = 1
		msg = "Error"
	}
	resp.StatusMsg = &msg
	resp.StatusCode = &code
	for _, followId := range followIds {
		user := getUser(ctx, followId)
		resp.UserList = append(resp.UserList, user)
	}
	return
}

// ListFollower implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) ListFollower(ctx context.Context, req *relation.DouyinRelationFollowerListRequest) (resp *relation.DouyinRelationFollowerListResponse, err error) {
	var followerIds []int
	followerIds, err = db.MGetFollowers(ctx, int(*req.UserId))
	var code int32 = 0
	var msg string = "Success"
	if err != nil {
		code = 1
		msg = "Error"
	}
	resp.StatusMsg = &msg
	resp.StatusCode = &code
	for _, followId := range followerIds {
		user := getUser(ctx, followId)
		resp.UserList = append(resp.UserList, user)
	}
	return
}
