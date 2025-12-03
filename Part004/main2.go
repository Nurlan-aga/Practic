/*package main

import "fmt"

type MyType struct {
	Name string
}

func (MyType) Val() {
}

func (*MyType) Ptr() {
}

type V interface {
	Val()
}

type P interface {
	Ptr()
}

type Complex interface {
	V
	P
}

func main() {
	fmt.Println("Hello, World!")

	var _ V = &MyType{}
	var _ V = MyType{}

	var _ Complex = &MyType{}
	var _ P = &MyType{}

}
*/