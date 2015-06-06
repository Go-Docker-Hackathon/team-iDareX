package mongo

import (
	"labix.org/v2/mgo"
)

var Session *mgo.Session
var Collection *mgo.Collection

func Connect() *mgo.Collection{
	Session, _ = mgo.Dial("localhost")

	Session.SetMode(mgo.Monotonic, true)		
	Collection = Session.DB("docker").C("task")
	
	return Collection
}