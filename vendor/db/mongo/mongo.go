package mongo

import (
	"labix.org/v2/mgo"
	"os"
)

var Session *mgo.Session
var Collection *mgo.Collection

func Connect() *mgo.Collection{

MONGODB_PORT
MONGODB_PORT_27017_TCP_ADDR

MONGODB_PORT_27017_TCP_PROTO
MONGODB_PORT_27017_TCP_PORT




	MONGODB := os.Getenv("MongoDB")
	MONGODB_USERNAME := os.Getenv("MONGODB_USERNAME")
	MONGODB_PASSWORD := os.Getenv("MONGODB_PASSWORD")
	MONGODB_PORT_27017_TCP := os.Getenv("MONGODB_PORT_27017_TCP")	
	MONGODB_INSTANCE_NAME := os.Getenv("MONGODB_INSTANCE_NAME")

//	mongodb://myuser:mypass@localhost:40001/mydb	
	Session, _ = mgo.Dial(MONGODB+"://" + MONGODB_USERNAME + ":" + MONGODB_PASSWORD + "@(" + MONGODB_PORT_27017_TCP+")/" + MONGODB_INSTANCE_NAME)
	

	Session.SetMode(mgo.Monotonic, true)		
	Collection = Session.DB(MONGODB_INSTANCE_NAME).C("task")
	
	return Collection
}