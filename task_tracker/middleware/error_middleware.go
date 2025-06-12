package middleware

import (
	"gin/task_tracker/utils"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			httpCode := err.Meta.(int)

			c.JSON(httpCode, utils.MakeResp(nil, httpCode, err.Error()))
			return
		}
	}
}
