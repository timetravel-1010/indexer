package main

import (
	"github.com/timetravel-1010/indexer/cmd/program"
	"testing"
)

var (
	re = program.HttpRequest{
		Creds: program.Credentials{
			User:     "admin",
			Password: "Complexpass#123",
		},
		BaseURL: "localhost",
		Index:   "profiling",
		Type:    "_doc",
		Port:    "4080",
	}

	directory = "../enron_mail_20110402"
	indexer   = program.Indexer{
		Parser: program.Parser{},
	}
)

func BenchmarkXxx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		indexer.Index(directory, re)
	}
}
