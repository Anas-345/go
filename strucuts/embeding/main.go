package main

import "fmt"

type user struct {
	name   string
	number int
}

type sender struct {
	user
	rateLimit int
}

func main() {
	s := sender{
		rateLimit: 10,
		user: user{
			name:   "Anas",
			number: 123456,
		},
	}
	fmt.Println(s.name, s.number, s.rateLimit)
}
