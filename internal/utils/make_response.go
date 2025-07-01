package utils

import responses "github.com/mlinciko/TaskTrackerGin/internal/handlers/dtos/response"

func MakeResp(Data any, Code int, ErrorMessage string) *responses.Response {
	return &responses.Response{Data, Code, ErrorMessage}
}
