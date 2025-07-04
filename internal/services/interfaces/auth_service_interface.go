package i_services

import requests "github.com/mlinciko/TaskTrackerGin/internal/handlers/dtos/request"

type AuthService interface {
	Login(request requests.LoginRequestDto) (string, error)
}
