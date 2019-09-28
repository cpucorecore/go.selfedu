package main

import "fmt"

var (
	i = 10
	s = "abc"
)

func main() {
	fmt.Println("Hello, World")
	fmt.Println(i, s, "\n")
	i = 100
	s = "edf"
	fmt.Println(i, s, "\n")
	fmt.Println("\n")
}
