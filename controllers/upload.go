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

// UploadController handles the /route route
func UploadController(w http.ResponseWriter, r *http.Request) {
	file, headers, _ := r.FormFile("payload")

	var (
		ID    = bson.NewObjectId()
		fExt  = filepath.Ext(headers.Filename)
		cType = headers.Header.Get("Content-Type")
		oHex  = fmt.Sprintf("%s%s", ID.Hex(), fExt)
	)

	_, err := modules.MinioClient.PutObject(
		helpers.Config.Minio.Bucket,
		oHex,
		file,
		headers.Size,
		minio.PutObjectOptions{ContentType: cType},
	)

	if err != nil {
		log.Fatalln(err)
	}

	c := modules.Database.C("illuminate")

	err = c.Insert(&modules.Shrine{
		ID:          ID,
		Ext:         fExt,
		ContentType: cType,
		Views:       0,
	})

	if err != nil {
		log.Println(err)
	}

	io.WriteString(w, fmt.Sprintf("%s/%s%s", helpers.Config.Root, ID.Hex(), fExt))
}
