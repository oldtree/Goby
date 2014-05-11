package Exmple

import (
	"fmt"
	"time"
)

func NonBlock_main() {
	messges := make(chan string, 1)
	signals := make(chan int, 1)

	//messges := make(chan string)
	//signals := make(chan int)

	select {
	case msg := <-messges:
		fmt.Println("recevied message ", msg)
	default:
		fmt.Println("no message received ")
	}

	msg := "hi"

	select {
	case messges <- msg:
		fmt.Println("sent message ", msg)
	default:
		fmt.Println("no messge sent ")
	}

	select {
	case msg := <-messges:
		fmt.Println("recevied message ", msg)
	case sig := <-signals:
		fmt.Println("recevied signals ", sig)
	default:
		fmt.Println("no activity")
	}

	fmt.Println("-------------------------------------------------------")

}

func ChannelOp2() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("recevied job", j)
			} else {
				fmt.Println("recevied all job !!")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 5; i++ {
		jobs <- i
		if i == 3 {
			//close(jobs) panic: runtime error: send on closed channel
		}
		fmt.Println("sent jobs ")
	}

	close(jobs)
	for elem := range jobs {
		fmt.Println("get from closed channel : ", elem)
	}
	<-done
	fmt.Println("sent all  jobs ")
	fmt.Println("-------------------------------------------------------")
}

func ChanOpTime() {
	timer := time.NewTimer(time.Second * 3)

	fmt.Println(<-timer.C)
	fmt.Println("Timer 1 has expired  ")

	timer2 := time.NewTimer(time.Second * 1)
	go func() {
		fmt.Println(<-timer2.C)
		fmt.Println("Timer 2 has expired  ")
	}()
	time.Sleep(time.Second * 2)
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 has stoped  ")
	}

	ticker := time.NewTicker(time.Millisecond * 100)
	//ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for st := range ticker.C {
			fmt.Println("tick at ", st)
		}
	}()
	time.Sleep(time.Second * 2)
	ticker.Stop()
	fmt.Println("ticker stoped")

}
