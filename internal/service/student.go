package service

import (
	"easy-forum/internal/model"
	"time"
)

func CreatePost(uid int, content string) (err error) {
	err = d.CreatePost(ctx,&model.Post{
		UserID:  uid,
		Content: content,
		Time:   time.Now(),
	})
	return err
}


func GetPostList() (postList []model.Post, err error) {
	postList, err = d.GetPostList(ctx)
	return postList, err
}

func UpdatePost(uid int, pid int, content string) (err error) {
	err = d.UpdatePost(ctx, content,pid)
	return err
}

func DeletePost(uid int, pid int) (err error) {
	err = d.DeletePost(ctx, pid)
	return err
}

func ReportPost(uid int, pid int, reason string) (err error) {
	err = d.CreateReport(ctx,&model.Report{
		UserID: uid,
		PostID: pid,
		Reason: reason,
		Status: 0,
	})
	return err
}

func GetReportList(uid int)(reportList []model.Report, err error){
	reportList, err = d.GetReportList(ctx,uid)
	return reportList, err
}

func GetPostByID(pid int) (post *model.Post, err error) {
	post, err = d.GetPostByID(ctx, pid)
	return post, err
}