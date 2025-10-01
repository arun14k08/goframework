package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Prop struct {
	ServerPort string
	DBUrl string
}

var AppProp Prop

func InitializeAppProps() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading environment variables")
	}

	AppProp = Prop{
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetAppProp() Prop {
	return AppProp
}