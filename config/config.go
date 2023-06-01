package config

import (
	"github.com/joho/godotenv"
	"log"
)

type Configuration struct {
	Server   SetupServer
	Database SetupDatabase
}

type SetupServer struct {
	Port string
}

type SetupDatabase struct {
	DBUser string
	DBPass string
	DBName string
	DBHost string
	DBPort string
}

var (
	Config *Configuration
)

func SetupConfiguration() {
	var envs map[string]string
	envs, err := godotenv.Read(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Config = new(Configuration)
	Config.Server.Port = envs["PORT"]
	Config.Database.DBHost = envs["DB_HOST"]
	Config.Database.DBPort = envs["DB_PORT"]
	Config.Database.DBName = envs["DB_NAME"]
	Config.Database.DBUser = envs["DB_USER"]
	Config.Database.DBPass = envs["DB_PASS"]
}
