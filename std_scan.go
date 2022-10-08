package mailer_app

import (
	"bufio"
	"os"
	"strings"
)

type StdScan struct {
	scanner *bufio.Scanner
}

func NewScanner() *StdScan {
	return &StdScan{
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (s *StdScan) Scan(mails chan string) {

	for s.scanner.Scan() {
		data := s.scanner.Text()
		end := strings.Contains(data, "end")
		if end {
			close(mails)
			return
		}
		mails <- data
	}

}
