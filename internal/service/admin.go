package service

import "easy-forum/internal/model"

func GetUserByID(id int) (user *model.User, err error) {
	user, err = d.GetUserByID(ctx, id)
	return user, err
}

func GetAdminReportList() (reports []model.Report, err error) {
	reports, err = d.GetAdminReportList(ctx)
	return reports, err
}

func ApproveReport(pid int,approval int) (err error) {
	err = d.ApproveReport(ctx, pid, approval)
	return err
}