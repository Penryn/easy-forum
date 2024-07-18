package user

import (
	"easy-forum/internal/service"
	"easy-forum/pkg/utils"

	"github.com/gin-gonic/gin"
)

type RegDate struct {
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	UserType int    `json:"user_type" binding:"required"`
}

func Reg(c *gin.Context) {
	var data RegDate
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	// 参数校验
	//用户名只有数字
	if !utils.IsNumber(data.Username) {
		utils.JsonErrorResponse(c, 200502, "用户名必须为纯数字")
		return
	}
	//密码长度
	if len(data.Password) < 8 || len(data.Password) > 16 {
		utils.JsonErrorResponse(c, 200503, "密码长度必须在8-16位")
		return
	}
	// 类型校验
	if data.UserType != 1 && data.UserType != 2 {
		utils.JsonErrorResponse(c, 200504, "用户类型错误")
		return
	}

	// 业务逻辑
	// 1.用户名是否存在
	_, err = service.GetUserByUsername(data.Username)
	if err == nil {
		utils.JsonErrorResponse(c, 200505, "用户名已存在")
		return
	}

	// 2.创建用户
	err = service.CreateUser(data.Username, data.Name, data.Password, data.UserType)
	if err != nil {
		utils.JsonErrorResponse(c, 200505, "注册失败")
		return
	}

	utils.JsonSuccess(c, nil)

}
