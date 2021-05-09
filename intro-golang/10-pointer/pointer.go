package main

import "fmt"

func main() {
	me := "thongchai"
	fmt.Printf("You are %v\n", me)

	// addr := &me
	var addr *string = &me

	fmt.Println(&me)
	fmt.Println(&addr)
	fmt.Printf("%T\n", addr)

	*addr = "Penguin" //derefernt
	fmt.Printf("You are %v\n", me)
}
