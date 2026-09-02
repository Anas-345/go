package main

import "fmt"

func (e email) cost() int {
	if e.isSubscribed {
		return 2
	}
	return 5
}

func (e email) format() string {
	if e.isSubscribed {
		return fmt.Sprintf("%s | Subscribed", e.body)
	}
	return fmt.Sprintf("%s | Not Subscribed", e.body)
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}

func main() {
	e := email{
		isSubscribed: true,
		body:         "Welcome Back",
	}
	fmt.Println(e.cost())
	fmt.Println(e.format())
}
