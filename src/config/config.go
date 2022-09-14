package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadConf() {
	// err := nil //godotenv.Load(".env")
	// if err != nil {
	// 	panic("Load .env fail ...")
	// }
	// console.Log("Load .env successfully ...")
}

type Conf map[string]interface{}

type Config struct {
	Port  string
	DbUri string
}

func GetConfig(key string) {
	conf := Config{
		Port:  os.Getenv("PORT"),
		DbUri: os.Getenv("DB_URI"),
	}
	var result Conf

	data, _ := json.Marshal(conf)
	json.Unmarshal(data, &result)
	fmt.Println(result["Port"])
}
