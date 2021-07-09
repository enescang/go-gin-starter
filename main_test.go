package main

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestEnvVariable(t *testing.T) {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port != "3000" {
		t.Error("PORT variable not 3000")
	} else {
		t.Log("godotenv package is running correctly")
	}
}
