package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func Run2() {
	var wg sync.WaitGroup
	wg.Add(1)
	resultChannel := make(chan string)
	go func() {
		Calculate2("job1", resultChannel)
		wg.Done()
	}()

	for msg := range resultChannel {
		fmt.Println(msg)
	}

	wg.Wait()
}

func Calculate2(todo string, result chan string) {
	if todo == "job1" {

	} else if todo == "job2" {

	}
	result <- "Starting "
	result <- todo
	time.Sleep(time.Second * 5)
	result <- "Ending " + todo

	// REMEBER channel.Close()
	close(result)
}
