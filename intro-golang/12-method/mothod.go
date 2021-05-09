package main

import "fmt"

type User struct {
	Username          string
	FullName          string
	ProfilePictureUtl string
}

//info is method of User
func (u User) Info() {
	fmt.Printf("Username : %v\n", u.Username)
	fmt.Printf("FullName : %v\n", u.FullName)
	fmt.Printf("ProfilePictureUtl : %v\n", u.ProfilePictureUtl)
}

func main() {
	//instance of user
	u := User{
		Username:          "thongchai",
		FullName:          "FF SS",
		ProfilePictureUtl: "google.com",
	}

	//info(u)
	u.Info()
}
