package main

import (
	ct "testPackages/counter"
	"testPackages/log"
)

func main() {
	l := log.NewLogger()
	if l == nil {
		panic("logger not initialized")
	}
	l.Log("counter package initialized in logs")
	ct.Count()
	ct.Print()
	ct.Count()
	ct.Print()
	ct.Count()
	ct.Print()
	ct.Count()
	ct.Print()
}
