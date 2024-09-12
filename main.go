package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DatabaseCache *gorm.DB

func main() {
	r := gin.Default()

	// Routes
	r.POST("/api/user/reg", Register)
	r.POST("/api/user/login", Login)

	r.POST("/api/student/post", ReleasePost)
	r.GET("/api/student/post", FetchAllPosts)
	r.PUT("/api/student/post", ModifyPost)
	r.DELETE("/api/student/post", DeletePost)

	r.POST("/api/student/report-post", ReportPost)
	r.GET("/api/student/report-post", ViewReport)

	r.GET("/api/admin/report", FetchUnauditedReports)
	r.POST("/api/admin/report", AuditedReport)

	r.Run()
}

func getDataBase() *gorm.DB {
	if DatabaseCache != nil {
		return DatabaseCache
	}

	db, error := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if error != nil {
		panic("failed to connect database:" + error.Error())
	}

	DatabaseCache = db

	if !db.Migrator().HasTable(&UserDataRecord{}) {
		db.AutoMigrate(&UserDataRecord{})
	}
	if !db.Migrator().HasTable(&PostDataRecord{}) {
		db.AutoMigrate(&PostDataRecord{})
	}
	if !db.Migrator().HasTable(&ReportDataRecord{}) {
		db.AutoMigrate(&ReportDataRecord{})
	}

	return db
}

type ResponseJsonObject struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"msg"`
}
