package modules

import (
	"log"

	"github.com/RumbleFrog/Illuminate/helpers"
	"github.com/minio/minio-go"
)

// MinioClient Pointers for Minio connection
var MinioClient *minio.Client

// MinioConnect fufills the MinioClient pointer
func MinioConnect(creds *helpers.MinioCredentials) {
	var err error

	MinioClient, err = minio.New(creds.Endpoint, creds.AccessKey, creds.PrivateKey, creds.Secure)

	if err != nil {
		log.Fatalln(err)
	}

	var exists bool

	exists, err = MinioClient.BucketExists(creds.Bucket)

	if err != nil {
		log.Fatalln(err)
	}

	if !exists {
		err = MinioClient.MakeBucket(creds.Bucket, "us-east-1")

		if err != nil {
			log.Fatalln(err)
		}
	}
}
