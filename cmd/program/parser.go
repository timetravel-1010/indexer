package program

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	regexEmailAddress = `[\w._%+-]+@[\w.-]+\.[A-Za-z]{2,}`
	regexName         = `^[a-zA-ZÀ-ÿ0-9 ()-]*$`
)

// An Email contains all the information of an e-mail.
type Email struct {
	MessageID               string   `json:"Message-ID"`
	Date                    string   `json:"Date"`
	From                    string   `json:"From"`
	To                      []string `json:"To"`
	CC                      []string `json:"CC"`
	BCC                     []string `json:"BCC"`
	Subject                 string   `json:"Subject"`
	MimeVersion             string   `json:"Mime-Version"`
	ContentType             string   `json:"Content-Type"`
	ContentTransferEncoding string   `json:"Content-Transfer-Encoding"`
	XFrom                   string   `json:"X-From"`
	XTo                     []string `json:"X-To"`
	Xcc                     []string `json:"X-cc"`
	Xbcc                    []string `json:"X-bcc"`
	XFolder                 string   `json:"X-Folder"`
	XOrigin                 string   `json:"X-Origin"`
	XFileName               string   `json:"X-FileName"`
	Body                    string   `json:"Body"`
}

// A Document contains the path of the email and the email itself.
type Document struct {
	Path  string `json:"path"` // path to the email.
	Email *Email `json:"email"`
}

// A Parser
type Parser struct {
}

// Parse parses the txt email file into the Email structure.
// If there is an error, it will be of type *PathError.
func (p *Parser) Parse(filePath string) (*Email, error) {
	em := Email{}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentField := ""
	isEmpty := true

	for scanner.Scan() {
		isEmpty = false
		line := scanner.Text()
		subStrings := strings.SplitN(line, ":", 2)
		if len(subStrings) == 2 && currentField != "X-FileName" {
			currentField = subStrings[0]
			l := strings.TrimSpace(subStrings[1])
			switch currentField {
			case "Message-ID":
				em.MessageID = l
			case "Date":
				em.Date = l
			case "From":
				em.From = l
			case "To":
				em.To = append(em.To, parseNames(l)...)
			case "Cc":
				em.CC = parseAddresses(l)
				em.CC = append(em.CC, parseNames(l)...)
			case "Bcc":
				em.BCC = parseAddresses(l)
				em.BCC = append(em.BCC, parseNames(l)...)
			case "Subject":
				em.Subject = l
			case "Mime-Version":
				em.MimeVersion = l
			case "Content-Type":
				em.ContentType = l
			case "Content-Transfer-Encoding":
				em.ContentTransferEncoding = l
			case "X-From":
				em.XFrom = l
			case "X-To":
				em.XTo = MapStrings(strings.Split(l, ","), strings.TrimSpace)
			case "X-cc":
				em.Xcc = parseAddresses(l)
				em.Xcc = append(em.Xcc, parseNames(l)...)
			case "X-bcc":
				em.Xbcc = parseAddresses(l)
				em.Xbcc = append(em.Xbcc, parseNames(l)...)
			case "X-Folder":
				em.XFolder = l
			case "X-Origin":
				em.XOrigin = l
			case "X-FileName":
				em.XFileName = l
			default:
				fmt.Println("No match found and currentLine=", currentField)
			}
		} else if currentField == "X-FileName" { // Body content
			em.Body += "\n"
			if subStrings != nil {
				em.Body += subStrings[0]
			}
		} else if currentField == "To" {
			em.To = append(em.To, parseAddresses(subStrings[0])...)
		}
		if em.MessageID == "" {
			log.Printf("The file %s is not an email, skipped.\n", filePath)
			break
		}
	}
	if isEmpty {
		fmt.Printf("The file %s is empty.", filePath)
		return nil, nil
	}
	return &em, nil
}

// parseAddresses
func parseAddresses(s string) []string {
	return GetStringsByRegexp(s, regexEmailAddress)
}

// parseNames
func parseNames(s string) []string {
	return GetStringsByRegexp(s, regexName)
}

// GetStringsByRegexp
func GetStringsByRegexp(s string, regex string) []string {
	return regexp.MustCompile(regex).FindAllString(s, -1)
}

// MapStrings
func MapStrings(arr []string, f func(string) string) []string {
	newArr := make([]string, len(arr))
	for i, s := range arr {
		newArr[i] = f(s)
	}
	return newArr
}
