package program

import (
	"bytes"
	"encoding/json"
	"io/fs"
	"log"
	"path/filepath"
)

// Indexer
type Indexer struct {
	Parser Parser
	path   string
}

// Index
func (in *Indexer) Index(dir string, re HttpRequest) {
	var counter int = 0
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	emails := []Document{}

	log.Println("Indexing documents...")
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			em, err := in.Parser.Parse(path)
			if err != nil {
				return err
			}
			emails = append(emails, Document{Path: path, Email: em})
			counter++
			if counter == 100 {
				encoder.Encode(Payload{Index: re.Index, DocumentData: emails})
				Upload(re, buf)
				buf.Reset()
				counter = 0
			}
		}
		return nil
	})
	if err != nil {
		panic("Error while opening the files!")
	}
	if counter > 0 {
		encoder.Encode(Payload{Index: re.Index, DocumentData: emails})
		Upload(re, buf)
	}
	log.Println("Indexing completed successfully completed.")
}
