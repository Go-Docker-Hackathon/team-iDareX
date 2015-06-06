package youtube

import (
	"fmt"
	"net/http"
	"html/template"
)

var WorkQueue = make(chan WorkRequest, 100)

func Collector (w http.ResponseWriter, r *http.Request) {
	
	// make sure we can only be called with an HTTP POST request.
	if r.Method != "POST" {
		t, _ := template.ParseFiles("./collectorLinkForm.gtpl")
		t.Execute(w, nil)
		return
	}
	
	url := r.FormValue("url")
	
	if url == "" {
		http.Error(w, "You must specify a url.", http.StatusBadRequest)
		return
	}
	
	work := WorkRequest{ Url: url}
	
	// insert to mongo , status is in queue.
	
	WorkQueue <- work
	fmt.Println("work request queued")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Url: " + url + " 已经收，服务器正在努力处理！")
	return
}