package main

import (
	"mailer_app"
)

const (
	HouseLabel = "house"
	WorkLabel  = "work"
)

func main() {
	mail := mailer_app.NewMail()

	h := mailer_app.NewReceiver(HouseLabel)
	w := mailer_app.NewReceiver(WorkLabel)

	go h.Start()
	go w.Start()

	mail.AddReceiver(h)
	mail.AddReceiver(w)

	stdScanner := mailer_app.NewScanner()

	mailch := mail.GetMailChannel()

	go stdScanner.Scan(mailch)

	mail.AcceptMessages()
}
