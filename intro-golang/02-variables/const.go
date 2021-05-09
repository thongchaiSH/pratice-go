package main

import "fmt"

func main() {
	var day string = "Monday"
	fmt.Printf("day:%v\n", day)

	day = "Tues day"
	fmt.Printf("day:%v\n", day)

	const dayConst = "Monday"
	fmt.Printf("Const day : %v", dayConst)

	// const Sunday = 0
	// const Monday = 1
	// const Tuesday = 2
	// const Wednesday = 3
	// const Thursday = 4
	// const Friday = 5
	// const Saturday = 6
	const (
		Sunday = iota //auto assigned +1 start 0
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Printf("Sunday : %v\n", Sunday)
	fmt.Printf("Monday : %v\n", Monday)
	fmt.Printf("Tuesday : %v\n", Tuesday)
	fmt.Printf("Wednesday : %v\n", Wednesday)
	fmt.Printf("Thursday : %v\n", Thursday)
	fmt.Printf("Friday : %v\n", Friday)
	fmt.Printf("Saturday : %v\n", Saturday)
}
