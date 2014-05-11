package Exmple

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker : ", id, " processing job")
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func Pool_main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 9; a++ {
		//fmt.Println(<-results)
		//<-results
		<-results
	}
}

func Rate_main() {
	requests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		requests <- i
	}

	close(requests)

	limiter := time.Tick(time.Microsecond * 200)

	for req := range requests {
		<-limiter
		fmt.Println("tick request ", req, time.Now())
	}

	bursylimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		bursylimiter <- time.Now()
	}

	go func() {
		for ttt := range time.Tick(time.Microsecond * 200) {
			bursylimiter <- ttt
		}
	}()

	bursyrequest := make(chan int, 5)
	for i := 0; i < 5; i++ {
		bursyrequest <- i
	}
	close(bursyrequest)

	for req := range bursyrequest {
		//<-bursyrequest //holy shit
		fmt.Println("times request ", req, time.Now())
	}
}
