package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"login-project/data"
	"login-project/services"
	"login-project/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminPostController struct {
	UserService   services.UserService
	ReportService services.ReportService
}

func (controller AdminPostController) FetchUnauditedReports(c *gin.Context) {
	var UserId int64 = -1

	if user_id, isUserIdExist := c.GetQuery("user_id"); !isUserIdExist {
		utils.ResponseFrom(c, http.StatusBadRequest, "missing parameters")
		return
	} else {
		UserId, _ = strconv.ParseInt(user_id, 10, 64)
	}

	user, error := controller.UserService.FindUserByUserId(UserId)

	if errors.Is(error, gorm.ErrRecordNotFound) {
		utils.ResponseFrom(c, http.StatusBadRequest, "用户不存在")
	} else if error != nil {
		utils.ResponseFromError(c, error)
	} else if user.UserType != 2 {
		utils.ResponseFrom(c, http.StatusForbidden, "用户无权限")
	} else {
		data, error := controller.ReportService.AdminViewReports()

		if error != nil {
			utils.ResponseFromError(c, error)
		} else {
			utils.ResponseOkFromData(c, data)
		}
	}
}

func (controller AdminPostController) AuditedReport(c *gin.Context) {
	var requestBody data.AuditedReportRequestJsonObject

	if error := c.BindJSON(&requestBody); error != nil {
		utils.ResponseFrom(c, http.StatusBadRequest, "missing parameters")
		return
	}

	user, error := controller.UserService.FindUserByUserId(requestBody.UserId)

	if errors.Is(error, gorm.ErrRecordNotFound) {
		utils.ResponseFrom(c, http.StatusBadRequest, "用户不存在")
	} else if error != nil {
		utils.ResponseFromError(c, error)
	} else if user.UserType != 2 {
		utils.ResponseFrom(c, http.StatusForbidden, "用户无权限")
	} else {
		report, error := controller.ReportService.FindReportByReportId(requestBody.ReportId)

		if errors.Is(error, gorm.ErrRecordNotFound) {
			utils.ResponseFrom(c, http.StatusBadRequest, "举报不存在")
		} else if error != nil {
			utils.ResponseFromError(c, error)
		} else if error := controller.ReportService.AuditedReport(report.ReportId, report.PostId, requestBody.Approval); error != nil {
			utils.ResponseFromError(c, error)
		} else {
			utils.ResponseOk(c)
		}
	}
}
