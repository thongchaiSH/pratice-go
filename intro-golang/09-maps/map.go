package main

import "fmt"

func main() {
	status := map[int]string{
		200: "OK",
		308: "Permanet Redirect",
		400: "Bad Request",
		500: "Internal Server Error",
	}
	fmt.Printf("%#v\n", status)
	fmt.Printf("size %#v\n", len(status))

	status[200] = "Okie"   //change
	status[600] = "Custom" //append

	fmt.Printf("%#v\n", status)
	fmt.Printf("size %#v\n", len(status))

	// v, ok := status[9999]
	if v, ok := status[9999]; ok {
		fmt.Printf("Read  %#v\n", v)
	} else {
		fmt.Printf("Key is empty")
	}

}
