package mailer_app

import "fmt"

type Mail struct {
	mailch    chan string
	receivers map[string]*Receiver
}

func (m *Mail) GetMailChannel() chan string {
	return m.mailch
}

func NewMail() *Mail {
	return &Mail{
		mailch:    make(chan string),
		receivers: make(map[string]*Receiver),
	}
}

func (m *Mail) AddReceiver(r *Receiver) {
	m.receivers[r.label] = r
}

func (m *Mail) AcceptMessages() {

	for {

		data, ok := <-m.mailch
		if !ok {
			return
		}

		msg := FromString(data)

		rcv, ok := m.receivers[msg.Label]
		if !ok {
			fmt.Println("unknown receiver:", msg.Label)
			continue
		}

		rcv.Accept(msg.Data)
	}

}
