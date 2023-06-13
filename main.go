package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"hacktiv8_final/controller"
	"hacktiv8_final/posts"
	"hacktiv8_final/repository"
	"hacktiv8_final/router"
	"hacktiv8_final/user"
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

	//Setup DB And Migration
	setupDB()
	validate := validator.New()

	//Init Repo
	commentRepo := repository.CommentRepo{Db: DB}
	photoRepo := repository.PhotoRepo{Db: DB}
	userRepo := repository.UsersRepo{Db: DB}

	//Init Domain
	commentService := posts.CommentService{
		CommentRepo: &commentRepo,
		Validate:    validate,
	}
	photoService := posts.PhotoService{
		PhotoRepo: photoRepo,
		Validate:  validate,
	}

	userService := user.Authz{
		AuthRepo:  &userRepo,
		Validate:  validate,
		JwtSecret: os.Getenv("JWT_SECRET"),
	}

	//Init Controller
	commentController := controller.CommentController{Service: &commentService}
	photoController := controller.PhotoController{Service: &photoService}
	authController := controller.AuthController{Service: &userService}

	routes := router.NewRouter(&commentController, &photoController, &authController)

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
		log.Fatal("[INIT] failet connected to PostgreSQL")
		return
	}
	err = DB.AutoMigrate(&repository.User{}, &repository.Photo{}, &repository.SocialMedia{}, &repository.Comment{})
	if err != nil {
		panic(err)
	}
	log.Println("Migrasi Berhasil")
	log.Println("[INIT] connected to PostgreSQL")
}
