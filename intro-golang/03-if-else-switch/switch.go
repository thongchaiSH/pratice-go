package main

import (
	"fmt"
	"time"
)

func main() {

	today := time.Monday
	fmt.Printf("Time %v\n", today)

	switch today {
	case time.Monday:
		fmt.Printf("today is Monday\n")
		fallthrough //force down case
	case time.Tuesday:
		fmt.Printf("today is Tuesday\n")
	case time.Saturday, time.Sunday: //multi case
		fmt.Printf("today is weekendy\n")
	default:
		fmt.Printf("Default day\n", today)
	}
}
