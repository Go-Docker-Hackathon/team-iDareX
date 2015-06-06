package youtube

import (
	"fmt"
	"github.com/Go-Docker-Hackathon/team-iDareX/vendor/upload/qiniu"
	"github.com/Go-Docker-Hackathon/team-iDareX/vendor/db/mongo"
	"labix.org/v2/mgo/bson"
)

type Worker struct{
	ID int
	Work chan WorkRequest
	WorkerQueue chan chan WorkRequest
	QuitChan chan bool
}

func NewWroker(id int, workerQueue chan chan WorkRequest) Worker {
	// Create, and return the worker.
	
	worker := Worker{
		ID: id,
		Work: make(chan WorkRequest),
		WorkerQueue: workerQueue,
		QuitChan: make(chan bool)}
	
	return worker
}

func (w Worker) Start() {
	go func() {
		for {
			// Add ourselves into the worker queue.
			w.WorkerQueue <- w.Work
			
			select {
			case work := <-w.Work:
				// Receive a work request.
				fmt.Printf("worker%d: Received work request\n", w.ID)
				fmt.Printf("worker %d: Url: %s\n", w.ID, work.Url)
				
				fileName, err := YoutubeDl(work.Url)
				if err != nil {
					fmt.Println("error with YoutubeDl:", err)
				}
				fmt.Println("filename:", fileName, "on worker")
				
				// change task status
				C := mongo.Connect()
				C.Update(bson.M{"fetchurl": work.Url}, bson.M{"$set": bson.M{"status": 2}}) // downloading
				fmt.Println("mongodb downloading")
				
				fmt.Println("upload to qiniu: ", fileName)
				C.Update(bson.M{"fetchurl": work.Url}, bson.M{"$set": bson.M{"status": 3}}) // start upload
				fmt.Println("mongodb start upload")

				key, err1 := qiniu.UploadQiniu(fileName)
				if err1 != nil {
					fmt.Println("upload file to qiniu error:", err1, "filename:", fileName)
				}else{
					C.Update(bson.M{"fetchurl": videoUrl}, bson.M{"$set": bson.M{"status": 4}}) // upload finish
					fmt.Println("mongodb upload finish")
					downloadurl := "http://7xjhxh.com1.z0.glb.clouddn.com" + key
					C.Update(bson.M{"fetchurl": work.Url}, bson.M{"$set": bson.M{"downloadurl": downloadurl}}) // set download key
					fmt.Println("mongodb set download key")
					fmt.Println("upload success, key:", key)
				}
				
			case <-w.QuitChan:
				// We have been asked to stop.
				fmt.Printf("worker %d stopping\n", w.ID)
			}
		}
	}()
}

// Stop tells the worker to stop listening for work requests.
// 
// Note that the worker will only stop *after* it has finished its work.
func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}