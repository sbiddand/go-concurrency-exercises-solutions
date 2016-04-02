package main

import (
	"sync"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	fetchSig := fetchSignalInstance()
	done := make(chan bool)

	start := time.Now()
	go func() {
		for {
			switch {
			case <-fetchSig:
				// Check if signal arrived earlier than a second (with error margin)
				if time.Since(start).Nanoseconds() < 990000000 {
					t.Log("There exists a two crawls who were executed less than 1 sec apart.")
					t.Log("Solution is incorrect.")
					t.FailNow()
				}
				start = time.Now()
			case <-done:
				break
			}
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	Crawl("http://golang.org/", 4, &wg)
	wg.Wait()
	done <- true
}
