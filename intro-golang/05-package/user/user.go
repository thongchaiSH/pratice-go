package user

import "fmt"

func Info(name, msg string, age int) {
	fmt.Printf("My name is %s\n", name)
	fmt.Printf("Message is %s\n", msg)
	fmt.Printf("Age is %d\n", age)
}
