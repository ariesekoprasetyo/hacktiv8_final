package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"hacktiv8_final/controller"
	"hacktiv8_final/posts"
	"hacktiv8_final/repository"
	"hacktiv8_final/router"
	"log"
	"net/http"
	"os"
	"time"
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

	//Init Repo
	commentRepo := repository.CommentRepo{Db: DB}

	//Init Domain
	commentService := posts.CommentService{CommentRepo: &commentRepo}

	//Init Controller
	commentController := controller.PostsController{PostsControl: &commentService}

	routes := router.NewRouter(&commentRepo, &commentController)

	server := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	serverErr := server.ListenAndServe()
	panic(serverErr)
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
