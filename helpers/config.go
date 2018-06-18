package helpers

import (
	"encoding/json"
	"os"
)

// Configuration for the Illuminate process
type Configuration struct {
	Port     uint8
	MongoURL string
	Minio    MinioCredentials
}

// MinioCredentials is a sub-struct of Configuration and provides
// info regarding the Minio instance
type MinioCredentials struct {
	Endpoint   string
	Port       uint8
	Secure     bool
	AccessKey  string
	PrivateKey string
	Bucket     string
}

// LoadConfig reads and parse the JSON file into readable struct
func LoadConfig() Configuration {
	file, _ := os.Open("config.json")

	decoder := json.NewDecoder(file)

	config := Configuration{}

	decoder.Decode(&config)

	defer file.Close()

	return config
}
