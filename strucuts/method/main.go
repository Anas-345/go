package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (a authenticationInfo) formatter() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", a.username, a.password)
}

func main() {
	s := authenticationInfo{
		username: "Anas",
		password: "123456",
	}
	fmt.Println(s.formatter())
}
