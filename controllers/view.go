package controllers

import (
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
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
	)

	if valid := bson.IsObjectIdHex(fName); !valid {
		w.WriteHeader(400)
		io.WriteString(w, "Invalid request")
		return
	}

	var (
		fHex = bson.ObjectIdHex(fName)
		err  error
	)

	c := modules.Database.C("illuminate")

	result := modules.Shrine{}

	if err = c.FindId(fHex).One(&result); err != nil {
		w.WriteHeader(500)
		io.WriteString(w, "Error retrieving object")
		log.Println(err)
		return
	}

	c.UpdateId(fHex, bson.M{
		"$inc": bson.M{
			"views": 1,
		},
	})

	var fObj *minio.Object

	fObj, err = modules.MinioClient.GetObject(
		helpers.Config.Minio.Bucket,
		object,
		minio.GetObjectOptions{},
	)

	var ObjInfo minio.ObjectInfo

	if ObjInfo, err = fObj.Stat(); err != nil {
		w.WriteHeader(500)
		io.WriteString(w, "Unable to fetch object")
		log.Println(err)
		return
	}

	if err != nil {
		w.WriteHeader(500)
		io.WriteString(w, "Unable to fetch object")
		log.Println(err)
		return
	}

	w.Header().Set("Views", strconv.FormatUint(result.Views, 10))
	w.Header().Set("Content-Type", ObjInfo.ContentType)

	io.Copy(w, fObj)
}
