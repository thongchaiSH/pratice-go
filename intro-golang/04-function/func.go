package main

import "fmt"

func main() {
	info("Gopher", "hello", 21)
	info("Thongchai", "yahoo", 22)

	fmt.Printf("Day : %v\n", today())

	a, b := swap("1", "2")
	fmt.Println(a, b)
}

func info(name, msg string, age int) {
	fmt.Printf("My name is %v\n", name)
}

func today() string {
	return "today"
}

func swap(x, y string) (string, string) {
	return y, x
}
