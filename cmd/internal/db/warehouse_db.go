package db

import (
	"github.com/FatiaGlacier/navigation-service/cmd/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var PostgresDB *gorm.DB

func InitPostgresDB() {
	dsn := config.GetEnv("POSTGRES_DSN", "")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	PostgresDB = db
}
