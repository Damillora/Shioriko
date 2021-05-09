package config

import "os"

type Config struct {
	PostgresDatabase    string
	AuthSecret          string
	DisableRegistration string
	DataDirectory       string
	BaseURL             string
}

var CurrentConfig Config

func InitializeConfig() {
	CurrentConfig = Config{
		PostgresDatabase:    os.Getenv("POSTGRES_DATABASE"),
		AuthSecret:          os.Getenv("AUTH_SECRET"),
		DisableRegistration: os.Getenv("DISABLE_REGISTRATION"),
		DataDirectory:       os.Getenv("DATA_DIR"),
		BaseURL:             os.Getenv("BASE_URL"),
	}
}
