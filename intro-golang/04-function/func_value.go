package main

import "fmt"

var name string = "Fluke"

// func plus(a, b int) int {
// 	return a + b
// }

// var plus = func(a, b int) int { return a + b }

func say(n string) {
	fmt.Printf("My name is %s\n", n)
}

func calc(op func(int, int) int) {
	r := op(4, 5)
	fmt.Println("Result : ", r)
}

func main() {
	fmt.Printf("value : %v\n", name)
	fmt.Printf("type : %T\n", name)

	say(name)

	plus := func(a, b int) int { return a + b }

	p := plus(1, 2)
	fmt.Printf("1+2 = %v\n", p)
	fmt.Printf("type : %T\n", plus)

	calc(plus)

	minus := func(a, b int) int { return a - b }
	calc(minus)

}
