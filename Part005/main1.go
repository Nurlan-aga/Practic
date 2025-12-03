/*package main

import "fmt"

type Greeter interface {
	Greet()
}

type User struct {
	Name string
}

func (u User) Greet() {
	fmt.Println("Hello, I am", u.Name)
}

type Admin struct {
	User
	Level int
}

func SayGreet(g Greeter) {
	g.Greet()
}

func main() {
	u := User{Name: "User1"}
	a := Admin{User: User{Name: "Admin1"}, Level: 10}

	SayGreet(u)
	SayGreet(a) // Admin реализует Greeter через встроенный User
}
*/