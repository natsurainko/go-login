package main

import (
	"net/http"
	"strconv"

	"github.com/ahmetb/go-linq/v3"
	"github.com/gin-gonic/gin"
)

func FetchUnauditedReports(c *gin.Context) {
	var UserId int64 = -1

	if user_id, isUserIdExist := c.GetQuery("user_id"); !isUserIdExist {
		c.JSON(http.StatusOK, &ResponseJsonObject{
			Code:    http.StatusBadRequest,
			Message: "missing parameters",
		})
	} else {
		UserId, _ = strconv.ParseInt(user_id, 10, 64)
	}

	var data any
	var message string = "success"
	var code int = http.StatusOK

	database := getDataBase()
	var userData UserDataRecord

	if result := database.Find(&userData, "user_id = ?", UserId); result.RowsAffected == 0 {
		code = 200506
		message = "用户不存在"
	} else if userData.UserType != 2 {
		code = http.StatusForbidden
		message = "用户无权限"
	} else {
		var reportsDatas []ReportDataRecord
		database.Find(&reportsDatas, "status = ?", 0)

		var ReportList []interface{} = linq.From(reportsDatas).SelectT(func(report ReportDataRecord) AdminViewReportItem {
			var post PostDataRecord
			var user UserDataRecord

			database.Find(&post, "post_id = ?", report.PostId)
			database.Find(&user, "user_id = ?", report.UserId)

			return AdminViewReportItem{
				PostId:   report.PostId,
				Reason:   report.Reason,
				Content:  post.Content,
				UserName: user.UserName,
				ReportId: report.ReportId,
			}
		}).Results()

		var DefaultReportList [0]AdminViewReportItem

		if len(ReportList) == 0 {
			data = &AdminViewReportResponseJsonObject{
				ReportList: DefaultReportList,
			}
		} else {
			data = &AdminViewReportResponseJsonObject{
				ReportList: ReportList,
			}
		}
	}

	c.JSON(http.StatusOK, &ResponseJsonObject{
		Code:    code,
		Data:    data,
		Message: message,
	})
}

func AuditedReport(c *gin.Context) {
	var requestBody AuditedReportRequestJsonObject

	if error := c.ShouldBindJSON(&requestBody); error != nil {
		c.JSON(http.StatusOK, &ResponseJsonObject{
			Code:    http.StatusBadRequest,
			Message: error.Error(),
		})
		return
	}

	var message string = "success"
	var code int = http.StatusOK

	database := getDataBase()
	var userData UserDataRecord

	if result := database.Find(&userData, "user_id = ?", requestBody.UserId); result.RowsAffected == 0 {
		code = 200506
		message = "用户不存在"
	} else if userData.UserType != 2 {
		code = http.StatusForbidden
		message = "用户无权限"
	} else {
		var reportData ReportDataRecord

		if result = database.Find(&reportData, "report_id = ?", requestBody.ReportId); result.RowsAffected == 0 {
			code = http.StatusBadRequest
			message = "举报不存在"
		} else if result = database.Model(ReportDataRecord{}).Where("report_id = ?", requestBody.ReportId).Update("Status", requestBody.Approval); result.Error != nil {
			code = http.StatusBadRequest
			message = "failed to audited report"
		}

		if requestBody.Approval == 1 {
			if result = database.Model(PostDataRecord{}).Where("post_id = ?", reportData.PostId).Update("IsDeleted", true); result.Error != nil {
				code = http.StatusBadRequest
				message = "failed to delete reported post"
			}
		}
	}

	c.JSON(http.StatusOK, &ResponseJsonObject{
		Code:    code,
		Message: message,
	})
}
