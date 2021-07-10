package utils

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestSignToken(t *testing.T) {
	myAuthClaims := AuthClaims{
		UserID: primitive.NewObjectID(),
	}
	token, err := SignToken(myAuthClaims)
	fmt.Println(token)
	assert.Nil(t, err)
	assert.Equal(t, len(strings.Split(token, ".")), 3)
}

func TestVerifyToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiI2MGU5ZDc4Y2U2ODE1ZDJiNGVmZGEyY2YifQ.WnwHaOaa0oWkD8aclHIS95k1Nag3V7hUhZpVCStdqxY"
	myClaims, err := VerifyToken(token)
	assert.Nil(t, err)
	assert.Equal(t, false, myClaims.UserID.IsZero())

	//disrupt the token
	token = "XXXXXX.YYYYY.ZZZZZ"
	myClaims, err = VerifyToken(token)
	assert.Error(t, err)
	assert.Equal(t, true, myClaims.UserID.IsZero())
}

func TestTokenExpiration(t *testing.T) {
	myAuthClaims := AuthClaims{
		primitive.NewObjectID(),
		jwt.StandardClaims{
			//Assume that we wait 180 second
			ExpiresAt: time.Now().Unix() - (3 * 60),
		},
	}
	token, err := SignToken(myAuthClaims)
	assert.Nil(t, err)
	claims, err := VerifyToken(token, true)
	assert.Equal(t, true, claims.UserID.IsZero())
	assert.Regexp(t, regexp.MustCompile("expired"), err)
}
