package mongo

import (
	"labix.org/v2/mgo"
	"os"
//	"fmt"
)

var Session *mgo.Session
var Collection *mgo.Collection

func Connect() *mgo.Collection{


	MONGODB_USERNAME := os.Getenv("MONGODB_USERNAME")
	MONGODB_PASSWORD := os.Getenv("MONGODB_PASSWORD")
	MONGODB_PORT_27017_TCP_ADDR := os.Getenv("MONGODB_PORT_27017_TCP_ADDR")	
	MONGODB_PORT_27017_TCP_PORT := os.Getenv("MONGODB_PORT_27017_TCP_PORT")	
	MONGODB_INSTANCE_NAME := os.Getenv("MONGODB_INSTANCE_NAME")

//	fmt.Println("username "+ MONGODB_USERNAME)	
	if MONGODB_USERNAME != "" {
		Session, _ = mgo.Dial("mongodb://" + MONGODB_USERNAME + ":" + MONGODB_PASSWORD + "@" + MONGODB_PORT_27017_TCP_ADDR + ":"  + MONGODB_PORT_27017_TCP_PORT + "/" + MONGODB_INSTANCE_NAME)		
	} else {
//		fmt.Println("username "+ MONGODB_USERNAME)
		
		if MONGODB_PORT_27017_TCP_ADDR != ""  {
			Session, _ = mgo.Dial("mongodb://" + MONGODB_PORT_27017_TCP_ADDR + ":"  + MONGODB_PORT_27017_TCP_PORT + "/" + "docker")			
		}else{
			Session, _ = mgo.Dial("localhost")
		}
	}

//	mongodb://myuser:mypass@localhost:40001/mydb	

	Session.SetMode(mgo.Monotonic, true)		
	
	if MONGODB_INSTANCE_NAME != "" {
		Collection = Session.DB(MONGODB_INSTANCE_NAME).C("task")
	} else {
		Collection = Session.DB("docker").C("task")
	}
	
	return Collection
}