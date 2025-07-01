package database

import (
	"log"

	"github.com/mlinciko/TaskTrackerGin/internal/models"
)

func Migrate() {
	err := DB.AutoMigrate(models.Registry()...)
	if err != nil {
		log.Fatalf("❌ Auto migration failed: %v", err)
	}

	log.Println("✅ Migration done!")
}
