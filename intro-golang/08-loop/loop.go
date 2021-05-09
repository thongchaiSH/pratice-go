package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// i := 0
	//loop forever
	// for {
	// 	fmt.Println(i)
	// }
	langs := []string{"golang", "python", "java"}
	for i := 0; i < len(langs); i++ {
		fmt.Println(i, ":", langs[i])
	}

	fmt.Println("For range slice")
	// for _, value := range langs { //value only
	// for index := range langs { //index only
	for index, value := range langs {
		fmt.Println(index, ":", value)
	}
}
