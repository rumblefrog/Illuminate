package controllers

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/RumbleFrog/Illuminate/helpers"
	"github.com/RumbleFrog/Illuminate/modules"
	"github.com/globalsign/mgo/bson"
	"github.com/minio/minio-go"
)

// ViewController handles all the base/non-routed requests
func ViewController(w http.ResponseWriter, r *http.Request) {
	object := r.URL.Path[1:]

	if object == "" {
		io.WriteString(w, "Illuminate the world")
		return
	}

	var (
		fExt  = filepath.Ext(object)
		fName = strings.TrimSuffix(object, fExt)
		err   error
	)

	c := modules.Database.C("illuminate")

	result := modules.Shrine{}

	if err = c.FindId(bson.ObjectIdHex(fName)).One(&result); err != nil {
		io.WriteString(w, "Error retrieving object")
		fmt.Println(err)
		return
	}

	var fObj *minio.Object

	fObj, err = modules.MinioClient.GetObject(
		helpers.Config.Minio.Bucket,
		fName,
		minio.GetObjectOptions{},
	)

	fmt.Println(result)
}
