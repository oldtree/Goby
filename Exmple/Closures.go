package Exmple

import (
	"fmt"
)

func Out(num int) func() int {
	i := num
	return func() int {
		if i == 0 {
			return 1
		}
		return i - 1
	}
}

func Closuer_Main() {
	in := Out(10)
	fmt.Println(in())
	fmt.Println(in())
	fmt.Println(in())
	fmt.Println(in())
	fmt.Println(in())

}
