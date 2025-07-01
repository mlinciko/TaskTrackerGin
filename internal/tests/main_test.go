package tests

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	m.Run()
	//TODO fix DB for testing
	//test_helpers.InitializeTestDB(m)
}
