package Exmple

import (
	"fmt"
	"math"
)

func Values_main() {
	fmt.Println("7.0 / 3.0  ", 7.0/3.0)
	fmt.Println(^uint64(1))
	fmt.Println(^0)
	fmt.Println(1 << 2)
	fmt.Println(1 > 2)
	fmt.Printf("max value of MaxFloat32  :%d\n", math.MaxFloat32)
	fmt.Printf("max value of MaxFloat64  :%d\n", math.MaxFloat64)
	fmt.Printf("max value of MaxInt64   :%d\n", math.MaxInt64)
	fmt.Printf("min value of MinInt16  :%d\n", math.MinInt64)
	fmt.Printf("max value of MaxInt32  :%d\n", math.MaxInt32)
	fmt.Printf("min value of MinInt16   :%d\n", math.MinInt32)
	fmt.Printf("max value of MaxInt16  :%d\n", math.MaxInt16)
	fmt.Printf("min value of MinInt16   :%d\n", math.MinInt16)
	fmt.Printf("max value of MaxInt8  :%d\n", math.MaxInt8)
	fmt.Printf("min value of MinInt16  :%d\n", math.MinInt8)

	fmt.Println(!(false && true) && (true && true) || (false) || (true))
}
