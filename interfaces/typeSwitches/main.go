package main

import "fmt"

func getExpenseReport(e expense) (string, float64) {
	switch res := e.(type) {
	case email:
		return res.toAddress, res.cost()
	case sms:
		return res.toPhoneNumber, res.cost()
	default:
		return "", 0.0
	}
}

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

func main() {
	e := email{
		toAddress:    "Hello World!",
		body:         "Welcome Back",
		isSubscribed: true,
	}
	fmt.Println(getExpenseReport(e))
	fmt.Println(getExpenseReport(sms{
		isSubscribed:  false,
		body:          "Welcome Back",
		toPhoneNumber: "123456789",
	}))
	fmt.Println(getExpenseReport(invalid{}))
}
