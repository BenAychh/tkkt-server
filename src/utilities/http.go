package utilities

import "github.com/gin-gonic/gin"

func ResponseError(c *gin.Context, statusCode int, errorMessageToUser string) {
	c.JSON(statusCode, gin.H{
		"isError":     true,
		"userMessage": errorMessageToUser,
	})
}

func RespondData(c *gin.Context, statusCode int, data interface{}, messageToUser ...string) {
	if len(messageToUser) > 0 {
		c.JSON(statusCode, gin.H{
			"isError":     false,
			"userMessage": messageToUser[0],
			"data":        data,
		})
	} else {
		c.JSON(statusCode, gin.H{
			"isError": false,
			"data":    data,
		})
	}
}
