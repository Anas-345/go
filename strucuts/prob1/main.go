package main

import "fmt"

type User struct {
	MemberShip
	Name string
}

type MemberShip struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	lim := 100
	if membershipType == "premium" {
		lim = 1000
	}
	return User{Name: name, MemberShip: MemberShip{
		Type:             membershipType,
		MessageCharLimit: lim,
	}}
}

func main() {
	u1 := newUser("Anas", "basic")
	fmt.Println(u1.Name, u1.MessageCharLimit, u1.Type)
}
