package program

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/timetravel-1010/indexer/cmd/util"
)

var (
	fields = []string{
		"Message-ID",
		"Date",
		"From",
		"To",
		"Cc",
		"Bcc",
		"Subject",
		"Mime-Version",
		"Content-Type",
		"Content-Transfer-Encoding",
		"X-From",
		"X-To",
		"X-cc",
		"X-bcc",
		"X-Folder",
		"X-Origin",
		"X-FileName",
		"Body",
	}
)

const (
	regexEmailAddress = `[\w._%+-]+@[\w.-]+\.[A-Za-z]{2,}`
	regexName         = `^[a-zA-ZÀ-ÿ0-9 ()-]*$`
)

// An Email contains all the information of an e-mail.
type Email struct {
	MessageID               string   `json:"messageId"`
	Date                    string   `json:"date"`
	From                    string   `json:"from"`
	To                      []string `json:"to"`
	CC                      []string `json:"cc"`
	BCC                     []string `json:"bcc"`
	Subject                 string   `json:"subject"`
	MimeVersion             string   `json:"mimeVersion"`
	ContentType             string   `json:"contentType"`
	ContentTransferEncoding string   `json:"contentTransferEncoding"`
	XFrom                   string   `json:"xFrom"`
	XTo                     []string `json:"xTo"`
	Xcc                     []string `json:"xcc"`
	Xbcc                    []string `json:"xbcc"`
	XFolder                 string   `json:"xFolder"`
	XOrigin                 string   `json:"xOrigin"`
	XFileName               string   `json:"xFileName"`
	Body                    string   `json:"body"`
}

type EmailBuilder struct {
	MessageID               strings.Builder
	Date                    strings.Builder
	From                    strings.Builder
	To                      []string
	CC                      []string
	BCC                     []string
	Subject                 strings.Builder
	MimeVersion             strings.Builder
	ContentType             strings.Builder
	ContentTransferEncoding strings.Builder
	XFrom                   strings.Builder
	XTo                     []string
	Xcc                     []string
	Xbcc                    []string
	XFolder                 strings.Builder
	XOrigin                 strings.Builder
	XFileName               strings.Builder
	Body                    strings.Builder
}

func (eb *EmailBuilder) build() *Email {
	return &Email{
		MessageID:               eb.MessageID.String(),
		Date:                    eb.Date.String(),
		From:                    eb.From.String(),
		To:                      eb.To,
		CC:                      eb.CC,
		BCC:                     eb.BCC,
		Subject:                 eb.Subject.String(),
		MimeVersion:             eb.MimeVersion.String(),
		ContentType:             eb.ContentType.String(),
		ContentTransferEncoding: eb.ContentTransferEncoding.String(),
		XFrom:                   eb.XFrom.String(),
		XTo:                     eb.XTo,
		Xcc:                     eb.Xcc,
		Xbcc:                    eb.Xbcc,
		XFolder:                 eb.XFolder.String(),
		XOrigin:                 eb.XOrigin.String(),
		XFileName:               eb.XFileName.String(),
		Body:                    eb.Body.String(),
	}
}

// A Document contains the path of the email and the email itself.
type Document struct {
	Path  string `json:"path"` // path to the email.
	Email *Email `json:"email"`
}

// A Parser
type Parser struct {
}

type NotEmailError struct{}

func (ner *NotEmailError) Error() string {
	return ""
}

var ErrorMail = errors.New("is not email")

// Parse parses the txt email file into the Email structure.
// If there is an error, it will be of type *PathError.
func (p *Parser) Parse(filePath string) (*Email, error) {
	//em := Email{}
	var eb EmailBuilder = EmailBuilder{}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, ErrorMail
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentField := ""
	inBody := false

	// Check if the file is an email
	scanner.Scan()
	line := scanner.Text()
	saveLine(&eb, &line, &currentField, filePath, &inBody)
	isEmail := eb.MessageID.Len() != 0
	if !isEmail {
		return nil, errors.New(fmt.Sprintf("the file %s is not an email", filePath))
	}

	for scanner.Scan() {
		line := scanner.Text()
		saveLine(&eb, &line, &currentField, filePath, &inBody)
	}
	em := eb.build()
	return em, nil
}

// CheckEmpty
func CheckEmpty(filePath string) (bool, error) {
	fi, err := os.Stat(filePath)
	if err != nil {
		return true, err
	}

	return (fi.Size() == 0), nil
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

// saveLine
func saveLine(em *EmailBuilder, line *string, currentField *string, filePath string, inBody *bool) {

	if *inBody {
		em.Body.WriteString("\n")
		em.Body.WriteString(*line)
		return
	}
	subStrings := strings.SplitN(*line, ":", 2)

	// Check if continues in a section.
	// Probably this will not work always.
	if idx := util.IndexOf(subStrings[0], fields); idx == -1 && em.Body.Len() == 0 {
		addLine(line, *currentField, em, filePath)
	} else if len(subStrings) == 2 {
		*currentField = subStrings[0]
		line := strings.TrimSpace(subStrings[1])
		setValue(*currentField, &line, em, filePath)
	}
	*inBody = *currentField == "X-FileName"
}

// setValue
func setValue(currentField string, l *string, em *EmailBuilder, filePath string) {
	switch currentField {
	case "Message-ID":
		em.MessageID.WriteString(*l)
	case "Date":
		em.Date.WriteString(*l)
	case "From":
		em.From.WriteString(*l)
	case "To":
		em.To = parseAddresses(*l)
		em.To = append(em.To, parseNames(*l)...)
	case "Cc":
		em.CC = parseAddresses(*l)
		em.CC = append(em.CC, parseNames(*l)...)
	case "Bcc":
		em.BCC = parseAddresses(*l)
		em.BCC = append(em.BCC, parseNames(*l)...)
	case "Subject":
		em.Subject.WriteString(*l)
	case "Mime-Version":
		em.MimeVersion.WriteString(*l)
	case "Content-Type":
		em.ContentType.WriteString(*l)
	case "Content-Transfer-Encoding":
		em.ContentTransferEncoding.WriteString(*l)
	case "X-From":
		em.XFrom.WriteString(*l)
	case "X-To":
		em.XTo = MapStrings(strings.Split(*l, ","), strings.TrimSpace)
	case "X-cc":
		em.Xcc = parseAddresses(*l)
		em.Xcc = append(em.Xcc, parseNames(*l)...)
	case "X-bcc":
		em.Xbcc = parseAddresses(*l)
		em.Xbcc = append(em.Xbcc, parseNames(*l)...)
	case "X-Folder":
		em.XFolder.WriteString(*l)
	case "X-Origin":
		em.XOrigin.WriteString(*l)
	case "X-FileName":
		em.XFileName.WriteString(*l)
	default:
		fmt.Println(fmt.Sprintf(`
        ===================ERROR NO MATCH FOUND
        function: setValue
        l: %s
        currentLine: %s
        file: %s
        ===================END ERROR`, *l, currentField, filePath))
	}
}

// addLine
func addLine(l *string, currentField string, em *EmailBuilder, filePath string) {

	switch currentField {
	case "Message-ID":
		em.MessageID.WriteString("\n")
		em.MessageID.WriteString(*l)
	case "Date":
		em.Date.WriteString("\n")
		em.Date.WriteString(*l)
	case "From":
		em.From.WriteString("\n")
		em.From.WriteString(*l)
	case "To":
		em.To = append(em.To, parseAddresses(*l)...)
		em.To = append(em.To, parseNames(*l)...)
	case "Cc":
		em.CC = append(em.CC, parseAddresses(*l)...)
		em.CC = append(em.CC, parseNames(*l)...)
	case "Bcc":
		em.BCC = append(em.BCC, parseAddresses(*l)...)
		em.BCC = append(em.BCC, parseNames(*l)...)
	case "Subject":
		em.Subject.WriteString("\n")
		em.Subject.WriteString(*l)
	case "Mime-Version":
		em.MimeVersion.WriteString("\n")
		em.MimeVersion.WriteString(*l)
	case "Content-Type":
		em.ContentType.WriteString("\n")
		em.ContentType.WriteString(*l)
	case "Content-Transfer-Encoding":
		em.ContentTransferEncoding.WriteString("\n")
		em.ContentTransferEncoding.WriteString(*l)
	case "X-From":
		em.XFrom.WriteString("\n")
		em.XFrom.WriteString(*l)
	case "X-To":
		em.XTo = append(em.XTo, MapStrings(strings.Split(*l, ","), strings.TrimSpace)...)
	case "X-cc":
		em.Xcc = append(em.Xcc, parseAddresses(*l)...)
		em.Xcc = append(em.Xcc, parseNames(*l)...)
	case "X-bcc":
		em.Xbcc = append(em.Xbcc, parseAddresses(*l)...)
		em.Xbcc = append(em.Xbcc, parseNames(*l)...)
	case "X-Folder":
		em.XFolder.WriteString("\n")
		em.XFolder.WriteString(*l)
	case "X-Origin":
		em.XOrigin.WriteString("\n")
		em.XOrigin.WriteString(*l)
	case "X-FileName":
		em.XFileName.WriteString("\n")
		em.XFileName.WriteString(*l)
	default:
		fmt.Println("addLine: No match found and currentLine =", currentField, "file:", filePath)
	}
}
