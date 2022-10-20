package concurrency

import (
	"fmt"
	"strconv"
	"time"
)

func Run4() int {
	return 23
}

func Run3() {
	resultChannel := make(chan string)
	workChannel := make(chan string, 100)

	for i := 0; i < 100; i++ {
		workChannel <- "Jobb" + strconv.FormatInt(int64(i), 10)
	}

	go Calculate3(workChannel, resultChannel)
	go Calculate3(workChannel, resultChannel)
	go Calculate3(workChannel, resultChannel)
	go Calculate3(workChannel, resultChannel)

	for msg := range resultChannel {
		fmt.Println(msg)
	}

}

func Calculate3(workChannel chan string, result chan string) {
	for {
		nextJob := <-workChannel

		result <- "Starting "
		result <- nextJob
		time.Sleep(time.Second * 5)
		result <- "Ending " + nextJob

	}
}
