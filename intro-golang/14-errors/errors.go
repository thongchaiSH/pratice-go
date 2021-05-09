package main

import (
	"fmt"
	"log"
)

func ReadFile(name string) (string, error) {
	//read file
	// return "", errors.New("File not found") //err
	return "data..", nil
}

func main() {
	data, err := ReadFile("profile.txt")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("read file success", data)
}
