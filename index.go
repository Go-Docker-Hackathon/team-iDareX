package main

import (
	"net/http" 
	"html/template" 
	"fmt"
	"log"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	//	"strings"
	//	"os" 
)

// 任务的结构
type Task struct { 
	Fetchurl string
	Downloadurl string
	Status int
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/addurl", addurl)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {		
		session, err := mgo.Dial("localhost")
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)		
		c := session.DB("docker").C("task")
		
		var tasks []Task
		c.Find(nil).All(&tasks) //查询所有
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

		fetchurl := r.Form["fetchurl"][0] // 

		fmt.Println("url : = " + fetchurl)

		session, err := mgo.Dial("localhost")
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)		
		c := session.DB("docker").C("task") 
		
		task := Task{} 
		err = c.Find(bson.M{"Fetchurl": fetchurl}).One(&task)

		if err != nil {
			fmt.Println("任务已经在下载队列中")
		} else {
			err = c.Insert(&Task{fetchurl, "", 0}) 
			checkError(err)				
		}		 
		http.Redirect(w, r, "/", 302)	
		
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
}