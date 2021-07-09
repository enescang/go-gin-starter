package main

import (
	"fmt"
	"os"

	"github.com/enescang/go-gin-starter/db"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Database

func main() {
	godotenv.Load()
	client = db.Init()
	port := os.Getenv("PORT")
	fmt.Println("The PORT is: ", port)
}
