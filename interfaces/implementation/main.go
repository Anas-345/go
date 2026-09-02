package main

import "fmt"

func getDetails(e employee) string {
	return fmt.Sprintf("%s is earning Rs. %d", e.getName(), e.getSalary())
}

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

func main() {
	c := contractor{
		name:         "Anas",
		hourlyPay:    500,
		hoursPerYear: 100,
	}
	fmt.Println(c.getName())
	ft := fullTime{
		name:   c.getName(),
		salary: c.hourlyPay * c.hoursPerYear,
	}
	fmt.Println(getDetails(ft))
}
