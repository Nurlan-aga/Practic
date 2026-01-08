package main

import (
	"fmt"

	"github.com/Nurlan-aga/greeting"
	"github.com/google/uuid"
)

func main() {
	fmt.Println(uuid.New().String())
	fmt.Println(greeting.Greet())
	fmt.Println(greeting.Bye())
	fmt.Println(greeting.Add())
}
