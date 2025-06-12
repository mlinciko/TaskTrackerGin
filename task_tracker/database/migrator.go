package database

import (
	"gin/task_tracker/models"
	"log"
)

func Migrate() {
	err := DB.AutoMigrate(models.Registry()...)
	if err != nil {
		log.Fatalf("❌ Auto migration failed: %v", err)
	}

	log.Println("✅ Migration done!")
}
