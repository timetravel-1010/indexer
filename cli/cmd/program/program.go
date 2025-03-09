package program

import (
	"bytes"
	"encoding/json"
	"io/fs"
	"log"
	"path/filepath"
	"sync"
	"time"

	"github.com/timetravel-1010/indexer/cli/cmd/util"
	"github.com/timetravel-1010/indexer/cli/internal/parser"
)

// Indexer
type Indexer struct {
	Parser parser.ParserI
}

// Index indexes files in the given directory
func (in *Indexer) Index(dir string, req HttpRequest) error {
	start := time.Now()
	var totalIndexed int

	fileChan := make(chan string, 100) // Buffered channel to avoid blocking
	var wg sync.WaitGroup

	const numWorkers = 4
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			in.processFiles(fileChan, req, &totalIndexed)
		}()
	}

	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Printf("error accessing file %s: %v", path, err)
			return nil
		}

		if !info.IsDir() {
			fileChan <- path
		}
		return nil
	})

	close(fileChan)
	wg.Wait()

	if err != nil {
		return err
	}

	log.Printf(
		"process finished: %d files indexed in %.2f seconds\n",
		totalIndexed,
		time.Since(start).Seconds(),
	)
	return nil
}

// processFiles processes files from the channel and uploads data in batches
func (in *Indexer) processFiles(
	fileChan <-chan string,
	req HttpRequest,
	totalIndexed *int,
) {
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	counter := 0

	for path := range fileChan {
		if isEmpty, err := util.CheckEmpty(path); err != nil {
			log.Printf("error checking file %s: %v", path, err)
			continue
		} else if isEmpty {
			log.Printf("skipping empty file: %s", path)
			continue
		}

		em, err := in.Parser.Parse(path)
		if err != nil {
			log.Printf("error parsing file %s: %v", path, err)
			continue
		}

		enErr := encoder.Encode(IndexAction{Index: IndexDocument{Index: req.Index}})
		if enErr != nil {
			log.Printf("error encondig index action")
		}

		enErr = encoder.Encode(parser.Document{Path: path, Email: em})
		if enErr != nil {
			log.Printf("error encondig documents")
		}

		counter++

		if counter >= 100 {
			if err := Upload(req, buf); err != nil {
				log.Printf("error uploading files: %v", err)
			} else {
				*totalIndexed += counter
			}
			buf.Reset() // Reset buffer after upload
			counter = 0
		}
	}

	// Upload any remaining data
	if counter > 0 {
		if err := Upload(req, buf); err != nil {
			log.Printf("error uploading remaining files: %v", err)
		} else {
			*totalIndexed += counter
		}
	}
}
