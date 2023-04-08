package lib

import (
	"api/models"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var (
	DB  *gorm.DB
	err error
)

func InitDatabase() {
	err = godotenv.Load()
	if err != nil {
		panic(err)
	}

	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_DBNAME"), port, os.Getenv("PG_SSLMODE"), os.Getenv("PG_TIMEZONE"),
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel:                  logger.Info,
			ParameterizedQueries:      false,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
		})

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	// Auto migration for now...
	DB.AutoMigrate(&models.Message{})

	fmt.Printf("Connected to PostgreSQL on %s:%s\n", host, port)
	if err != nil {
		panic(err)
	}
}
