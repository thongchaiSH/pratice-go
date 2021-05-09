package main

import "fmt"

var name string = "Gopher ทดสอบ" //package scrope

func main() {
	fmt.Printf("name : %v\n", name)
	fmt.Printf("Type %T\n", name)

	var name2 = "Type String"
	fmt.Printf("Type of name2 %T\n", name2) //%T type of variable

	name3 := "Short hand"
	fmt.Printf("Short hand %v\n", name3)

	// var s string = "Go"
	// var i int = 5
	// var f float64 = 3.7
	// var b bool = true
	// var r rune = '@'

	s := "Go"
	i := 5
	f := 3.7
	b := true
	r := '@'

	fmt.Printf("string %v\n", s)
	fmt.Printf("int %v\n", i)
	fmt.Printf("float64 %v\n", f)
	fmt.Printf("bool %v\n", b)
	fmt.Printf("rune %v\n", r)
}

func funcnum() {
	fmt.Printf("name : %v", name)
}
