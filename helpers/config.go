package helpers

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	MongoURL string `json:"title"`
	S3       map[string]string
}

var Config Configuration

func LoadConfig() Configuration {
	file, _ := os.Open("config.json")

	decoder := json.NewDecoder(file)

	Config = Configuration{}

	decoder.Decode(&Config)

	defer file.Close()

	return Config
}
