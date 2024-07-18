package dao

import (
	"context"
	"easy-forum/internal/model"

	"gorm.io/gorm"
)

type Dao struct {
	orm *gorm.DB
}

func New(orm *gorm.DB) *Dao {
	return &Dao{orm: orm}
}

type Daos interface{
	// User
	GetUserByUsername(ctx context.Context,username string) (user *model.User, err error)
	CreateUser(ctx context.Context,user *model.User) (err error)
	GetUserByID(ctx context.Context,uid int) (user *model.User, err error)
	// Post
	CreatePost(ctx context.Context,post *model.Post) (err error)
	GetPostList(ctx context.Context) (postList []*model.Post, err error)
	UpdatePost(ctx context.Context,content string,postid int) (err error)
	DeletePost(ctx context.Context,postid int) (err error)
	// Report
	CreateReport(ctx context.Context,report *model.Report) (err error)
	GetReportList(ctx context.Context,uid int) (reportList []*model.Report, err error)
	GetAdminReportList(ctx context.Context) (reportList []*model.Report, err error)
	ApproveReport(ctx context.Context,rid int,approval int,uid int) (err error)
}