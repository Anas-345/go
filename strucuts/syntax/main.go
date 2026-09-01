package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message     string
}

func main() {
	msg1 := messageToSend{
		phoneNumber: 123456,
		message:     "Welcome Back",
	}
	fmt.Println(msg1)
}
