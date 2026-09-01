package main

import "fmt"

func (u User) SendMessage(msg string, msgLen int) (message string, success bool) {
	success = false
	if msgLen <= u.MessageCharLimit {
		success = true
		message = msg
	}
	return
}

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	membership := Membership{Type: membershipType}
	if membershipType == "premium" {
		membership.MessageCharLimit = 1000
	} else {
		membership.Type = "standard"
		membership.MessageCharLimit = 100
	}
	return User{Name: name, Membership: membership}
}

func main() {
	u1 := newUser("Anas", "premium")
	fmt.Println(u1.SendMessage("Hello!!", 7))
}
