package modules

import (
	"gopkg.in/mgo.v2"
)

var session mgo.v2.Session

func mongoConnect() {
	session, err := mgo.dial(config.MongoURL)
}
