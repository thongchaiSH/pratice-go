package main

import "fmt"

func main() {
	langs := []string{"golang", "python", "java"}
	fmt.Printf("langs: %#v\n", langs)

	fmt.Printf("%#v\n", langs[0])

	fmt.Printf("length %#v\n", len(langs)) //get size

	langs = append(langs, "F#")             //add item
	langs = append(langs, "F#", "Em", "Am") //add multi item
	fmt.Printf("langs: %#v\n", langs)
}
