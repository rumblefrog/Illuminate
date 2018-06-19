package modules

import (
	"log"

	"github.com/RumbleFrog/Illuminate/helpers"
	"github.com/minio/minio-go"
)

// MinioClient Pointers for Minio connection
var MinioClient *minio.Client

// MinioConnect initiates the MinioClient pointer
func MinioConnect(creds helpers.MinioCredentials) {
	var err error

	MinioClient, err = minio.New(creds.Endpoint, creds.AccessKey, creds.PrivateKey, creds.Secure)

	if err != nil {
		log.Fatal(err)
	}
}
