package main

import "fmt"

type formatter interface {
	format() string
}

type plainText struct {
	message string
}

func (p plainText) format() string {
	return p.message
}

type bold struct {
	message string
}

func (b bold) format() string {
	return fmt.Sprintf("**%s**", b.message)
}

type code struct {
	message string
}

func (c code) format() string {
	return fmt.Sprintf("`%s`", c.message)
}

func sendMessage(format formatter) string {
	return format.format()
}

func main() {
	p := plainText{
		message: "Plain Text",
	}
	b := bold{
		message: "Bold Text",
	}
	c := code{
		message: "Code Text",
	}

	fmt.Println(sendMessage(p))
	fmt.Println(sendMessage(b))
	fmt.Println(sendMessage(c))
}
