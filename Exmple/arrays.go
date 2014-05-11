package Exmple

import (
	"fmt"
)

var a [5]int
var b = [10]int{1, 2, 3, 4, 5}
var c [10][10]int

func ArrayMain() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			c[i][j] = i + j
		}
	}
	fmt.Println(c)
}
