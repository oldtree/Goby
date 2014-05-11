package Exmple

import (
	"fmt"
	"math"
	"reflect"
)

type Person struct {
	name string
	age  int
}

type Geometry interface {
	area() float32
	perim() float32
}

type rect struct {
	length float32
	width  float32
}

type circle struct {
	radius float32
}

func (r rect) area() float32 {
	return r.length * r.width
}
func (r rect) perim() float32 {
	return 2*r.length + 2*r.width
}

func (ce circle) area() float32 {
	return ce.radius * math.Pi * ce.radius
}

func (ce circle) perim() float32 {
	return ce.radius * 2 * math.Pi
}

func Shape(ge Geometry) {
	fmt.Println("this geometry info :", reflect.TypeOf(ge))
	fmt.Printf("the area is : %d \n", ge.area())
	fmt.Printf("the perim is : %d \n", ge.perim())
}

func Struct_Main() {
	fmt.Println(Person{"Bob", 20})
	fmt.Println(Person{name: "Green", age: 40})
	fmt.Println(Person{name: "linux"})
	fmt.Println(Person{age: 30})
	fmt.Println(Person{})
	fmt.Println(reflect.TypeOf(Person{}))

	fmt.Println("------------------------------------")

	T := rect{10.0, 10.0}
	Shape(T)

	C := circle{10.0}
	Shape(C)

}
