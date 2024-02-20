package parser

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/timetravel-1010/indexer/internal/email"
)

// A Document contains the path of the email and the email itself.
type Document struct {
	Path  string       `json:"path"` // path to the email.
	Email email.EmailI `json:"email"`
}

// A Parser
type ParserI interface {
	Parse(string) (email.EmailI, error)
}

type Parser struct{}

// Parse parses the txt email file into the Email structure.
// If there is an error, it will be of type *PathError.
func (p Parser) Parse(filePath string) (email.EmailI, error) {
	eb := email.NewEmailBuilder()

	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("Error opening the file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Check if the file is an email
	scanner.Scan()
	line := scanner.Text()
	eb.SaveLine(&line, filePath)
	isEmail := eb.MessageID.Len() != 0
	if !isEmail {
		return nil, errors.New(fmt.Sprintf("the file %s is not an email", filePath))
	}

	for scanner.Scan() {
		line := scanner.Text()
		eb.SaveLine(&line, filePath)
	}
	return eb.Build(), nil
}
