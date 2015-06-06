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
	Status int // 0 队列中 1 下载中 2 下载完成 3 上传中 4 上传完成[可供用户下载], 上传完成需要将Downloadurl修改成上传完的地址
}

func main() {
	// Parse the command-line flags.
	flag.Parse()

	// Start the dispatcher.
	fmt.Println("Starting the dispatcher")
	youtube.StartDispatcher(*NWorkers)
	
	http.HandleFunc("/", index)
	http.HandleFunc("/addurl", addurl)
	http.HandleFunc("/search", search)
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

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {		
		C := mongo.Connect()

		var tasks []Task
		err := C.Find(nil).Sort("-_id").Limit(20).All(&tasks) //查询所有
		checkError(err)
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
		formatId := r.Form["formatId"][0]
		if fetchurl == "" {
			http.Redirect(w, r, "/", 302)
		}
		
		fmt.Println("url : = " + fetchurl)

		C := mongo.Connect()
		
		task := Task{} 
		err := C.Find(bson.M{"fetchurl": fetchurl}).One(&task)
		
		if err != nil {
			fmt.Println("err", err.Error())
			
			err = C.Insert(&Task{fetchurl, "", 0})
			youtube.Collector(fetchurl, formatId)
		} else {
			fmt.Println("任务已经在下载队列中")
		}		 

		http.Redirect(w, r, "/", 302)		
	}
}

func search(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()	
	fetchurl := r.Form["fetchurl"][0]
	data := youtube.GetVideoQuality(fetchurl)
	fmt.Fprintf(w, data)	
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
}