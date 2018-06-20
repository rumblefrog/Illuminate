package modules

import (
	"github.com/globalsign/mgo/bson"
)

// Shrine struct
type Shrine struct {
	ID    bson.ObjectId `json:"id" bson:"_id"`
	Views uint32
}
