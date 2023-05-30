package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	HOST string `json:"host" binding:"required"`
	DSN  string `json:"dsn"`
}

var configFile = map[string]string{
	"debug":   "config.dev.toml",
	"release": "config.toml",
}

var CONFIG = Config{
	HOST: "127.0.0.1:8080",
	DSN:  "",
}

func GetConfig(mode string) *Config {
	if mode == "" {
		mode = "debug"
	}
	viper.SetConfigFile(configFile[mode])
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
	host := viper.GetString("HOST")
	CONFIG.HOST = host
	var postgresPassword string
	if mode == "debug" {
		postgresPassword = viper.GetString("postgresPassword")
	} else {
		postgresPassword = os.Getenv("postgresPassword")
	}
	dbConnectString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		viper.GetString("postgresHost"),
		viper.GetString("postgresUser"),
		postgresPassword,
		viper.GetString("postgresDB"),
		viper.GetInt("postgresPort"),
	)
	CONFIG.DSN = dbConnectString
	return &CONFIG
}
