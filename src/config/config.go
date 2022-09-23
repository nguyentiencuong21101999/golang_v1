package config

import (
	console "main/src/helpers/consoles"
	"os"

	"github.com/joho/godotenv"
)

func LoadConf() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Load .env fail ...")
	}
	console.Info("Load .env successfully ...")
}

type Config struct {
	Port  string `json:"port"`
	DbUri string `json:"dbUri"`
}

func GetConfig() *Config {
	conf := Config{
		Port:  os.Getenv("PORT"),
		DbUri: os.Getenv("DB_URI"),
	}
	return &conf
}
