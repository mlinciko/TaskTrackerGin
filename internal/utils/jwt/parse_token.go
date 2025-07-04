package jwt_utils

import (
	"github.com/golang-jwt/jwt"
	"github.com/mlinciko/TaskTrackerGin/config"
	appstructs "github.com/mlinciko/TaskTrackerGin/internal/structs"
)

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &appstructs.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JWTSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
