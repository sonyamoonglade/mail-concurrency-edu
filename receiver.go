package mailer_app

import (
	"fmt"
	"strings"
)

type Message struct {
	Label string
	Data  string
}

func FromString(str string) *Message {

	// data = "[house]: hello world"

	arr := strings.Split(str, "]:")

	label := strings.TrimPrefix(arr[0], "[")
	message := strings.TrimPrefix(arr[1], " ")

	return &Message{
		Label: label,
		Data:  message,
	}

}

type Receiver struct {
	label string
	in    chan string
}

func NewReceiver(label string) *Receiver {
	return &Receiver{
		label: label,
		in:    make(chan string),
	}
}

func (r *Receiver) Start() {
	//start should read from channel r.in and print all incoming msgs
	for data := range r.in {
		fmt.Printf("[%s]: received: %s\n", r.label, data)
	}
}

func (r *Receiver) Accept(msg string) {
	r.in <- msg
}
