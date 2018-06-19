package helpers

import (
	"encoding/json"
	"log"
	"os"
)

// Config helds the pointers to the Configuration struct
var Config Configuration

// Configuration for the Illuminate process
type Configuration struct {
	Port     uint64
	MongoURL string
	Minio    *MinioCredentials
}

// MinioCredentials is a sub-struct of Configuration and provides
// info regarding the Minio instance
type MinioCredentials struct {
	Endpoint   string
	Secure     bool
	AccessKey  string
	PrivateKey string
	Bucket     string
}

// LoadConfig reads and parse the JSON file into readable struct
func LoadConfig() {
	file, err := os.Open("config.json")

	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)

	Config = Configuration{}

	decoder.Decode(&Config)

	defer file.Close()
}
