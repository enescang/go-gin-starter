package controllers

import (
	"context"
	"strings"
	"time"

	"github.com/enescang/go-gin-starter/db"
	"github.com/enescang/go-gin-starter/models"
	"github.com/enescang/go-gin-starter/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type authControllers struct {
	Signup func(*gin.Context)
	Login  func(*gin.Context)
}

var Auth authControllers = authControllers{
	Signup: signup,
	Login:  login,
}

//Signup >>>
type signupSchema struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"fullname" binding:"required"`
}

func signup(c *gin.Context) {
	var schema signupSchema
	DB, _ := db.Init()
	err := c.BindJSON(&schema)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"error": err.Error(),
		})
		return
	}
	var user models.UserSchema = models.UserSchema{
		Email:     strings.ToLower(schema.Email),
		Password:  schema.Password,
		FullName:  schema.FullName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hash)
	filter := bson.M{"email": user.Email}
	count, err := DB.Collection("users").CountDocuments(context.TODO(), filter)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"error": "Error while checking email",
		})
		return
	}
	if count > 0 {
		c.AbortWithStatusJSON(409, gin.H{
			"error": "Email already exists",
		})
		return
	}
	_, err = DB.Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User Signed up",
	})
}

//Signup <<<

//Login >>>

type loginSchema struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func login(c *gin.Context) {
	var schema loginSchema
	DB, _ := db.Init()
	err := c.BindJSON(&schema)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	var user models.UserSchema
	filter := bson.M{"email": strings.ToLower(schema.Email)}
	err = DB.Collection("users").FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"error": "User not found",
		})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(schema.Password))
	if err != nil {
		c.AbortWithStatusJSON(422, gin.H{
			"error": "Password incorrect",
		})
		return
	}
	var authClaim = utils.AuthClaims{
		UserID: user.ID,
	}
	accessToken, err := utils.SignToken(authClaim)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"error": "Please try again :(",
		})
		return
	}

	c.JSON(200, gin.H{
		"access_token": accessToken,
		"account":      user,
	})
}

//Login <<<
