package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wait()
}

func wait() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			worker(id)
		}(i)
	}

	wg.Wait()
}

func worker(id int) {
	fmt.Printf("worker %d starting\n", id)
	fmt.Printf("posting to get workflow statuses\n")
	fmt.Printf("logic to check response data\n")
	fmt.Println("if data found then sleep for Xs")
	fmt.Println("if data not found then sleep for 10s and check one more time")
	fmt.Printf("worker %d done\n", id)
	sleep(10)
}

func sleep(seconds int) {
	time.Sleep(time.Second * time.Duration(seconds))
}
