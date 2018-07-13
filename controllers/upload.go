package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"

	"github.com/RumbleFrog/Illuminate/helpers"
	"github.com/RumbleFrog/Illuminate/modules"
	"github.com/globalsign/mgo/bson"
	"github.com/minio/minio-go"
)

// UploadController handles the /upload route
func UploadController(w http.ResponseWriter, r *http.Request) {
	file, headers, err := r.FormFile("payload")

	if err != nil {
		w.WriteHeader(400)
		io.WriteString(w, "Invalid payload")
		return
	}

	var (
		ID   = bson.NewObjectId()
		fExt = filepath.Ext(headers.Filename)
		oHex = fmt.Sprintf("%s%s", ID.Hex(), fExt)
	)

	_, err = modules.MinioClient.PutObject(
		helpers.Config.Minio.Bucket,
		oHex,
		file,
		headers.Size,
		minio.PutObjectOptions{ContentType: headers.Header.Get("Content-Type")},
	)

	if err != nil {
		w.WriteHeader(500)
		io.WriteString(w, "Unable to upload object")
		log.Println(err)
		return
	}

	c := modules.Database.C("illuminate")

	err = c.Insert(&modules.Shrine{
		ID:    ID,
		Views: 0,
	})

	if err != nil {
		w.WriteHeader(500)
		io.WriteString(w, "Unable to insert struct")
		log.Println(err)
		return
	}

	io.WriteString(w, fmt.Sprintf("%s/%s%s", helpers.Config.Root, ID.Hex(), fExt))
}
