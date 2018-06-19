package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RumbleFrog/Illuminate/controllers"
	"github.com/RumbleFrog/Illuminate/helpers"
	"github.com/RumbleFrog/Illuminate/middlewares"
	"github.com/RumbleFrog/Illuminate/modules"
)

func main() {
	helpers.LoadConfig()
	modules.MinioConnect(helpers.Config.Minio)
	modules.MongoConnect(helpers.Config.MongoURL)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print(r.URL.Path[1:])
		fmt.Fprint(w, "Illuminate the world")
	})

	http.HandleFunc("/upload", middlewares.IsPost(middlewares.IsFishy(controllers.UploadController)))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", helpers.Config.Port), nil))
}
