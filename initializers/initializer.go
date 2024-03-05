package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	log.Print(".env loaded")
}

func InitDb() *gorm.DB {
	dsn := "user=" + os.Getenv("DB_USERNAME") + " password=" + os.Getenv("DB_PASSWORD") + " host=" + os.Getenv("DB_HOSTNAME") + " port=" + os.Getenv("DB_PORT") + " dbname=" + os.Getenv("DB_NAME") + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Print("Connection to database established.", db)
	return db
}
