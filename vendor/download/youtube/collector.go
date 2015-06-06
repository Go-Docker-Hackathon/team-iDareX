package youtube

var WorkQueue = make(chan WorkRequest, 100)

func Collector(url string) {

	work := WorkRequest{ Url: url}
	
	// insert to mongo , status is in queue.
	
	WorkQueue <- work
}