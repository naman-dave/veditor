package pkg

import (
	"log"

	"github.com/naman-dave/veditor/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() error {
	var err error

	DB, err = gorm.Open(sqlite.Open("database.sqlite"), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func Migrate() error {
	// Migrate the schema
	err := DB.AutoMigrate(&entity.Video{})
	if err != nil {
		return err
	}

	log.Println("Database connection established and schema migrated.")

	return nil
}
