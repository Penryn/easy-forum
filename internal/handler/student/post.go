package student

import (
	"easy-forum/internal/model"
	"easy-forum/internal/service"
	"easy-forum/pkg/utils"

	"github.com/gin-gonic/gin"
)

type CreatePostData struct {
	UserID  int    `json:"user_id"`
	Content string `json:"content"`
}

func CreatePost(c *gin.Context) {
	var data CreatePostData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	// 业务逻辑
	// 1.用户是否存在
	_, err := service.GetUserByID(data.UserID)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "用户不存在")
		return
	}
	// 2.创建帖子
	err = service.CreatePost(data.UserID, data.Content)
	if err != nil {
		utils.JsonErrorResponse(c, 200508, "创建失败")
		return
	}
	utils.JsonSuccess(c, nil)
}

type GetPostListResponse struct {
	PostList []model.Post `json:"post_list"`
}

func GetPostList(c *gin.Context) {
	// 业务逻辑
	// 1.获取帖子列表
	var postList []model.Post
	postList, err := service.GetPostList()
	if err != nil {
		utils.JsonErrorResponse(c, 200509, "获取失败")
		return
	}
	utils.JsonSuccess(c, gin.H{
		"post_list": postList,
	})
}

type UpdatePostData struct {
	UserID  int    `json:"user_id"`
	PostID  int    `json:"post_id"`
	Content string `json:"content"`
}

func UpdatePost(c *gin.Context) {
	var data UpdatePostData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	// 业务逻辑
	// 1.用户是否存在
	_, err := service.GetUserByID(data.UserID)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "用户不存在")
		return
	}
	// 2.帖子是否存在
	post, err := service.GetPostByID(data.PostID)
	if err != nil {
		utils.JsonErrorResponse(c, 200509, "帖子不存在")
		return
	}
	// 3.帖子是否属于用户
	if post.UserID != data.UserID {
		utils.JsonErrorResponse(c, 200510, "无权限")
		return
	}
	// 4.更新帖子
	err = service.UpdatePost(data.UserID, data.PostID, data.Content)
	if err != nil {
		utils.JsonErrorResponse(c, 200510, "更新失败")
		return
	}
	utils.JsonSuccess(c, nil)

}

type DeletePostData struct {
	UserID int `form:"user_id"`
	PostID int `form:"post_id"`
}

func DeletePost(c *gin.Context) {
	var data DeletePostData
	if err := c.ShouldBindQuery(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	// 业务逻辑
	// 1.用户是否存在
	user, err := service.GetUserByID(data.UserID)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "用户不存在")
		return
	}
	// 2.帖子是否存在
	post, err := service.GetPostByID(data.PostID)
	if err != nil {
		utils.JsonErrorResponse(c, 200509, "帖子不存在")
		return
	}
	// 3.帖子是否属于用户
	if user.ID != post.UserID {
		utils.JsonErrorResponse(c, 200510, "无权限")
		return
	}
	// 4.删除帖子和更新举报记录
	err = service.UpdataReportStatus(data.PostID, 1)
	if err != nil {
		utils.JsonErrorResponse(c, 200511, "删除失败")
		return
	}
	utils.JsonSuccess(c, nil)
}

type ReportPostData struct {
	UserID int    `json:"user_id"`
	PostID int    `json:"post_id"`
	Reason string `json:"reason"`
}

func ReportPost(c *gin.Context) {
	var data ReportPostData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	// 业务逻辑
	// 1.用户是否存在
	user, err := service.GetUserByID(data.UserID)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "用户不存在")
		return
	}
	// 2.帖子是否存在
	post, err := service.GetPostByID(data.PostID)
	if err != nil {
		utils.JsonErrorResponse(c, 200509, "帖子不存在")
		return
	}
	// 3.是否重复举报
	_, err = service.GetReportByPostIDAndUserID(user.ID, post.ID)
	if err == nil {
		utils.JsonErrorResponse(c, 200512, "重复举报")
		return
	}
	// 4.举报帖子
	err = service.ReportPost(data.UserID, data.PostID, data.Reason)
	if err != nil {
		utils.JsonErrorResponse(c, 200512, "举报失败")
		return
	}
	utils.JsonSuccess(c, nil)
}

type GetReportListData struct {
	UserID int `form:"user_id"`
}

type GetReportListResponse struct {
	PostID  int    `json:"post_id"`
	Content string `json:"content"`
	Reason  string `json:"reason"`
	Status  int    `json:"status"`
}

func GetReportList(c *gin.Context) {
	var data GetReportListData
	if err := c.ShouldBindQuery(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	_, err := service.GetUserByID(data.UserID)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "用户不存在")
		return
	}
	// 业务逻辑
	// 1.获取举报列表
	var reportList []model.Report
	reportList, err = service.GetReportList(data.UserID)
	if err != nil {
		utils.JsonErrorResponse(c, 200513, "获取失败")
		return
	}
	var reportListResponse []GetReportListResponse
	for _, report := range reportList {
		// 2.获取帖子内容
		post, err := service.GetAllPostByID(report.PostID)
		if err != nil {
			utils.JsonErrorResponse(c, 200514, "获取失败")
			return
		}
		// 3.返回帖子内容
		reportListResponse = append(reportListResponse, GetReportListResponse{
			PostID:  post.ID,
			Content: post.Content,
			Reason:  report.Reason,
			Status:  report.Status,
		})
	}

	utils.JsonSuccess(c, gin.H{
		"report_list": reportListResponse,
	})
}
