package main

import "fmt"

type Phone interface {
	Call(number string)
}

type Samsung struct {
	Name string
}

func (s Samsung) Call(number string) {
	fmt.Println(s.Name, " calling ", number)
}

func (s Samsung) Answer() {
	fmt.Println("hello there!")
}

func Dial(p Phone) {
	p.Call("0980960604")
}

func main() {
	s := Samsung{
		Name: "Thongchai",
	}
	s.Answer()
	Dial(s)
}
