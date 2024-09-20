package utils

import (
	"login-project/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseFrom(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, &data.ResponseJsonObject{
		Code:    code,
		Message: message,
	})
}

func ResponseFromError(c *gin.Context, _error error) {
	c.JSON(http.StatusOK, &data.ResponseJsonObject{
		Code:    http.StatusBadRequest,
		Message: _error.Error(),
	})
}

func ResponseOk(c *gin.Context) {
	c.JSON(http.StatusOK, &data.ResponseJsonObject{
		Code:    http.StatusOK,
		Message: "success",
	})
}

func ResponseOkFromData(c *gin.Context, _data any) {
	c.JSON(http.StatusOK, &data.ResponseJsonObject{
		Code:    http.StatusOK,
		Data:    _data,
		Message: "success",
	})
}
