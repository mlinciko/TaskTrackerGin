package tests

import (
	test_helpers "gin/task_tracker/tests/utils"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	test_helpers.InitializeTestDB(m)
}
