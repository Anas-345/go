package main

import "fmt"

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

func filterMessages(messages []Message, filterType string) []Message {
	slice := []Message{}
	for _, v := range messages {
		if v.Type() == filterType {
			slice = append(slice, v)
		}
	}
	return slice
}

func main() {
	fmt.Println(filterMessages([]Message{TextMessage{Sender: "Anas", Content: "Hello"}, MediaMessage{Sender: "Anas", MediaType: "Any", Content: "Hello"}, LinkMessage{Sender: "Anas", URL: "https://google.com", Content: "Hello"}}, "text"))
}
