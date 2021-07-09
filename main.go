package main

import (
	"fmt"
	"os"

	"github.com/enescang/go-gin-starter/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Database

func main() {
	godotenv.Load()
	client, _ = db.Init()
	port := os.Getenv("PORT")

	router := gin.Default()

	fmt.Println("The PORT is: ", port, router)
}
