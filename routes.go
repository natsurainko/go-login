package main

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.POST("/api/user/reg", RegisterController.Register)
	r.POST("/api/user/login", LoginController.Login)

	r.POST("/api/student/post", StudentPostController.ReleasePost)
	r.GET("/api/student/post", StudentPostController.FetchAllPosts)
	r.PUT("/api/student/post", StudentPostController.ModifyPost)
	r.DELETE("/api/student/post", StudentPostController.DeletePost)

	r.POST("/api/student/report-post", StudentPostController.ReportPost)
	r.GET("/api/student/report-post", StudentPostController.ViewReport)

	r.GET("/api/admin/report", AdminPostController.FetchUnauditedReports)
	r.POST("/api/admin/report", AdminPostController.AuditedReport)
}
