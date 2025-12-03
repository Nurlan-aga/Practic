/*package main

import "fmt"

type Greeter interface {
	Greet() string
}

type User struct {
	Name string
}

func (u User) Greet() string {
	return "Hello, " + u.Name
}

type Android struct {
	Model string
}

func (a Android) Greet() string {
	return "Hi, I am Android " + a.Model
}

func SayHello(g Greeter) string {
	return g.Greet()
}

func main() {

	var SomeGreeter Greeter

	SomeGreeter = User{Name: "Bob"}

	SomeGreeter = User{Name: "Alice"}
	greet := SayHello(SomeGreeter)
	fmt.Println(greet)
	SomeGreeter = Android{Model: "X1000"}
	greet = SayHello(SomeGreeter)
	fmt.Println(greet)
}
*/