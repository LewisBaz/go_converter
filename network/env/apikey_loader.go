package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func Getapikey() (string, error) {
	err := godotenv.Load()
	if err != nil {
        fmt.Println("Error loading .env file:", err)
		return "", err
    }
	apikey := os.Getenv("ExchangeRate_API_KEY")
	return apikey, nil
}