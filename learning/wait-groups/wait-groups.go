package main

import (
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
	sleep(10)
}

func sleep(seconds int) {
	time.Sleep(time.Second * time.Duration(seconds))
}
