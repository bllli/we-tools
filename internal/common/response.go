package common

import "github.com/gin-gonic/gin"

func WriteResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(200, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}

func OK(c *gin.Context, data interface{}) {
	WriteResponse(c, 0, "OK", data)
}

func Fail(c *gin.Context, message string) {
	WriteResponse(c, 1, message, nil)
}

func FailWithData(c *gin.Context, message string, data interface{}) {
	WriteResponse(c, 1, message, data)
}
