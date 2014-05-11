package Exmple

import (
	"fmt"
	"reflect"
)

var s []string

func Slice_main() {
	fmt.Println(reflect.TypeOf("A"))
	s = append(s, "A", "B", "C")
	s = append(s, "E", "F", "G")

	c := make([]string, len(s[3:]))
	copy(c, s[3:])
	fmt.Println(len(s[3:]))
	for str := range s {
		fmt.Printf("S  this is %x \n", str)
	}
	for str := range c {
		fmt.Printf("C  this is %x \n", str)
	}
}
