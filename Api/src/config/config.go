package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DATABASE_FULL_ADD = ""
	APP_HOST_PORT     = 0
)

// Carrega configurações
func Load() {

	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	APP_HOST_PORT, err = strconv.Atoi(os.Getenv("APP_HOST_PORT"))
	if err != nil {
		APP_HOST_PORT = 9000
	}

	DATABASE_FULL_ADD = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASS"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
	)
}
