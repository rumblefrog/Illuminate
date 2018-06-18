package modules

import (
	"gopkg.in/mgo.v2"
)

// MongoConnect returns a session pointer to the MongoDB
func MongoConnect(MongoURL string) (session *mgo.Session) {
	session, err := mgo.Dial(MongoURL)

	return
}
