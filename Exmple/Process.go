package Exmple

import (
	"fmt"
	"time"
)

func Process() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("we can do evey thing")
		fallthrough
	case time.Monday:
		fmt.Println("nothing will leave!")
		fallthrough
	case time.Friday:
		fmt.Println("going crazy")
		fallthrough
	default:
		fmt.Println("working now")

	}
	t := time.Now().UTC()
	time.Sleep(time.Second)
	switch {
	case t.After(time.Now()):
		fmt.Println(t.Format("yy-mm-dd"))
		fmt.Printf("After time %d :", t)

	case t.Before(time.Now()):
		fmt.Println(t.Format("yy-mm-dd"))
		fmt.Printf("before time %d :", t)
		fallthrough
	case t.Hour() < 12:
		fmt.Println("before noon")
	case t.Hour() > 12:
		fmt.Println("after noon")
	default:
		fmt.Println(t)

	}
}
