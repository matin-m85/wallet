package repository

import (
	"gift-service/config"
	"gift-service/internal/models"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(cfg *config.Config) {
	var err error

	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(postgres.Open(cfg.GetDSN()), &gorm.Config{})
		if err == nil {
			break
		}
		log.Println("Waiting for database, retrying in 2s...")
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatal("[error] failed to initialize database:", err)
	}
	if err := DB.AutoMigrate(&models.GiftCardGroup{}, &models.GiftCard{}); err != nil {
		log.Fatal("[error] failed to migrate tables:", err)
	}
	log.Println("Connected to database successfully")
}
