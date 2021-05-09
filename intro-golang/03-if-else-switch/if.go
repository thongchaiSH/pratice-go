package main

import "fmt"

func main() {
	num := 14
	if num%2 == 0 {
		fmt.Printf("Even")
	} else if num == 31 {
		fmt.Printf("Number 31")
	} else {
		fmt.Printf("Odd")
	}
}
