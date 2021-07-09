package db

import (
	"os"
	"regexp"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestDbConnection(t *testing.T) {
	godotenv.Load("../.env")
	_, err = Init()
	assert.Nil(t, err)

	os.Setenv("MONGO_URI", "broken-mongo-uri")
	_, err = Init()
	assert.Regexp(t, regexp.MustCompile("parsing uri"), err)
}
