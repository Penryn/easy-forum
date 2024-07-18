package router

import (
	s "easy-forum/internal/handler/student"
	u "easy-forum/internal/handler/user"
	a "easy-forum/internal/handler/admin"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine){
	const pre ="api"

	api := r.Group(pre)
	{
		user := api.Group("/user")
		user.POST("/reg",u.Reg)
		user.POST("/login",u.Login)
		
		student := api.Group("/student")
		student.POST("/post",s.CreatePost)
		student.GET("/post",s.GetPostList)
		student.PUT("/post",s.UpdatePost)
		student.DELETE("/post",s.DeletePost)
		student.POST("/report-post",s.ReportPost)
		student.GET("/report-post",s.GetReportList)

		admin := api.Group("/admin")
		admin.GET("/report",a.GetReportList)
		admin.POST("/report",a.ApproveReport)

	}

}