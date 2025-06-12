package utils

import responses "gin/task_tracker/handlers/dtos/response"

func MakeResp(Data any, Code int, ErrorMessage string) *responses.Response {
	return &responses.Response{Data, Code, ErrorMessage}
}
