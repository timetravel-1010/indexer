package stdparser

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"log"
	"net/mail"
	"os"

	"github.com/timetravel-1010/indexer/internal/email"
)

var listField map[string]bool = map[string]bool{
	"Bcc":   true,
	"Cc":    true,
	"To":    true,
	"X-Bcc": true,
	"X-Cc":  true,
	"X-To":  true,
}

type StdEmail map[string]any

func (se StdEmail) DoSomething() {}

type StdParser struct{}

func (sp StdParser) Parse(filePath string) (email.EmailI, error) {
	file, _ := os.Open(filePath)
	defer file.Close()
	reader := bufio.NewReader(file)

	msg, err := mail.ReadMessage(reader)
	if err != nil {
		log.Println("error in ReadMessage")
		return nil, err
	}
	return getStdEmail(msg)
}

func getStdEmail(msg *mail.Message) (*StdEmail, error) {
	m := StdEmail{}
	buf := &bytes.Buffer{}
	var err error

	for k, lv := range msg.Header {
		if listField[k] {
			if _, ok := m[k]; !ok {
				continue
			}
			m[k], err = msg.Header.AddressList(k)

			if err != nil {
				if errors.Is(err, mail.ErrHeaderNotPresent) {
					continue
				}
				continue
			}
			continue
		}
		m[k] = lv[0]
	}
	io.Copy(buf, msg.Body)
	m["body"] = buf.String()

	return &m, nil
}
