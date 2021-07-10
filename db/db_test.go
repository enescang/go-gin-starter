package db

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestDbConnection(t *testing.T) {
	godotenv.Load("../.env")
	_, err = Init()
	assert.Nil(t, err)

	//Even the MONGO_URI change Init() must not throwing any error
	//Because we use singleton pattern to connection DB
	os.Setenv("MONGO_URI", "broken-mongo-uri")
	_, err = Init()
	assert.Nil(t, err)
}
