package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ahmetb/go-linq/v3"
	"github.com/gin-gonic/gin"
)

func ReleasePost(c *gin.Context) {
	var requestBody ReleasePostRequestJsonObject

	if error := c.BindJSON(&requestBody); error != nil {
		c.JSON(http.StatusOK, &ResponseJsonObject{
			Code:    http.StatusBadRequest,
			Message: error.Error(),
		})
		return
	}

	var userData UserDataRecord
	database := getDataBase()

	//var count int64
	var code int = http.StatusOK
	var message string = "success"

	//database.Model(PostDataRecord{}).Count(&count)

	if database.Find(&userData, "user_id = ?", requestBody.UserId).RowsAffected == 0 {
		code = http.StatusBadRequest
		message = "用户不存在"
	} else if result := database.Create(&PostDataRecord{
		Content: requestBody.Content,
		UserId:  requestBody.UserId,
		Time:    time.Now().String(),
		//PostId:  count + 1,
	}); result.Error != nil {
		code = http.StatusBadRequest
		message = "failed when trying to write to the database"
	}

	c.JSON(http.StatusOK, &ResponseJsonObject{
		Code:    code,
		Message: message,
	})
}

func FetchAllPosts(c *gin.Context) {
	datebase := getDataBase()

	var postDataRecords []PostDataRecord
	datebase.Find(&postDataRecords, "is_deleted = ?", false)

	c.JSON(http.StatusOK, &ResponseJsonObject{
		Code: http.StatusOK,
		Data: &FetchAllPostsResponseJsonObject{
			PostList: postDataRecords,
		},
		Message: "success",
	})
}

func ModifyPost(c *gin.Context) {
	var requestBody ModifyPostRequestJsonObject

	if error := c.BindJSON(&requestBody); error != nil {
		c.JSON(http.StatusOK, &ResponseJsonObject{
			Code:    http.StatusBadRequest,
			Message: error.Error(),
		})
		return
	}

	var message string = "success"
	var code int = http.StatusOK

	database := getDataBase()
	var postData PostDataRecord
	result := database.Find(&postData, "post_id = ?", requestBody.PostId)

	if result.Error != nil || result.RowsAffected == 0 {
		message = "帖子不存在"
		code = http.StatusBadRequest
	} else if postData.IsDeleted {
		message = "帖子已删除"
		code = http.StatusBadRequest
	} else if postData.UserId != requestBody.UserId {
		message = "无权修改"
		code = http.StatusBadRequest
	} else if updateResult := database.Model(PostDataRecord{}).Where("post_id = ?", requestBody.PostId).Update("Content", requestBody.Content); updateResult.Error != nil {
		message = "failed to edit post"
		code = http.StatusBadRequest
	}

	c.JSON(http.StatusOK, &ResponseJsonObject{
		Code:    code,
		Message: message,
	})
}

func DeletePost(c *gin.Context) {
	var UserId int64 = -1
	var PostId int64 = -1

	user_id, isUserIdExist := c.GetQuery("user_id")
	post_id, isPostIdExist := c.GetQuery("post_id")

	if !isUserIdExist || !isPostIdExist {
		c.JSON(http.StatusOK, &ResponseJsonObject{
			Code:    http.StatusBadRequest,
			Message: "missing parameters",
		})
		return
	}

	UserId, _ = strconv.ParseInt(user_id, 10, 64)
	PostId, _ = strconv.ParseInt(post_id, 10, 64)

	var message string = "success"
	var code int = http.StatusOK

	var postData PostDataRecord

	database := getDataBase()
	result := database.Find(&postData, "post_id = ?", PostId)

	if result.RowsAffected == 0 || result.Error != nil {
		code = http.StatusBadRequest
		message = "帖子不存在"
	} else if postData.IsDeleted {
		message = "帖子已删除"
		code = http.StatusBadRequest
	} else if postData.UserId != UserId {
		code = http.StatusForbidden
		message = "无权删除"
	} else if deleteResult := database.Model(PostDataRecord{}).Where("post_id = ?", PostId).Update("IsDeleted", true); deleteResult.Error != nil {
		code = http.StatusBadRequest
		message = "failed to delete post"
	}

	c.JSON(http.StatusOK, &ResponseJsonObject{
		Code:    code,
		Message: message,
	})
}

func ReportPost(c *gin.Context) {
	var requestBody ReportPostRequestJsonObject

	if error := c.BindJSON(&requestBody); error != nil {
		c.JSON(http.StatusOK, &ResponseJsonObject{
			Code:    http.StatusBadRequest,
			Message: error.Error(),
		})
		return
	}

	var message string = "success"
	var code int = http.StatusOK

	var postData PostDataRecord

	datebase := getDataBase()
	result := datebase.Find(&postData, "post_id = ?", requestBody.PostId)

	if result.Error != nil || result.RowsAffected == 0 {
		message = "帖子不存在"
		code = http.StatusBadRequest
	} else if postData.IsDeleted {
		message = "帖子已删除"
		code = http.StatusBadRequest
	} else if reportResult := datebase.Create(&ReportDataRecord{
		Reason: requestBody.Reason,
		UserId: requestBody.UserId,
		PostId: requestBody.PostId,
		Status: 0,
	}); reportResult.Error != nil {
		message = "failed when trying to write to the database"
		code = http.StatusBadRequest
	}

	c.JSON(http.StatusOK, &ResponseJsonObject{
		Code:    code,
		Message: message,
	})
}

func ViewReport(c *gin.Context) {
	var UserId int64 = -1

	if user_id, isUserIdExist := c.GetQuery("user_id"); !isUserIdExist {
		c.JSON(http.StatusOK, &ResponseJsonObject{
			Code:    http.StatusBadRequest,
			Message: "missing parameters",
		})
		return
	} else {
		UserId, _ = strconv.ParseInt(user_id, 10, 64)
	}

	var data any
	var message string = "success"
	var code int = http.StatusOK

	var userData UserDataRecord
	var reportsOfUser []ReportDataRecord

	database := getDataBase()

	if result := database.Find(&userData, "user_id = ?", UserId); result.RowsAffected == 0 {
		code = http.StatusBadRequest
		message = "用户不存在"
	} else if reportResult := database.Find(&reportsOfUser, "user_id = ?", UserId); reportResult.Error != nil {
		code = http.StatusBadRequest
		message = "failed when trying to read the database"
	}

	var ReportList []interface{} = linq.From(reportsOfUser).SelectT(func(report ReportDataRecord) ViewReportItem {
		var post PostDataRecord
		database.Find(&post, "post_id = ?", report.PostId)

		return ViewReportItem{
			PostId:  report.PostId,
			Reason:  report.Reason,
			Content: post.Content,
			Status:  report.Status,
		}
	}).Results()

	var DefaultReportList [0]ViewReportItem

	if code != http.StatusBadRequest {
		if len(ReportList) == 0 {
			data = &ViewReportResponseJsonObject{
				ReportList: DefaultReportList,
			}
		} else {
			data = &ViewReportResponseJsonObject{
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
