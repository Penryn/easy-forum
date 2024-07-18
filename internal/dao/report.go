package dao

import (
	"context"
	"easy-forum/internal/model"
)

func (d *Dao) CreateReport(ctx context.Context, report *model.Report) (err error) {
	err = d.orm.WithContext(ctx).Create(&report).Error
	return err
}

func (d *Dao) GetReportList(ctx context.Context, uid int) (reportList []model.Report, err error) {
	err = d.orm.WithContext(ctx).Where("user_id = ?", uid).Find(&reportList).Error
	return reportList, err
}

func (d *Dao) GetAdminReportList(ctx context.Context) (reportList []model.Report, err error) {
	err = d.orm.WithContext(ctx).Where("status = ?", 0).Find(&reportList).Error
	return reportList, err
}

func (d *Dao) ApproveReport(ctx context.Context, pid int, approval int, uid int) (err error) {
	var report model.Report
	var post model.Post
	err = d.orm.WithContext(ctx).Where("post_id = ?", pid).Model(&report).Update("status", approval).Error
	if err != nil {
		return err
	}
	if approval == 1 {
		err = d.orm.WithContext(ctx).Where("id = ?", pid).Delete(&post).Error
		if err != nil {
			return err
		}
	}
	return nil
}
