package main

import (
	"fmt"
	"time"
)

func hello(name string) (a, b string) {
	a = "Hello " + name
	b = "Ahoj " + name
	return
}

func main() {
	fmt.Println("Time: ", time.Now())
	a, b := hello("Vojta")
	fmt.Println(a)
	fmt.Println(b)
}
