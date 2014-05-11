package Exmple

import (
	//"errors"
	"fmt"
	"time"
)

type Error struct {
	info  string
	level int
}

func (ee Error) String() string {
	return fmt.Sprintf("this is %d level error : %x ", ee.level, ee.info)
}

func Ch_Main() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	mesg := <-messages
	fmt.Println(mesg)
}

var messages = make(chan string, 2)

func Ch_Buf_main() {
	messages <- "first"
	messages <- "second"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

func ping(pings chan string, msg string) {
	pings <- msg
}

func pong(pings chan string, pongs chan string) {
	pongs <- <-pings
}
func PingPong() {
	fmt.Println("begin ping pong ")

	pin := make(chan string, 1)
	pon := make(chan string, 2)
	msg := "GZ hello"
	ping(pin, msg)
	pong(pin, pon)
	fmt.Println(<-pon)
	fmt.Println("end ping pong ")
}

func Multi() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	b1 := make(chan bool, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "i am the first "
	}()

	go func() {
		time.Sleep(time.Microsecond * 1)
		c2 <- "i am the second "
	}()
	tt := time.Now()
	//for {
	b1 <- (time.Now().Minute() > tt.Minute()+1)
	//fmt.Println((time.Now().Minute() > tt.Minute()+1))
	select {
	case msg1 := <-c1:
		fmt.Println("recivied :", msg1)
	case msg2 := <-c2:
		fmt.Println("recivied :", msg2)
	case <-time.After(time.Second * 2):
		fmt.Println(" i am out of the loop")

		//	}
	}

}
