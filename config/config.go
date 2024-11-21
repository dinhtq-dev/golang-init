package config

import (
	"golang/constants"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
	AppLang string

	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

// Config returns a new AppConfig instance with values loaded from environment variables
func Config(envFile string) (*AppConfig, error) {
	err := godotenv.Load(envFile)
	
	if err != nil {
		log.Printf("Could not load %s, using environment variables instead.", envFile)
	}

	return &AppConfig{
		AppName: getEnv("APP_NAME", constants.Config.AppEnv),
		AppEnv:  getEnv("APP_ENV", constants.Config.AppEnv),
		AppPort: getEnv("APP_PORT", constants.Config.AppPort),
		AppLang: getEnv("APP_LANG", constants.Config.AppLang),

		DBHost: getEnv("DB_HOST", constants.Config.DBHost),
		DBPort: getEnv("DB_PORT", constants.Config.DBPort),
		DBUser: getEnv("DB_USER", constants.Config.DBUser),
		DBPass: getEnv("DB_PASS", constants.Config.DBPass),
		DBName: getEnv("DB_NAME", constants.Config.DBName),
		
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
