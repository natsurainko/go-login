package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var requestBody LoginRequestJsonObject

	if error := c.BindJSON(&requestBody); error != nil {
		c.JSON(http.StatusOK, &ResponseJsonObject{
			Code:    http.StatusBadRequest,
			Message: error.Error(),
		})
		return
	}

	var data any
	var code int = http.StatusOK
	var message string = "success"

	database := getDataBase()
	var userData UserDataRecord

	if result := database.Find(&userData, "user_name = ?", requestBody.UserName); result.RowsAffected == 0 {
		code = 200506
		message = "用户不存在"
	} else if userData.Password != requestBody.Password {
		code = 200507
		message = "密码错误"
	} else {
		data = &LoginDataResponseJsonObject{
			UserType: userData.UserType,
			UserId:   userData.UserId,
		}
	}

	c.JSON(http.StatusOK, &ResponseJsonObject{
		Code:    code,
		Data:    data,
		Message: message,
	})
}
