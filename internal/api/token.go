package api

import (
	"log"
	"os"
)

var jwtSecret []byte

func init() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET not set")
	}

	jwtSecret = []byte(secret)
}
