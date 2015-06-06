package main

import (
	"net/http" 
	"html/template" 
	"log"
	"fmt"
//	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	//	"strings"
	//	"os" 
	
	// youtube
	"flag"
	"github.com/Go-Docker-Hackathon/team-iDareX/vendor/download/youtube"
	"github.com/Go-Docker-Hackathon/team-iDareX/vendor/db/mongo"
)

var NWorkers = flag.Int("n", 5, "The number of workers to start")

// 任务的结构
type Task struct { 
	Fetchurl string
	Downloadurl string
	Status int // 0 队列中 1 下载中 2 下载完成 3 上传中 4 上传完成[可供用户下载]
}

func main() {
	// Parse the command-line flags.
	flag.Parse()

	// Start the dispatcher.
	fmt.Println("Starting the dispatcher")
	youtube.StartDispatcher(*NWorkers)
	
	http.HandleFunc("/", index)
	http.HandleFunc("/addurl", addurl)
	http.HandleFunc("/test", testedit)
	http.HandleFunc("/clear", clear)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func clear(w http.ResponseWriter, r *http.Request) {
	C := mongo.Connect()	
	C.Remove(bson.M{"fetchurl":"https://www.youtube.com/watch?v=SbY33tZl_pY"})
}

func testedit(w http.ResponseWriter, r *http.Request) {
	C := mongo.Connect()
	err := C.Update(bson.M{"fetchurl": "https://www.youtube.com/watch?v=SbY33tZl_pY"}, bson.M{"$set": bson.M{"downloadurl": "http://www.baidu.com"}})
	checkError(err)
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {		
		C := mongo.Connect()

		var tasks []Task
		C.Find(nil).All(&tasks) //查询所有
//	    for _, r := range tasks {
//			fmt.Println(" r.fetchurl = "+r.Fetchurl+" r.status = ")
//		}
		t, err := template.ParseFiles("views/index.gtpl")
		checkError(err)
		
		err = t.Execute(w, tasks)
		checkError(err)
	}
}

// 添加Youtube链接到队列
func addurl(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/index.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()	

		fetchurl := r.Form["fetchurl"][0]

		fmt.Println("url : = " + fetchurl)

		C := mongo.Connect()
		
		task := Task{} 
		err := C.Find(bson.M{"fetchurl": fetchurl}).One(&task)

		if err != nil {
			checkError(err)	
			youtube.Collector(fetchurl)
		} else {
			fmt.Println("任务已经在下载队列中")
		}		 
		
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
}