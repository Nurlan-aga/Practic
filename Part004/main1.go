/*package main

import "fmt"

type Greeter interface {
	Greet() string
}

type User struct {
	Name string
}

func (u User) Greet() string {
	return fmt.Sprintf("hello %s", u.Name)
}

func PrintGreeter(g Greeter) {
	u, ok := g.(User)

	if ok {
		fmt.Println(u.Greet())
	}
}

func PrintType(x any) {
	switch v := x.(type) {
	case int:
		fmt.Printf("This is an int: %d\n", v)
	case string:
		fmt.Printf("This is a string: %s\n", v)
	case Greeter:
		fmt.Printf("This is a Greeter: %s\n", v.Greet())
	default:
		fmt.Println("Unknown type")
	}
}

func main() {

	var u User = User{Name: "Alice"}
	PrintGreeter(u)

	PrintType(42)
	PrintType("Hello, World!")
	PrintType(u)

}
*/