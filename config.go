package main

import (
	"encoding/json"
	"os"
)

type configuration struct {
	MongoURL string `json:"title"`
	S3       map[string]string
}

var config configuration

func loadConfig() Configuration {
	file, _ := os.Open("config.json")

	decoder := json.NewDecoder(file)

	config = configuration{}

	decoder.Decode(&config)

	defer file.Close()

	return config
}
