package dao

import (
	"context"
	"easy-forum/internal/model"
)

func (d *Dao) CreatePost(ctx context.Context, post *model.Post) (err error) {
	err = d.orm.WithContext(ctx).Create(&post).Error
	return err
}

func (d *Dao) GetPostList(ctx context.Context) (postList []model.Post, err error) {
	err = d.orm.WithContext(ctx).Find(&postList).Error
	return postList, err
}

func(d *Dao) UpdatePost(ctx context.Context, content string, postid int) (err error) {
	err = d.orm.WithContext(ctx).Model(&model.Post{}).Where("id = ?", postid).Update("content", content).Error
	return err
}

func(d *Dao) DeletePost(ctx context.Context, postid int) (err error) {
	err = d.orm.WithContext(ctx).Delete(&model.Post{}, postid).Error
	return err
}

func (d *Dao)GetPostByID(ctx context.Context, postid int) (post *model.Post, err error) {
	post = new(model.Post)
	err = d.orm.WithContext(ctx).Where("id = ?", postid).First(post).Error
	return post, err
}

func (d *Dao)GetAllPostByID(ctx context.Context, postid int) (post *model.Post, err error) {
	post = new(model.Post)
	err = d.orm.WithContext(ctx).Unscoped().Where("id = ?", postid).First(post).Error
	return post, err
}