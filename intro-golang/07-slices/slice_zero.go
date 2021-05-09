package main

import "fmt"

func main() {
	langs := []string{}

	if langs == nil {
		fmt.Printf("langs is nil")
	} else {
		fmt.Printf("langs: %#v\n", langs)
	}
}
