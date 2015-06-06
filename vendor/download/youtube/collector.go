package youtube

import (
	"github.com/Go-Docker-Hackathon/team-iDareX/vendor/db/mongo"
		"labix.org/v2/mgo/bson"
)

var WorkQueue = make(chan WorkRequest, 100)

func Collector(url string) {

	work := WorkRequest{ Url: url}
	
	// insert to mongo , status is in queue.
	
	WorkQueue <- work
	
	// change task status
	C := mongo.Connect()
	C.Update(bson.M{"fetchurl": url}, bson.M{"$set": bson.M{"status": 1}}) // downloading
}