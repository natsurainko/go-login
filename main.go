package main

import (
	"github.com/gin-gonic/gin"

	"login-project/controllers"
	"login-project/dao"
	"login-project/services"
)

var UserDao dao.UserDao
var PostDao dao.PostDao
var ReportDao dao.ReportDao

var DataBaseService services.DataBaseService = services.DataBaseService{}
var UserService services.UserService
var PostService services.PostService
var ReportService services.ReportService

var RegisterController controllers.RegisterController
var LoginController controllers.LoginController
var StudentPostController controllers.StudentPostController
var AdminPostController controllers.AdminPostController

func main() {
	r := gin.Default()

	InitServices()
	InitControllers()
	InitRoutes(r)

	r.Run()
}

func InitServices() {
	DataBaseService.InitDataBase()

	UserDao = dao.UserDao{
		DataBase: DataBaseService.DataBase,
	}
	PostDao = dao.PostDao{
		DataBase: DataBaseService.DataBase,
	}
	ReportDao = dao.ReportDao{
		DataBase: DataBaseService.DataBase,
	}

	UserService = services.UserService{
		UserDao: UserDao,
	}
	PostService = services.PostService{
		PostDao: PostDao,
	}
	ReportService = services.ReportService{
		ReportDao:   ReportDao,
		PostService: PostService,
	}
}

func InitControllers() {
	RegisterController = controllers.RegisterController{
		UserService: UserService,
	}

	LoginController = controllers.LoginController{
		UserService: UserService,
	}

	StudentPostController = controllers.StudentPostController{
		UserService:   UserService,
		PostService:   PostService,
		ReportService: ReportService,
	}

	AdminPostController = controllers.AdminPostController{
		UserService:   UserService,
		ReportService: ReportService,
	}
}
