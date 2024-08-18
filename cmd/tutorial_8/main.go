package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.RWMutex{} // this allows for locking of the goroutine so that executions can wait for each other at certain points instead of happening all at once 
var wg = sync.WaitGroup{} // this allows the program to not finish executing until the goroutines are finished, utilizing it as a counter for how many are still running
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	var t0 = time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v\n", time.Since(t0))
	fmt.Printf("The results are: %v", results)
}

// Simulate a call to a database with a set time delay
func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	m.Lock() // need to lock the goroutine so that no two threads can modify this results slice at the same time, to avoid corruption/missing data
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("The current results are: %v\n", results)
	m.RUnlock()
}