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
	err = d.orm.WithContext(ctx).Unscoped().Where("user_id = ?", uid).Find(&reportList).Error
	return reportList, err
}

func (d *Dao) GetAdminReportList(ctx context.Context) (reportList []model.Report, err error) {
	err = d.orm.WithContext(ctx).Unscoped().Where("status = ?", 0).Find(&reportList).Error
	return reportList, err
}

func (d *Dao) ApproveReport(ctx context.Context, pid int, approval int) (err error) {
	err = d.orm.WithContext(ctx).Model(&model.Report{}).Where("post_id = ?", pid).Update("status", approval).Error
	if err != nil {
		return err
	}
	if approval == 1 {
		err = d.orm.WithContext(ctx).Where("id = ?", pid).Delete(&model.Post{}).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Dao) GetReportByPostID(ctx context.Context, pid int) (report *model.Report, err error) {
	report = new(model.Report)
	err = d.orm.WithContext(ctx).Where("post_id = ?", pid).First(&report).Error
	return report, err
}

func (d *Dao)DeleteReport(ctx context.Context, rid int) (err error) {
	err = d.orm.WithContext(ctx).Delete(&model.Report{}, rid).Error
	return err
}

func (d *Dao)GetReportByPostIDAndUserID(ctx context.Context, uid int, pid int) (report *model.Report, err error) {
	report = new(model.Report)
	err = d.orm.WithContext(ctx).Where("user_id = ? AND post_id = ?", uid, pid).First(&report).Error
	return report, err
}