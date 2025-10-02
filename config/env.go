package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Prop struct {
	ServerPort string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBUrl 	   string
	JwtSecret string
}

var AppProp Prop

func InitializeAppProps() {
		// Load .env only in local dev
	if os.Getenv("RAILWAY_ENVIRONMENT") == "" {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, using system environment")
		}
	}

	AppProp = Prop{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "local_db"),
		DBUrl: 		getEnv("DB_URL", "postgres://root:pass@localhost:5432/finance_tracker?sslmode=disable"),
		JwtSecret:  getEnv("JWT_SECRET", ""),
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