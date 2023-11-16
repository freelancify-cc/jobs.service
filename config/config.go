package config

import (
	"encoding/json"
	"log"
	"os"
)

// Config : struct
type Config struct {
	Environment         string `json:"env"`
    Host                string `json:"host"`
	Port                string `json:"port"`
	DbString            string `json:"db"`
}

var config *Config

func init() {
	file, err := os.Open("config/config.json")
	if err != nil {
		log.Fatal("[x] error : ", err.Error())
	}
	decoder := json.NewDecoder(file)
	conf := Config{}
	err = decoder.Decode(&conf)
	if err != nil {
		log.Fatal("[x] error", err.Error())
	}
	config = &conf
}

func GetConfig() *Config {
	return config
}
