package main

import (
	"fmt"
	"log"
	"os"

	"github.com/enescang/go-gin-starter/controllers"
	"github.com/enescang/go-gin-starter/db"
	"github.com/enescang/go-gin-starter/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	client     *mongo.Database
	mongoError error
)

func main() {
	godotenv.Load()
	client, mongoError = db.Init()
	if mongoError != nil {
		log.Fatal("Mongodb Connection Error: ", mongoError)
	}
	port := os.Getenv("PORT")

	router := gin.New()

	router.GET("/test", middlewares.RequiresAuth, func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello"})
	})

	router.POST("/auth/signup", controllers.Auth.Signup)

	router.POST("/auth/login", controllers.Auth.Login)

	router.Run("127.0.0.1:" + port)
	fmt.Println("The PORT is: ", port, router)
}
