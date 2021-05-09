package main

import "fmt"

func main() {
	fmt.Print("Hello Gopher!!")

	fmt.Printf("Hello %s!!\n", "Thongchai") //string
	fmt.Println("next line")

	fmt.Printf("Hello %d", 1000) //number

	fmt.Printf("Hello %v\n", "Thongchai") //dynamic type
	fmt.Printf("Age %v\n", "26")

	fmt.Printf("Hello %v : %v\n", "Thongchai", 26) //multi parameter

}
