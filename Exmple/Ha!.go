package Exmple

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func Atom_main() {
	var ops uint64 = 0

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(atomic.AddUint64(&ops, 1))
			runtime.Gosched()
		}()
	}
	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println(opsFinal)
}
