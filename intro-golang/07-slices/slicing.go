package main

import "fmt"

func main() {
	langs := []string{"golang", "python", "java"}
	fmt.Printf("langs: %#v\n", langs)

	a := langs[0:2] //halfopen [0-2) not 2
	fmt.Printf("langs: %#v\n", a)

	a = langs[1:3] //[1:3)
	fmt.Printf("langs: %#v\n", a)

	a = langs[0:len(langs)] //langs[0:]
	fmt.Printf("langs: %#v\n", a)

	r := langs[0:3]
	s := langs[:3]
	t := langs[0:]
	u := langs[:]

	fmt.Printf("r : %v\n", r)
	fmt.Printf("s : %v\n", s)
	fmt.Printf("t : %v\n", t)
	fmt.Printf("t : %v\n", u)

	printSlice(langs)
	cords := [4]string{"A", "B", "C", "D"}
	printSlice(cords[:]) //convert array to slice
}

func printSlice(ns []string) {
	fmt.Printf("printSlice : %v\n", ns)
}
