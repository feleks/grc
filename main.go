package main

import (
	"net/http"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// HttpServer
	wg.Add(1)
	go func() {
		setupRoutes()
		err := http.ListenAndServe(":8971", nil)
		if err != nil {
			panic(err)
		}
	}()

	wg.Wait()

	panic("All goroutines terminated")
}
