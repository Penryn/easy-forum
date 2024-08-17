package admin

import (
	"easy-forum/internal/service"
	"easy-forum/pkg/utils"


	"github.com/gin-gonic/gin"
)

type GetReportListData struct {
	UserID int `form:"user_id"`
}

type GetReportListResponse struct {
	UserNmae string `json:"username"`
	PostID   int    `json:"post_id"`
	Content  string `json:"content"`
	Reason   string `json:"reason"`
}

func GetReportList(c *gin.Context) {
	var data GetReportListData
	if err := c.ShouldBindQuery(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	// 业务逻辑
	// 1.用户是否存在
	user, err := service.GetUserByID(data.UserID)
	if err != nil {
		utils.JsonErrorResponse(c, 200512, "用户不存在")
		return
	}
	// 2.是否是管理员
	if user.Type != 2 {
		utils.JsonErrorResponse(c, 200512, "用户不是管理员")
		return
	}
	// 3.获取未审批的举报列表
	list, err := service.GetAdminReportList()
	if err != nil {
		utils.JsonErrorResponse(c, 200513, "获取未审批的举报列表失败")
		return
	}
	var listResponse []GetReportListResponse
	for _, item := range list {
		// 4.获取帖子内容
		post, err := service.GetPostByID(item.PostID)
		if err != nil {
			utils.JsonErrorResponse(c, 200514, "获取帖子内容失败")
			return
		}
		user, err := service.GetUserByID(post.UserID)
		if err != nil {
			utils.JsonErrorResponse(c, 200514, "获取用户信息失败")
			return
		}
		// 5.返回帖子内容
		listResponse = append(listResponse, GetReportListResponse{
			UserNmae: user.Username,
			PostID:   post.ID,
			Content:  post.Content,
			Reason:   item.Reason,
		})
	}
	utils.JsonSuccess(c, gin.H{
		"report_list": listResponse,
	})
}

type ApproveReportData struct {
	PostID   int `json:"post_id" binding:"required"`
	Approval int `json:"approval" binding:"required"`
	UserID   int `json:"user_id" binding:"required"`
}

func ApproveReport(c *gin.Context) {
	var data ApproveReportData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	// 业务逻辑
	// 1.用户是否存在
	user, err := service.GetUserByID(data.UserID)
	if err != nil {
		utils.JsonErrorResponse(c, 200512, "用户不存在")
		return
	}
	// 2.是否是管理员
	if user.Type != 2 {
		utils.JsonErrorResponse(c, 200512, "用户不是管理员")
		return
	}
	// 3.是否是审批过的举报
	report, err := service.GetReportByPostID(data.PostID)
	if err != nil {
		utils.JsonErrorResponse(c, 200514, "举报不存在")
		return
	}
	if report.Status != 0 {
		utils.JsonErrorResponse(c, 200514, "举报已审批")
		return
	}
	// 4.审批举报
	err = service.ApproveReport(data.PostID, data.Approval)
	if err != nil {
		utils.JsonErrorResponse(c, 200515, "审批举报失败")
		return
	}
	utils.JsonSuccess(c, nil)
}
