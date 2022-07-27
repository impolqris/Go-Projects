package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter your URL!")
	fmt.Println("------------------")
	fmt.Printf("-> ")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	fmt.Printf("\n   %d characters.\n", len(text)-1)

	if len(text)-1 > 2048 || len(text)-1 < 4 {
		fmt.Printf("\n[WARNING] Heads up, your text is not suitable for any valid URL.")
	}
}
