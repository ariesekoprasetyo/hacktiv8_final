package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"hacktiv8_final/repository"
	"log"
	"os"
)

var DB *gorm.DB

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	setupDB()
	err = DB.AutoMigrate(&repository.User{}, &repository.Photo{}, &repository.SocialMedia{}, &repository.Comment{})
	if err != nil {
		panic(err)
	}
	log.Println("Migrasi Berhasil")
}

func setupDB() {
	var err error
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed To Connect")
		return
	}
	log.Println("[INIT] connected to PostgreSQL")
}
