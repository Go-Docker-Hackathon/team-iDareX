package youtube

import "fmt"

var WorkerQueue chan chan WorkRequest

func StartDispatcher(nworkers int) {
	// First, initialize the channel we are going to but the workers' work channels into.
	WorkerQueue := make(chan chan WorkRequest, nworkers)
	
	// Now, create all of our workers.
	for i := 1; i<=nworkers; i++ {
		fmt.Println("Starting worker", i)
		worker := NewWroker(i, WorkerQueue)
		worker.Start()
	}
	
	go func() {
		for {
			select {
			case work := <-WorkQueue:
				fmt.Println("Received work request")
				go func() {
					worker := <-WorkerQueue
					fmt.Println("Dispatching work request ")
					worker <- work
				}()
			}
		}
	}()
}