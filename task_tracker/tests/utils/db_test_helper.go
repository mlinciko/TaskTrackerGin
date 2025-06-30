package test_helpers

import (
	"database/sql"
	"fmt"
	"gin/task_tracker/config"
	"gin/task_tracker/database"
	"gin/task_tracker/models"
	"log"
	"os"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TestDB holds the test database connection
var TestDB *gorm.DB

func createTestDB() {
	config.LoadConfig()
	dsn := config.AppConfig.GetPostgresDSN()

	DB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to default database: %v", err)
	}
	defer DB.Close()

	// Test DB name
	testDBName := config.AppConfig.DBName + "_test"

	// Terminate all connections to the test database
	_, err = DB.Exec(fmt.Sprintf(`
    -- Terminate all connections to the test database
    SELECT pg_terminate_backend(pg_stat_activity.pid)
    FROM pg_stat_activity
    WHERE pg_stat_activity.datname = '%s'
    AND pid <> pg_backend_pid();
   `, testDBName))

	if err != nil {
		log.Fatalf("Failed to terminate connections to the test database %s: %v", testDBName, err)
	}

	// Drop the test database if it exists
	_, err = DB.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", testDBName))
	if err != nil {
		log.Fatalf("Failed to drop test database %s: %v", testDBName, err)
	}

	// Create the test database
	_, err = DB.Exec(fmt.Sprintf("CREATE DATABASE %s", testDBName))
	if err != nil {
		log.Fatalf("Failed to create test database %s: %v", testDBName, err)
	} else {
		log.Printf("Test Database %s created successfully", testDBName)
	}
}

func SetupTestDB() {
	createTestDB()
	dsn := config.AppConfig.GetTestPostgresDSN()

	var err error
	TestDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to test database: %v", err)
	}

	// Run migrations (you can add your models here)
	TestDB.AutoMigrate(models.Registry()...)
}

// ResetTestDatabase resets the database after each test
func ResetTestDatabase() {
	TestDB.Exec("TRUNCATE TABLE items RESTART IDENTITY CASCADE")
}

// TearDownTestDatabase closes the database connection
func TearDownTestDatabase() {
	sqlDB, err := TestDB.DB()
	if err != nil {
		log.Fatalf("Failed to close the database: %v", err)
	}
	sqlDB.Close()
}

// PatchDatabase replaces the global database.DB with TestDB for testing
func PatchDatabase() {
	database.DB = TestDB
}

// UnpatchDatabase restores the original database.DB (if needed)
func UnpatchDatabase() {
	// Reset database.DB back to the development DB after testing
	config.LoadConfig()
	dsn := config.AppConfig.GetPostgresDSN()
	database.Connect(dsn)
}

// InitializeTestDB sets up the test environment and patches the database connection
func InitializeTestDB(m *testing.M) {
	// Initialize the test database
	SetupTestDB()
	// Patch the global database.DB to use the TestDB
	PatchDatabase()
	// Run all tests
	code := m.Run()
	// Teardown the database connection
	TearDownTestDatabase()
	// Unpatch the database
	UnpatchDatabase()
	// Exit with the test run status code
	os.Exit(code)
}
