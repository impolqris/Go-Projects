package main

import (
	"fmt"
)

func main() {
	i := ""

	fmt.Printf("What is your name? ")
	fmt.Scanln(&i)
	fmt.Println("Your name: ", i)
}
