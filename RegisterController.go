package main

import (
	"errors"
	"net/http"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	var requestBody RegisterRequestJsonObject

	if error := c.BindJSON(&requestBody); error != nil {
		c.JSON(http.StatusOK, &ResponseJsonObject{
			Code:    http.StatusBadRequest,
			Message: error.Error(),
		})
		return
	}

	var code int = http.StatusOK
	var message string = "success"

	if !isValidUserName(requestBody.UserName) {
		code = 200502
		message = "用户名必须为纯数字"
	} else if isContainedSameUserName(requestBody.UserName) {
		code = 200505
		message = "用户名已存在"
	} else if lengthOfPassword := utf8.RuneCountInString(requestBody.Password); lengthOfPassword < 8 || lengthOfPassword > 16 {
		code = 200503
		message = "密码长度必须在8-16位"
	} else if requestBody.UserType != 1 && requestBody.UserType != 2 {
		code = 200504
		message = "用户类型错误"
	} else if !writeUserDataRecord(requestBody) {
		code = http.StatusBadRequest
		message = "failed when trying to write to the database"
	}

	c.JSON(http.StatusOK, &ResponseJsonObject{
		Code:    code,
		Message: message,
	})
}

func isValidUserName(username string) bool {
	for _, char := range username {
		if char < '0' || char > '9' {
			return false
		}
	}

	return true
}

func isContainedSameUserName(username string) bool {
	datebase := getDataBase()

	var userData UserDataRecord
	result := datebase.First(&userData, "user_name = ?", username)

	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func writeUserDataRecord(requestBody RegisterRequestJsonObject) bool {
	datebase := getDataBase()

	//var count int64
	//datebase.Model(&UserDataRecord{}).Count(&count)

	result := datebase.Create(&UserDataRecord{
		Name:     requestBody.Name,
		UserName: requestBody.UserName,
		Password: requestBody.Password,
		UserType: requestBody.UserType,
		//UserId:   count + 1,
	})

	return result.Error == nil
}
