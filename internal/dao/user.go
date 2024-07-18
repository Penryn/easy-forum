package dao

import (
	"context"
	"easy-forum/internal/model"
)

func (d *Dao) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := d.orm.WithContext(ctx).Where("username = ?", username).First(&user).Error
	return &user, err
}

func (d *Dao) CreateUser(ctx context.Context, user *model.User) error {
	err := d.orm.WithContext(ctx).Create(&user).Error
	return err
}

func (d *Dao)GetUserByID(ctx context.Context, uid int) (*model.User, error) {
	var user model.User
	err := d.orm.WithContext(ctx).Where("id = ?", uid).First(&user).Error
	return &user, err
}