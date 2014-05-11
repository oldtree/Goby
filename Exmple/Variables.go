package Exmple

import (
	"fmt"
	"strconv"
)

const (
	FIRST = iota
	SECOND
	THRID
	FORTH
)

type I interface {
	Show(key string)
}
type S struct {
}
type f func(info string) string

type value2func map[string]func(info string) string

func (v value2func) Show(key string) {
	value := v[key]
	fmt.Printf("Key : %s == Value :%d \n", key, value)
}

var vv value2func

var info = "o.o"

var list []int
var list2 [10]int

func What(info string) string {
	fmt.Println(info)
	return info
}

func Show() {
	vv = make(map[string]func(info string) string)
	for i := 0; i < 20; i++ {
		F := func(info string) string {
			fmt.Println(info)
			return info
		}
		vv[strconv.Itoa(i)] = F
		list = append(list, i)
	}
	fmt.Println(len(list2))
	fmt.Println(len(list))
	fmt.Println(len(vv))
	FF := vv["1"]
	vv.Show("2")
	FF("test")
}
