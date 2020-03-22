package main

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestDotEnv(t *testing.T) {
	godotenv.Load()

	if err := godotenv.Load(".env"); err != nil {
		t.Error("Error loading .env file")
	}

	if os.Getenv("AHREFS_TOKEN") == "" {
		t.Error("Empty token")
	}
}
