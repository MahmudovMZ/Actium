package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database Db `json:"db"`
}

type Db struct {
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
	Address  string `json:"address"`
}

var conf Config

func ReadConfig(path string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, &conf)
}

func GetConf() *Config {
	return &conf
}

var JWTSecret []byte

func LoadEnv() {
	// Загружаем переменные из .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	JWTSecret = []byte(os.Getenv("JWT_SECRET"))
}
