package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func (p Point) DistanceFromOrigin() int {
	if p.X < 0 {
		p.X = -p.X
	}
	if p.Y < 0 {
		p.Y = -p.Y
	}
	return p.X + p.Y
}

func (p *Point) Move(dx, dy int) {
	p.X += dx
	p.Y += dy
}

type Greeter interface {
	Greet() string
}

type Person struct {
	Name string
}

func (p Person) Greet() string {
	return "Hello, human " + p.Name
}

type Robot struct {
	ID string
}

func (r Robot) Greet() string {
	return "Hello, robot " + r.ID
}

func SayHello(g Greeter) {
	fmt.Println(g.Greet())
}

type Updater interface {
	Update(msg string)
}

// реализуй интерфейс по указателю
type Logger struct {
	Last string
}

func (l *Logger) Update(msg string) {
	l.Last = msg
}

type User struct {
	Name string
}

type Greeter1 interface {
	Greet()
}

func (u User) Greet() {
	fmt.Println("Hello, I am", u.Name)
}

type Admin struct {
	User
	Level int
}

func (a Admin) Ban(userName string) {
	fmt.Println(a.Name, "bans", userName) // Name из встраиваемого User
}

func SayGreet(g Greeter1) {
	g.Greet()
}

type Reader interface {
	Read() string
}

type Writer interface {
	Write(s string)
}

type ReadWriter interface {
	Reader
	Writer
}

type MemoryBuffer struct {
	data string
}

func (m *MemoryBuffer) Write(s string) {
	m.data = s
}

func (m *MemoryBuffer) Read() string {
	return m.data
}

func Describe(v any) {
	switch val := v.(type) {
	case int:
		fmt.Println("int:", val)
	case string:
		fmt.Println("string:", val)
	case bool:
		fmt.Println("bool:", val)
	default:
		fmt.Printf("unknown (%T): %v\n", val, val)
	}
}

func main() {
	//1
	fmt.Println("==============================================")
	fmt.Println("Task1")
	p := Point{X: -3, Y: 5}
	fmt.Println("Distance:", p.DistanceFromOrigin())

	//2
	fmt.Println("==============================================")
	fmt.Println("Task2")
	p = Point{X: 1, Y: 2}
	p.Move(3, -1)
	fmt.Println("Point:", p.X, p.Y) // 4 1

	//3
	fmt.Println("==============================================")
	fmt.Println("Task3")
	p1 := Person{Name: "Фёдор"}
	SayHello(p1)

	//4
	fmt.Println("==============================================")
	fmt.Println("Task4")
	p2 := Person{Name: "Фёдор"}
	r1 := Robot{ID: "RX-78"}
	SayHello(p2)
	SayHello(r1)

	//5
	fmt.Println("==============================================")
	fmt.Println("Task5")
	var greeters []Greeter = []Greeter{
		Person{Name: "Vasya"},
		Person{Name: "Petya"},
		Robot{ID: "Robocop"},
		Robot{ID: "Wall-E"},
	}
	for _, g := range greeters {
		fmt.Println(g.Greet())
	}

	//6
	fmt.Println("==============================================")
	fmt.Println("Task6")
	//var u1 Updater
	var u2 Updater

	l := &Logger{}

	// u1 = l      // <- так не компилируется
	u2 = l // <- а так компилируется
	u2.Update("hello")
	fmt.Println("Last:", l.Last)

	//7
	fmt.Println("==============================================")
	fmt.Println("Task7")
	a := Admin{
		User:  User{Name: "SuperAdmin"},
		Level: 10,
	}

	a.Greet() // метод User "поднялся" на уровень Admin
	a.Ban("troll42")

	//8
	fmt.Println("==============================================")
	fmt.Println("Task8")
	u := User{Name: "User1"}
	a1 := Admin{User: User{Name: "Admin1"}, Level: 10}

	SayGreet(u)
	SayGreet(a1) // Admin реализует Greeter через встроенный User

	//9
	fmt.Println("==============================================")
	fmt.Println("Task9")
	var rw ReadWriter = &MemoryBuffer{}
	rw.Write("test")
	fmt.Println(rw.Read())

	//10
	fmt.Println("==============================================")
	fmt.Println("Task10")
	Describe(10)
	Describe("hello")
	Describe(true)
	Describe(3.14)
}
