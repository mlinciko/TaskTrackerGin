package jwt_utils

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	xerrors "github.com/mlinciko/TaskTrackerGin/internal/errors"
	appstructs "github.com/mlinciko/TaskTrackerGin/internal/structs"
)

const TOKEN_HEADER_KEY = "Authorization"

func VerifyToken(c *gin.Context) (bool, error) {
	header := c.GetHeader(TOKEN_HEADER_KEY)

	tokenString := strings.TrimSpace(strings.Replace(header, "Bearer", "", 1))
	token, err := ParseToken(tokenString)

	if err != nil {
		return false, err
	}

	claims, ok := token.Claims.(*appstructs.JWTClaims)
	if ok && token.Valid {
		if float64(time.Now().Unix()) > float64(claims.ExpiresAt) {
			return false, xerrors.ErrExpiredToken
		}
	} else {
		return false, xerrors.ErrInvalidToken
	}

	return true, nil
}
