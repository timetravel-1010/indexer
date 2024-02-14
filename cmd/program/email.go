package program

import (
	"fmt"
	"strings"

	"github.com/timetravel-1010/indexer/cmd/util"
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

// An EmailBuilder
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
	inBody                  bool
	currentField            string
}

var setterMap map[string]func(*string)

// NewEmailBuilder returns a pointer to a new EmailBuilder struct with zero values.
func NewEmailBuilder() *EmailBuilder {
	eb := &EmailBuilder{}
	setterMap = setterMapBuilder(eb)
	return eb
}

// build
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

// saveLine
func (eb *EmailBuilder) SaveLine(line *string, filePath string) {
	if eb.inBody {
		eb.Body.WriteString("\n")
		eb.Body.WriteString(*line)
		return
	}
	for _, field := range fields {
		if after, found := strings.CutPrefix(*line, field); found {
			eb.currentField = field
			eb.setValue(&after, filePath)
			//setValue(*currentField, &after, eb, filePath)
			eb.inBody = eb.currentField == "X-FileName: "
			return
		}
	}
	// Continues in a field
	eb.addLine(line, filePath)
}

func setterMapBuilder(eb *EmailBuilder) map[string]func(*string) {
	return map[string]func(*string){
		"Message-ID: ": func(lineContent *string) { eb.MessageID.WriteString(*lineContent) },
		"Date: ":       func(lineContent *string) { eb.Date.WriteString(*lineContent) },
		"From: ":       func(lineContent *string) { eb.From.WriteString(*lineContent) },
		"To: ": func(lineContent *string) {
			eb.To = parseAddresses(*lineContent)
			eb.To = append(eb.To, parseNames(*lineContent)...)
		},
		"Cc: ": func(lineContent *string) {
			eb.CC = parseAddresses(*lineContent)
			eb.CC = append(eb.CC, parseNames(*lineContent)...)
		},
		"Bcc: ": func(lineContent *string) {
			eb.BCC = parseAddresses(*lineContent)
			eb.BCC = append(eb.BCC, parseNames(*lineContent)...)
		},
		"Subject: ":                   func(lineContent *string) { eb.Subject.WriteString(*lineContent) },
		"Mime-Version: ":              func(lineContent *string) { eb.MimeVersion.WriteString(*lineContent) },
		"Content-Type: ":              func(lineContent *string) { eb.ContentType.WriteString(*lineContent) },
		"Content-Transfer-Encoding: ": func(lineContent *string) { eb.ContentTransferEncoding.WriteString(*lineContent) },
		"X-From: ":                    func(lineContent *string) { eb.XFrom.WriteString(*lineContent) },
		"X-To: ": func(lineContent *string) {
			eb.XTo = util.MapStrings(strings.Split(*lineContent, ","), strings.TrimSpace)
		},
		"X-cc: ": func(lineContent *string) {
			eb.Xcc = parseAddresses(*lineContent)
			eb.Xcc = append(eb.Xcc, parseNames(*lineContent)...)
		},
		"X-bcc: ": func(lineContent *string) {
			eb.Xbcc = parseAddresses(*lineContent)
			eb.Xbcc = append(eb.Xbcc, parseNames(*lineContent)...)
		},
		"X-Folder: ":   func(lineContent *string) { eb.XFolder.WriteString(*lineContent) },
		"X-Origin: ":   func(lineContent *string) { eb.XOrigin.WriteString(*lineContent) },
		"X-FileName: ": func(lineContent *string) { eb.XFileName.WriteString(*lineContent) },
	}

}

// SetValue sets the value of a specific field in the email header.
func (eb *EmailBuilder) setValue(lineContent *string, filePath string) error {
	if lineContent == nil {
		return fmt.Errorf("line content cannot be nil")
	}
	if setter, ok := setterMap[eb.currentField]; ok {
		setter(lineContent)
	} else {
		return fmt.Errorf("no match found for field '%s' in file '%s'", eb.currentField, filePath)
	}

	return nil
}

// setValue
func (eb *EmailBuilder) setValueDeprecated(l *string, currentField, filePath string) {
	switch currentField {
	case "Message-ID: ":
		eb.MessageID.WriteString(*l)
	case "Date: ":
		eb.Date.WriteString(*l)
	case "From: ":
		eb.From.WriteString(*l)
	case "To: ":
		eb.To = parseAddresses(*l)
		eb.To = append(eb.To, parseNames(*l)...)
	case "Cc: ":
		eb.CC = parseAddresses(*l)
		eb.CC = append(eb.CC, parseNames(*l)...)
	case "Bcc: ":
		eb.BCC = parseAddresses(*l)
		eb.BCC = append(eb.BCC, parseNames(*l)...)
	case "Subject: ":
		eb.Subject.WriteString(*l)
	case "Mime-Version: ":
		eb.MimeVersion.WriteString(*l)
	case "Content-Type: ":
		eb.ContentType.WriteString(*l)
	case "Content-Transfer-Encoding: ":
		eb.ContentTransferEncoding.WriteString(*l)
	case "X-From: ":
		eb.XFrom.WriteString(*l)
	case "X-To: ":
		eb.XTo = util.MapStrings(strings.Split(*l, ","), strings.TrimSpace)
	case "X-cc: ":
		eb.Xcc = parseAddresses(*l)
		eb.Xcc = append(eb.Xcc, parseNames(*l)...)
	case "X-bcc: ":
		eb.Xbcc = parseAddresses(*l)
		eb.Xbcc = append(eb.Xbcc, parseNames(*l)...)
	case "X-Folder: ":
		eb.XFolder.WriteString(*l)
	case "X-Origin: ":
		eb.XOrigin.WriteString(*l)
	case "X-FileName: ":
		eb.XFileName.WriteString(*l)
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
func (eb *EmailBuilder) addLine(l *string, filePath string) {
	switch eb.currentField {
	case "Message-ID: ":
		eb.MessageID.WriteString("\n")
		eb.MessageID.WriteString(*l)
	case "Date: ":
		eb.Date.WriteString("\n")
		eb.Date.WriteString(*l)
	case "From: ":
		eb.From.WriteString("\n")
		eb.From.WriteString(*l)
	case "To: ":
		eb.To = append(eb.To, parseAddresses(*l)...)
		eb.To = append(eb.To, parseNames(*l)...)
	case "Cc: ":
		eb.CC = append(eb.CC, parseAddresses(*l)...)
		eb.CC = append(eb.CC, parseNames(*l)...)
	case "Bcc: ":
		eb.BCC = append(eb.BCC, parseAddresses(*l)...)
		eb.BCC = append(eb.BCC, parseNames(*l)...)
	case "Subject: ":
		eb.Subject.WriteString("\n")
		eb.Subject.WriteString(*l)
	case "Mime-Version: ":
		eb.MimeVersion.WriteString("\n")
		eb.MimeVersion.WriteString(*l)
	case "Content-Type: ":
		eb.ContentType.WriteString("\n")
		eb.ContentType.WriteString(*l)
	case "Content-Transfer-Encoding: ":
		eb.ContentTransferEncoding.WriteString("\n")
		eb.ContentTransferEncoding.WriteString(*l)
	case "X-From: ":
		eb.XFrom.WriteString("\n")
		eb.XFrom.WriteString(*l)
	case "X-To: ":
		eb.XTo = append(eb.XTo, util.MapStrings(strings.Split(*l, ","), strings.TrimSpace)...)
	case "X-cc: ":
		eb.Xcc = append(eb.Xcc, parseAddresses(*l)...)
		eb.Xcc = append(eb.Xcc, parseNames(*l)...)
	case "X-bcc: ":
		eb.Xbcc = append(eb.Xbcc, parseAddresses(*l)...)
		eb.Xbcc = append(eb.Xbcc, parseNames(*l)...)
	case "X-Folder: ":
		eb.XFolder.WriteString("\n")
		eb.XFolder.WriteString(*l)
	case "X-Origin: ":
		eb.XOrigin.WriteString("\n")
		eb.XOrigin.WriteString(*l)
	case "X-FileName: ":
		eb.XFileName.WriteString("\n")
		eb.XFileName.WriteString(*l)
	default:
		fmt.Println("addLine: No match found and currentLine =", eb.currentField, "file:", filePath)
	}
}
