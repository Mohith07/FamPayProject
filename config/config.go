package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Configuration struct {
	RestServerPort  string        `json:"rest_server_port"`
	DBHost          string        `json:"db_host"`
	DBUsername      string        `json:"db_username"`
	DBPassword      string        `json:"db_password"`
	DBPort          string        `json:"db_port"`
	DBType          string        `json:"db_type"`
	DBSSLMode       string        `json:"dbssl_mode"`
	DBName          string        `json:"db_name"`
	DefaultAPIKey   string		  `json:"default_youtube_apikey"`
}

var config *Configuration

func Init() {
	file, fileErr := os.Open("config/conf.json")
	if fileErr != nil {
		fmt.Printf(fileErr.Error())
		log.Fatal(fileErr.Error())
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Error on parsing configuration file. Error " + err.Error())
		}
	}(file)
	decoder := json.NewDecoder(file)
	config = &Configuration{}
	err := decoder.Decode(config)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("port is - %s\n", config.RestServerPort)
}

func GetConfig() *Configuration {
	return config
}
