package main

import "fmt"

func main() {
	pwr := make([]int, 16)
	for i := range pwr {
		pwr[i] = 2 << uint(i)
	}
	for _, value := range pwr {
		fmt.Printf("%d ", value)
	}
}
