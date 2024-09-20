package controllers

import (
	"net/http"
	"unicode/utf8"

	"login-project/data"
	"login-project/services"
	"login-project/utils"

	"github.com/gin-gonic/gin"
)

type RegisterController struct {
	UserService services.UserService
}

func (controller RegisterController) Register(c *gin.Context) {
	var requestBody data.RegisterRequestJsonObject

	if error := c.BindJSON(&requestBody); error != nil {
		utils.ResponseFrom(c, http.StatusBadRequest, "missing parameters")
		return
	}

	if !controller.isValidUserName(requestBody.UserName) {
		utils.ResponseFrom(c, 200502, "用户名必须为纯数字")
	} else if lengthOfPassword := utf8.RuneCountInString(requestBody.Password); lengthOfPassword < 8 || lengthOfPassword > 16 {
		utils.ResponseFrom(c, 200503, "密码长度必须在8-16位")
	} else if requestBody.UserType != 1 && requestBody.UserType != 2 {
		utils.ResponseFrom(c, 200504, "用户类型错误")
	} else if containsUser, error := controller.UserService.ContainsUserName(requestBody.UserName); error != nil || containsUser {
		if error != nil {
			utils.ResponseFromError(c, error)
		} else {
			utils.ResponseFrom(c, 200505, "用户名已存在")
		}
	} else if error := controller.UserService.RegisterUser(requestBody); error != nil {
		utils.ResponseFrom(c, http.StatusBadRequest, "failed when trying to write to the database")
	} else {
		utils.ResponseOk(c)
	}
}

func (controller RegisterController) isValidUserName(username string) bool {
	for _, char := range username {
		if char < '0' || char > '9' {
			return false
		}
	}

	return true
}
