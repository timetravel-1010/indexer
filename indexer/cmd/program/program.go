package program

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"

	"github.com/timetravel-1010/indexer/cmd/util"
	"github.com/timetravel-1010/indexer/internal/parser"
)

// Indexer
type Indexer struct {
	Parser parser.ParserI
	path   string
}

// Index
func (in *Indexer) Index(dir string, re HttpRequest) {
	var counter int = 0
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)

	log.Println("Indexing documents...")

	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			if isEmpty, err := util.CheckEmpty(path); isEmpty && err == nil {
				fmt.Printf("The file %s is empty.\n", path)
				return nil
			} else if err != nil {
				return errors.New(fmt.Sprintf("error checking empty file: %s", err.Error()))
			}

			em, err := in.Parser.Parse(path)
			if err != nil {
				return err
			}

			encoder.Encode(IndexAction{
				Index: IndexDocument{
					Index: re.Index,
				},
			})

			encoder.Encode(parser.Document{
				Path:  path,
				Email: em,
			})
			counter++
			if counter == 100 {
				if err := Upload(re, buf); err != nil {
					return err
				}
				counter = 0
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("error while indexing the files!", err)
		return
	}
	if counter > 0 {
		if err := Upload(re, buf); err != nil {
			log.Printf("error uploading the files: %v", err)
		}
	}
	log.Println("Indexing completed successfully.")
}
