package main

import "fmt"

type User struct { //similar to class
	Username          string
	FullName          string
	ProfilePictureUtl string
}

func main() {
	// username := "thongchai"
	// fullname := "FF SS"
	// ProfilePictureUtl := "https://xxx"

	u := User{
		Username:          "thongchai",
		FullName:          "FF SS",
		ProfilePictureUtl: "google.com",
	}

	fmt.Printf("%#v\n", u)
	fmt.Printf("%v\n", u.Username)
}
