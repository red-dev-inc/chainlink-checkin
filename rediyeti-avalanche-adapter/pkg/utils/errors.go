package utils

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"message"`
}

func SendError(c *gin.Context, code int, message string) {
	e := ErrorResponse{
		Code:  code,
		Error: message,
	}
	c.JSON(code, e)
}
