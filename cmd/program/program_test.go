package program

import (
	"bytes"
	"encoding/json"
	"testing"
)

var (
	re = HttpRequest{
		Creds: Credentials{
			User:     "admin",
			Password: "Complexpass#123",
		},
		BaseURL: "localhost",
		Index:   "profiling",
		Type:    "_doc",
		Port:    "4080",
	}

	directory = "../enron_mail_20110402"
	indexer   = Indexer{
		Parser: Parser{},
	}
)

func TestIndexLargeEmail(t *testing.T) {
	em, err := indexer.Parser.Parse("../../samples/big-email.txt")
	if err != nil {
		t.Fatalf("Error parsing the file, %v", err)
	}

	if em == nil {
		t.Fatalf("The file %s is empty or does not correspond to an email.", "big-email.txt")
	}
	buf := &bytes.Buffer{}

	json.NewEncoder(buf).Encode(IndexAction{
		Index: IndexDocument{
			Index: re.Index,
		},
	})
	json.NewEncoder(buf).Encode(em)
	err = Upload(re, buf)

	if err != nil {
		t.Fatalf("error uploading the email: %v", err)
	}
}
