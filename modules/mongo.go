package modules

import (
	"log"

	"gopkg.in/mgo.v2"
)

// Session for the MongoDB connection
var Session *mgo.Session

// MongoConnect returns a session pointer to the MongoDB
func MongoConnect(MongoURL string) {
	var err error

	Session, err = mgo.Dial(MongoURL)

	if err != nil {
		log.Fatal(err)
	}
}
