package services

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mlinciko/TaskTrackerGin/config"
	xerrors "github.com/mlinciko/TaskTrackerGin/internal/errors"
	requests "github.com/mlinciko/TaskTrackerGin/internal/handlers/dtos/request"
	i_repositories "github.com/mlinciko/TaskTrackerGin/internal/repositories/interfaces"
	i_services "github.com/mlinciko/TaskTrackerGin/internal/services/interfaces"
	appstructs "github.com/mlinciko/TaskTrackerGin/internal/structs"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	repo        i_repositories.UserRepository
	userService i_services.UserService
}

func NewAuthService(repo i_repositories.UserRepository, userService i_services.UserService) i_services.AuthService {
	return &authService{repo, userService}
}

func (s *authService) Login(request requests.LoginRequestDto) (string, error) {
	user, err := s.userService.GetUserByEmail(request.Email)
	if err != nil {
		return "", xerrors.ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return "", xerrors.ErrInvalidCredetionals
	}

	claims := appstructs.JWTClaims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			Subject:   strconv.FormatUint(uint64(user.ID), 10),
		},
	}

	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, tokenErr := tokenClaim.SignedString([]byte(config.AppConfig.JWTSecretKey))
	if tokenErr != nil {
		return "", tokenErr
	}

	return token, nil
}
