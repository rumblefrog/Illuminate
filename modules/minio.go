package modules

import (
	"github.com/RumbleFrog/Illuminate/helpers"
	"github.com/minio/minio-go"
)

// MinioClient Pointers for Minio connection
var MinioClient minio.Client

// MinioConnect initiates the MinioClient pointer
func MinioConnect(creds helpers.MinioCredentials) {
	MinioClient, err := minio.New(creds.Endpoint, creds.AccessKey, creds.PrivateKey, creds.Secure)
}
