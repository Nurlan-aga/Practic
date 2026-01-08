package main

import "fmt"

type IReader interface {
	Read() string
}

type IWriter interface {
	Write(s string)
}

type IReadWriter interface {
	IReader
	IWriter
}

type Memory struct {
	data string
}

func (m *Memory) Read() string {
	return m.data
}

func (m *Memory) Write(s string) {
	m.data = s
}

func main() {
	var a IReadWriter = &Memory{}
	a.Write("qqq")
	fmt.Println(a.Read())
}
