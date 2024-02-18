package stdparser

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/mail"
	"os"

	"github.com/timetravel-1010/indexer/cmd/program"
	"github.com/timetravel-1010/indexer/internal/email"
	"github.com/timetravel-1010/indexer/internal/parser"
)

var skipField map[string]bool = map[string]bool{
	"Bcc":   true,
	"Cc":    true,
	"To":    true,
	"X-Bcc": true,
	"X-Cc":  true,
	"X-To":  true,
}

type stdEmail struct{}

type StdParser struct{}

func (sp StdParser) Parse(filePath string) (*email.Email, error) {
	filePath = "hidden/sent/2."
	file, _ := os.Open(filePath)
	defer file.Close()
	reader := bufio.NewReader(file)

	msg, err := mail.ReadMessage(reader)
	if err != nil {
		return nil, err
	}
	em := &email.Email{}
	getHeaders(em, msg)

	return nil, nil
}

func getHeaders(em *email.Email, msg *mail.Message) { //mp mail.Header) {

	re := program.HttpRequest{
		Creds: program.Credentials{
			User:     "admin",
			Password: "Complexpass#123",
		},
		BaseURL: "localhost",
		Index:   "profiling",
		Type:    "_doc",
		Port:    "4080",
	}

	m := map[string]any{}
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	for k, lv := range msg.Header {
		if skipField[k] {
			m[k], _ = msg.Header.AddressList(k)
			continue
		}
		m[k] = lv[0]
	}
	//body, _ := io.ReadAll(msg.Body)
	//buf := new(bytes.Buffer)
	io.Copy(buf, msg.Body)
	body := buf.String()
	m["body"] = string(body)
	foo, _ := json.Marshal(m)
	fmt.Printf("marshaled: \n%s", foo)

	encoder.Encode(program.IndexAction{
		Index: program.IndexDocument{
			Index: "foo",
		},
	})

	encoder.Encode(parser.Document2{
		Path:  "somepath",
		Email: m,
	})
	program.Upload(re, buf)
}
