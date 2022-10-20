package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func Run() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		Calculate("job1")
		wg.Done()
	}()
	go func() {
		Calculate("job2")
		wg.Done()
	}()
	f := func() {
		Calculate("job3")
		wg.Done()
	}
	go f()

	wg.Wait()
}

func Calculate(todo string) {
	if todo == "job1" {

	} else if todo == "job2" {

	}
	fmt.Println("Starting ", todo)
	time.Sleep(time.Second * 5)
	fmt.Println("Ending ", todo)
}
