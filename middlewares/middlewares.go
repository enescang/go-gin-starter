package middlewares

import (
	"context"
	"fmt"
	"strings"

	"github.com/enescang/go-gin-starter/db"
	"github.com/enescang/go-gin-starter/models"
	"github.com/enescang/go-gin-starter/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func RequiresAuth(c *gin.Context) {
	var client, _ = db.Init()
	authorization := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(authorization, "Bearer ")
	fmt.Println(len(splitToken))
	fmt.Println(splitToken)
	if len(splitToken) < 2 {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "Authorization token is missing.",
		})
		return
	}
	claims, err := utils.VerifyToken(splitToken[1])
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"error": err,
		})
		return
	}
	var result models.UserSchema
	filter := bson.M{"_id": claims.UserID}
	err = client.Collection("users").FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"error": "User not found",
		})
		return
	}
}
