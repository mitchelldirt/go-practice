package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.Mutex{}
var rwM = sync.RWMutex{}
var wg sync.WaitGroup
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func goRoutines() {
	return
	// Concurrency doesn't == parallelism
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v\n", time.Since(t0))
	fmt.Printf("The results are %v", results)
}

func dbCall(i int) {
	// Simulate DB call
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	rwM.RLock()
	fmt.Printf("\nThe current results are: %v\n", results)
	rwM.RUnlock()
}
