package main

import (
	"fmt"
	"github/thongchaiSH/igapp/time"
	"github/thongchaiSH/igapp/user"
)

//package user
// func info(name, msg string, age int) {
// 	fmt.Printf("My name is %s\n", name)
// 	fmt.Printf("Message is %s\n", msg)
// 	fmt.Printf("Age is %d\n", age)
// }

//package time
// func today() string {
// 	return "today"
// }

func main() {
	user.Info("Thongchai", "Hi", 26)

	t := time.Today()
	fmt.Printf("today is %s", t)

}
