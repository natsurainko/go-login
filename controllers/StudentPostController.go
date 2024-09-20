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

type StudentPostController struct {
	UserService   services.UserService
	PostService   services.PostService
	ReportService services.ReportService
}

func (controller StudentPostController) ReleasePost(c *gin.Context) {
	var requestBody data.ReleasePostRequestJsonObject

	if error := c.BindJSON(&requestBody); error != nil {
		utils.ResponseFrom(c, http.StatusBadRequest, "missing parameters")
		return
	}

	if containsUserId, _ := controller.UserService.ContainsUserId(requestBody.UserId); !containsUserId {
		utils.ResponseFrom(c, http.StatusBadRequest, "用户不存在")
	} else if error := controller.PostService.ReleasePost(requestBody); error != nil {
		utils.ResponseFromError(c, error)
	} else {
		utils.ResponseOk(c)
	}
}

func (controller StudentPostController) FetchAllPosts(c *gin.Context) {
	posts, error := controller.PostService.FetchAllPosts()

	if error != nil {
		utils.ResponseFromError(c, error)
	} else {
		utils.ResponseOkFromData(c, &data.FetchAllPostsResponseJsonObject{
			PostList: posts,
		})
	}
}

func (controller StudentPostController) ModifyPost(c *gin.Context) {
	var requestBody data.ModifyPostRequestJsonObject

	if error := c.BindJSON(&requestBody); error != nil {
		utils.ResponseFrom(c, http.StatusBadRequest, "missing parameters")
		return
	}

	postData, error := controller.PostService.FindPostByPostId(requestBody.PostId)

	if errors.Is(error, gorm.ErrRecordNotFound) {
		utils.ResponseFrom(c, http.StatusBadRequest, "帖子不存在")
	} else if error != nil {
		utils.ResponseFromError(c, error)
	} else if postData.IsDeleted {
		utils.ResponseFrom(c, http.StatusBadRequest, "帖子已删除")
	} else if postData.UserId != requestBody.UserId {
		utils.ResponseFrom(c, http.StatusBadRequest, "无权修改")
	} else if error := controller.PostService.UpdatePostContent(requestBody.Content, requestBody.PostId); error != nil {
		utils.ResponseFromError(c, error)
	} else {
		utils.ResponseOk(c)
	}
}

func (controller StudentPostController) DeletePost(c *gin.Context) {
	var UserId int64 = -1
	var PostId int64 = -1

	user_id, isUserIdExist := c.GetQuery("user_id")
	post_id, isPostIdExist := c.GetQuery("post_id")

	if !isUserIdExist || !isPostIdExist {
		utils.ResponseFrom(c, http.StatusBadRequest, "missing parameters")
		return
	}

	UserId, _ = strconv.ParseInt(user_id, 10, 64)
	PostId, _ = strconv.ParseInt(post_id, 10, 64)

	postData, error := controller.PostService.FindPostByPostId(PostId)

	if errors.Is(error, gorm.ErrRecordNotFound) {
		utils.ResponseFrom(c, http.StatusBadRequest, "帖子不存在")
	} else if error != nil {
		utils.ResponseFromError(c, error)
	} else if postData.IsDeleted {
		utils.ResponseFrom(c, http.StatusBadRequest, "帖子已删除")
	} else if postData.UserId != UserId {
		utils.ResponseFrom(c, http.StatusBadRequest, "无权删除")
	} else if error := controller.PostService.DeletePost(PostId); error != nil {
		utils.ResponseFromError(c, error)
	} else {
		utils.ResponseOk(c)
	}
}

func (controller StudentPostController) ReportPost(c *gin.Context) {
	var requestBody data.ReportPostRequestJsonObject

	if error := c.BindJSON(&requestBody); error != nil {
		utils.ResponseFrom(c, http.StatusBadRequest, "missing parameters")
		return
	}

	postData, error := controller.PostService.FindPostByPostId(requestBody.PostId)

	if errors.Is(error, gorm.ErrRecordNotFound) {
		utils.ResponseFrom(c, http.StatusBadRequest, "帖子不存在")
	} else if error != nil {
		utils.ResponseFromError(c, error)
	} else if containsUserId, _ := controller.UserService.ContainsUserId(requestBody.UserId); !containsUserId {
		utils.ResponseFrom(c, http.StatusBadRequest, "用户不存在")
	} else if postData.IsDeleted {
		utils.ResponseFrom(c, http.StatusBadRequest, "帖子已删除")
	} else if error := controller.ReportService.AddReport(requestBody); error != nil {
		utils.ResponseFromError(c, error)
	} else {
		utils.ResponseOk(c)
	}
}

func (controller StudentPostController) ViewReport(c *gin.Context) {
	var UserId int64 = -1

	if user_id, isUserIdExist := c.GetQuery("user_id"); !isUserIdExist {
		utils.ResponseFrom(c, http.StatusBadRequest, "missing parameters")
		return
	} else {
		UserId, _ = strconv.ParseInt(user_id, 10, 64)
	}

	if containsUserId, _ := controller.UserService.ContainsUserId(UserId); !containsUserId {
		utils.ResponseFrom(c, http.StatusBadRequest, "用户不存在")
	} else {
		data, error := controller.ReportService.ViewReportsFromUserId(UserId)

		if error != nil {
			utils.ResponseFromError(c, error)
		} else {
			utils.ResponseOkFromData(c, data)
		}
	}
}
