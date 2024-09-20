package controllers

import (
	"errors"
	"login-project/data"
	"login-project/services"
	"login-project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginController struct {
	UserService services.UserService
}

func (controller LoginController) Login(c *gin.Context) {
	var requestBody data.LoginRequestJsonObject

	if error := c.BindJSON(&requestBody); error != nil {
		utils.ResponseFrom(c, http.StatusBadRequest, "missing parameters")
		return
	}

	user, error := controller.UserService.FindUserByUserName(requestBody.UserName)

	if errors.Is(error, gorm.ErrRecordNotFound) {
		utils.ResponseFrom(c, 200506, "用户不存在")
	} else if error != nil {
		utils.ResponseFromError(c, error)
	} else if user.Password != requestBody.Password {
		utils.ResponseFrom(c, 200507, "密码错误")
	} else {
		utils.ResponseOkFromData(c, data.LoginDataResponseJsonObject{
			UserType: user.UserType,
			UserId:   user.UserId,
		})
	}
}
