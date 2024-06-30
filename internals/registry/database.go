package registry

import (
	"fmt"

	config "github.com/masterghost2002/videotube/configs"
	"github.com/masterghost2002/videotube/internals/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var storage *gorm.DB

func StorageInit() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", config.ENVS.DBAddress, config.ENVS.DBUser, config.ENVS.DBPassword, config.ENVS.DBName, config.ENVS.DBPort)
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database %w", err)
	}

	fmt.Printf("Successfully connected to database")
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}
	storage = db
	return nil
}
func GetDB() *gorm.DB {
	return storage
}
