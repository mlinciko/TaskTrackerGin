package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	xerrors "github.com/mlinciko/TaskTrackerGin/internal/errors"
	"github.com/mlinciko/TaskTrackerGin/internal/handlers"
	requests "github.com/mlinciko/TaskTrackerGin/internal/handlers/dtos/request"
	responses "github.com/mlinciko/TaskTrackerGin/internal/handlers/dtos/response"
	i_handlers "github.com/mlinciko/TaskTrackerGin/internal/handlers/interfaces"
	"github.com/mlinciko/TaskTrackerGin/internal/routes"
	"github.com/mlinciko/TaskTrackerGin/internal/tests/factories"
	"github.com/mlinciko/TaskTrackerGin/internal/tests/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserHandlerTestSuite struct {
	suite.Suite
	router      *gin.Engine
	handler     i_handlers.UserHandler
	mockService *mocks.MockUserService
}

func (s *UserHandlerTestSuite) SetupTest() {
	mockService := new(mocks.MockUserService)

	r := gin.Default()

	s.router = r
	s.mockService = mockService
	s.handler = handlers.NewUserHandler(s.mockService)

	routes.SetupMiddleware(s.router)
	routes.SetUsersRoutes(s.router, s.handler)
}

// Test cases
func (s *UserHandlerTestSuite) TestUserHandler_CreateUser_BadRequest() {
	tests := []struct {
		name    string
		payload string
	}{
		{
			name:    "empty JSON",
			payload: `{}`,
		},
		{
			name:    "email email",
			payload: `{"firstName": "Alice"}`,
		},
		{
			name:    "invalid email",
			payload: `{"firstName": "Alice", "email": "not-an-email"}`,
		},
		{
			name:    "invalid JSON",
			payload: `{"firstName": "Alice", "email": `,
		},
	}

	for _, tt := range tests {
		req, _ := http.NewRequest("POST", "/users", bytes.NewBufferString(tt.payload))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		s.router.ServeHTTP(resp, req)

		s.Equal(http.StatusBadRequest, resp.Code, tt.name)
	}
}

func (s *UserHandlerTestSuite) TestUserHandler_CreateUser_InternalServerError() {
	s.mockService.On("CreateUser", mock.AnythingOfType("*models.User")).Return(assert.AnError)

	reqBody := requests.CreateUserRequestDto{
		FirstName: "Bob",
		Email:     "bob@example.com",
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	s.router.ServeHTTP(resp, req)

	s.Equal(http.StatusInternalServerError, resp.Code)
}

func (s *UserHandlerTestSuite) TestUserHandler_CreateUser_Success() {
	s.mockService.On("CreateUser", mock.AnythingOfType("*models.User")).Return(nil)
	body := `{"firstName":"Bob","email":"bob@example.com"}`

	req, _ := http.NewRequest("POST", "/users", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	s.router.ServeHTTP(resp, req)

	var parsedResp responses.Response

	err := json.Unmarshal(resp.Body.Bytes(), &parsedResp)
	s.NoError(err)
	s.Equal(http.StatusOK, resp.Code)
	s.Equal(http.StatusOK, parsedResp.Code)
	s.Equal("", parsedResp.ErrorMessage)
	s.Equal("Bob", parsedResp.Data.(map[string]interface{})["firstName"])
	s.Equal("bob@example.com", parsedResp.Data.(map[string]interface{})["email"])
}

func (s *UserHandlerTestSuite) TestUserHandler_GetUserHandler_BadRequest() {
	testTable := []struct {
		param string
		name  string
	}{
		{
			param: "abc",
			name:  "non-numeric id",
		},
		{
			param: "10.98",
			name:  "double id",
		},
		{
			param: "90*",
			name:  "special symbols in id",
		},
	}

	for _, tt := range testTable {
		url := "/users/" + tt.param

		req, _ := http.NewRequest("GET", url, nil)
		resp := httptest.NewRecorder()

		s.router.ServeHTTP(resp, req)

		var parsedResp responses.Response
		err := json.Unmarshal(resp.Body.Bytes(), &parsedResp)

		s.NoError(err)
		s.Equal(http.StatusBadRequest, resp.Code)
		s.Equal(nil, parsedResp.Data)
		s.Equal(http.StatusBadRequest, parsedResp.Code)
		s.Equal(xerrors.ErrInvalidUserId.Error(), parsedResp.ErrorMessage)
	}
}

func (s *UserHandlerTestSuite) TestUserHandler_GetUserHandler_NotFound() {
	s.mockService.On("GetUserByID", uint(18)).Return(nil, assert.AnError)

	req, _ := http.NewRequest("GET", "/users/18", nil)
	resp := httptest.NewRecorder()

	s.router.ServeHTTP(resp, req)

	var parsedResp responses.Response
	err := json.Unmarshal(resp.Body.Bytes(), &parsedResp)

	s.NoError(err)
	s.Equal(http.StatusNotFound, resp.Code)
	s.Equal(nil, parsedResp.Data)
	s.Equal(http.StatusNotFound, parsedResp.Code)
	s.Equal(xerrors.ErrUserNotFound.Error(), parsedResp.ErrorMessage)
}

func (s *UserHandlerTestSuite) TestUserHandler_GetUserHandler_Success() {
	user := factories.CreateUser()
	s.mockService.On("GetUserByID", uint(1)).Return(user, nil)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	resp := httptest.NewRecorder()

	s.router.ServeHTTP(resp, req)

	var parsedResp responses.Response
	err := json.Unmarshal(resp.Body.Bytes(), &parsedResp)
	s.NoError(err)

	s.Equal(http.StatusOK, resp.Code)

	s.Equal(float64(user.ID), parsedResp.Data.(map[string]interface{})["id"])
	s.Equal(user.FirstName, parsedResp.Data.(map[string]interface{})["firstName"])
	s.Equal(user.Email, parsedResp.Data.(map[string]interface{})["email"])
	s.Equal(user.CreatedAt.Format(time.RFC3339Nano), parsedResp.Data.(map[string]interface{})["createdAt"])

	s.Equal(http.StatusOK, parsedResp.Code)
	s.Equal("", parsedResp.ErrorMessage)
}

func TestUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}
