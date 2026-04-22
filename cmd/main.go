package main

import (
	"fmt"
	"log"
	"newsapi/internal"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := internal.NewDB(connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	repo := internal.NewRepository(db)
	svc := internal.NewService(repo)
	res := internal.NewHandler(svc)

	r := gin.Default()

	r.POST("/auth/login", res.Login)
	r.POST("/auth/register", res.Register)
	authorized := r.Group("/")
	authorized.Use(internal.AuthMiddleware())
	authorized.GET("/news", res.GetArticle)
	authorized.POST("/news/fetch", res.FetchNews)
	r.Run(os.Getenv("SERVER_PORT"))

}
