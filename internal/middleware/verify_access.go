package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mlinciko/TaskTrackerGin/internal/utils"
	jwt_utils "github.com/mlinciko/TaskTrackerGin/internal/utils/jwt"
)

func VerifyAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		ok, err := jwt_utils.VerifyToken(c)
		log.Println("VerifyToken: ", ok)

		if !ok {
			c.JSON(http.StatusUnauthorized, utils.MakeResp(nil, http.StatusUnauthorized, err.Error()))
			c.Abort()
			return
		}

		c.Next()
	}
}
