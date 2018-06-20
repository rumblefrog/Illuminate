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
		fHex  = bson.ObjectIdHex(fName)
		err   error
	)

	c := modules.Database.C("illuminate")

	result := modules.Shrine{}

	if err = c.FindId(fHex).One(&result); err != nil {
		io.WriteString(w, "Error retrieving object")
		fmt.Println(err)
		return
	}

	c.UpdateId(fHex, bson.M{
		"$inc": bson.M{
			"Views": 1,
		},
	})

	var fObj *minio.Object

	fObj, err = modules.MinioClient.GetObject(
		helpers.Config.Minio.Bucket,
		fName,
		minio.GetObjectOptions{},
	)

	io.Copy(w, fObj)
}
