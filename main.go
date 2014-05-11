// GoBy project main.go
package main

import (
	"GoBy/Exmple"
	"fmt"
	//"os"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	fmt.Println("Hello World!")
	Exmple.Atom_main()
	time.Sleep(time.Second)
}
