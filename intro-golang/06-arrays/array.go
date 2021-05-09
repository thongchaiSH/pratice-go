package main

import "fmt"

func showAll(ns [4]string) {
	fmt.Printf("showAll: %#v\n", ns)
}

func main() {
	var langs = [4]string{"golang", "python", "java"}
	fmt.Printf("langs: %#v\n", langs)

	n := langs[1]
	fmt.Printf("%#v\n", n)

	langs[1] = "dart"
	fmt.Printf("langs: %#v\n", langs)

	showAll(langs)

	cords := [3]string{"Am", "Gm", "F7-11"}
	showAll(cords)

}
