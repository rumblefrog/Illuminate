package modules

import (
	"log"

	"github.com/globalsign/mgo"
)

// Database for the MongoDB connection
var Database *mgo.Database

// MongoConnect returns a session pointer to the MongoDB
func MongoConnect(MongoURL string) {
	session, err := mgo.Dial(MongoURL)

	if err != nil {
		log.Fatal("MongoDB failed to connect: ", err)
	}

	Database = session.DB("illuminate")
}
