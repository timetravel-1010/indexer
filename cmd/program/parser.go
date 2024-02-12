package program

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var (
	fields = []string{
		"Message-ID: ",
		"Date: ",
		"From: ",
		"To: ",
		"Cc: ",
		"Bcc: ",
		"Subject: ",
		"Mime-Version: ",
		"Content-Type: ",
		"Content-Transfer-Encoding: ",
		"X-From: ",
		"X-To: ",
		"X-cc: ",
		"X-bcc: ",
		"X-Folder: ",
		"X-Origin: ",
		"X-FileName: ",
		"Body",
	}
)

// A Document contains the path of the email and the email itself.
type Document struct {
	Path  string `json:"path"` // path to the email.
	Email *Email `json:"email"`
}

// A Parser
type Parser struct{}

// Parse parses the txt email file into the Email structure.
// If there is an error, it will be of type *PathError.
func (p *Parser) Parse(filePath string) (*Email, error) {
	eb := NewEmailBuilder()

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
	return eb.build(), nil
}

// CheckEmpty
func CheckEmpty(filePath string) (bool, error) {
	fi, err := os.Stat(filePath)
	if err != nil {
		return true, err
	}

	return (fi.Size() == 0), nil
}

// MapStrings
func MapStrings(arr []string, f func(string) string) []string {
	newArr := make([]string, len(arr))
	for i, s := range arr {
		newArr[i] = f(s)
	}
	return newArr
}
