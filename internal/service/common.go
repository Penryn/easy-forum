package service

import (
	"context"
	"easy-forum/internal/dao"

	"gorm.io/gorm"
)

var (
	ctx =context.Background()
	d  *dao.Dao
)

func ServiceInit(db *gorm.DB){
	d=dao.New(db)
}