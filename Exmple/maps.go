package Exmple

import (
	"fmt"
)

var m = map[int]string{1: "one", 2: "two", 3: "three"}

func Map_main() {
	mp := make(map[string]int)
	mp["one"] = 1
	mp["two"] = 2
	fmt.Printf("length of mp %d \n", len(mp))
	delete(mp, "one")
	fmt.Printf("length of mp %d \n", len(mp))

	for k, v := range m {
		fmt.Printf("%s<---->%s  \n", k, v)
	}
	for i, c := range "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmeopqrstuvwxyz" {
		fmt.Println(i, c)
	}
	li := []int{1, 2, 3, 4, 5}
	ShowChar(li...)
}

func ShowChar(chars ...int) {
	for _, char := range chars {
		fmt.Println(char)
	}
}
