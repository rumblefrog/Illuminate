package controllers

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/RumbleFrog/Illuminate/helpers"
	"github.com/RumbleFrog/Illuminate/modules"
	"github.com/minio/minio-go"
	"github.com/satori/go.uuid"
)

// UploadController handles the /route route
func UploadController(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Test")

	file, headers, _ := r.FormFile("payload")

	uid, err := uuid.NewV4()
	if err != nil {
		log.Fatalln(err)
	}

	ObjectName := fmt.Sprintf("%s%s", uid, filepath.Ext(headers.Filename))

	var n int64

	n, err = modules.MinioClient.PutObject(
		helpers.Config.Minio.Bucket,
		ObjectName,
		file,
		headers.Size,
		minio.PutObjectOptions{ContentType: headers.Header.Get("Content-Type")},
	)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(n)

	// fmt.Println(file.)
}
