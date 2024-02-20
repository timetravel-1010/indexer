package email

import (
	"fmt"
	"log"
	"net/mail"
	"strings"

	"github.com/timetravel-1010/indexer/internal/regex"
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

// An Email contains all the information of an e-mail.
type Email struct {
	MessageID               string          `json:"Message-Id"`
	Date                    string          `json:"Date"`
	From                    string          `json:"From"`
	To                      []*mail.Address `json:"To"`
	CC                      []*mail.Address `json:"Cc"`
	BCC                     []*mail.Address `json:"Bcc"`
	Subject                 string          `json:"Subject"`
	MimeVersion             string          `json:"Mime-Version"`
	ContentType             string          `json:"Content-Type"`
	ContentTransferEncoding string          `json:"Content-Transfer-Encoding"`
	XFrom                   string          `json:"X-From"`
	XTo                     []*mail.Address `json:"X-To"`
	Xcc                     []*mail.Address `json:"X-Cc"`
	Xbcc                    []*mail.Address `json:"X-Bcc"`
	XFolder                 string          `json:"X-Folder"`
	XOrigin                 string          `json:"X-Origin"`
	XFileName               string          `json:"X-Filename"`
	Body                    string          `json:"Body"`
}

type EmailI interface {
	DoSomething()
}

func (em Email) DoSomething() {}

// An EmailBuilder
type EmailBuilder struct {
	MessageID               strings.Builder
	Date                    strings.Builder
	From                    strings.Builder
	To                      []*mail.Address
	CC                      []*mail.Address
	BCC                     []*mail.Address
	Subject                 strings.Builder
	MimeVersion             strings.Builder
	ContentType             strings.Builder
	ContentTransferEncoding strings.Builder
	XFrom                   strings.Builder
	XTo                     []*mail.Address
	Xcc                     []*mail.Address
	Xbcc                    []*mail.Address
	XFolder                 strings.Builder
	XOrigin                 strings.Builder
	XFileName               strings.Builder
	Body                    strings.Builder
	inBody                  bool
	currentField            string
}

var setterMap map[string]func(*string) error

// NewEmailBuilder returns a pointer to a new EmailBuilder struct with zero values.
func NewEmailBuilder() *EmailBuilder {
	eb := &EmailBuilder{}
	setterMap = setterMapBuilder(eb)
	return eb
}

// Build
func (eb *EmailBuilder) Build() *Email {
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
			// TODO: pending check if after (lineContent) is "", " ", etc.
			err := eb.setValue(&after, filePath)
			if err != nil {
				log.Println("error in SaveLine: ", err)
			}
			//setValue(*currentField, &after, eb, filePath)
			eb.inBody = eb.currentField == "X-FileName: "
			return
		}
	}
	// Continues in a field
	err := eb.setValue(line, filePath)
	if err != nil {
		log.Println("error in SaveLine:", err)
	}
}

func setterMapBuilder(eb *EmailBuilder) map[string]func(*string) error {
	return map[string]func(*string) error{
		"Message-ID: ": func(lineContent *string) error {
			_, err := eb.MessageID.WriteString(*lineContent)
			return err
		},
		"Date: ": func(lineContent *string) error {
			_, err := eb.Date.WriteString(*lineContent)
			return err
		},
		"From: ": func(lineContent *string) error {
			_, err := eb.From.WriteString(*lineContent)
			return err
		},
		"To: ": func(lineContent *string) error {
			setAddresses(&eb.To, lineContent)
			return nil
		},
		"Cc: ": func(lineContent *string) error {
			setAddresses(&eb.CC, lineContent)
			return nil
		},
		"Bcc: ": func(lineContent *string) error {
			setAddresses(&eb.BCC, lineContent)
			return nil
		},
		"Subject: ": func(lineContent *string) error {
			_, err := eb.Subject.WriteString(*lineContent)
			return err
		},
		"Mime-Version: ": func(lineContent *string) error {
			_, err := eb.MimeVersion.WriteString(*lineContent)
			return err
		},
		"Content-Type: ": func(lineContent *string) error {
			_, err := eb.ContentType.WriteString(*lineContent)
			return err
		},
		"Content-Transfer-Encoding: ": func(lineContent *string) error {
			_, err := eb.ContentTransferEncoding.WriteString(*lineContent)
			return err
		},
		"X-From: ": func(lineContent *string) error {
			_, err := eb.XFrom.WriteString(*lineContent)
			return err
		},
		"X-To: ": func(lineContent *string) error {
			setAddresses(&eb.XTo, lineContent)
			return nil
		},
		"X-cc: ": func(lineContent *string) error {
			setAddresses(&eb.Xcc, lineContent)
			return nil
		},
		"X-bcc: ": func(lineContent *string) error {
			setAddresses(&eb.BCC, lineContent)
			return nil
		},
		"X-Folder: ": func(lineContent *string) error {
			_, err := eb.XFolder.WriteString(*lineContent)
			return err
		},
		"X-Origin: ": func(lineContent *string) error {
			_, err := eb.XOrigin.WriteString(*lineContent)
			return err
		},
		"X-FileName: ": func(lineContent *string) error {
			_, err := eb.XFileName.WriteString(*lineContent)
			return err
		},
	}
}

// SetValue sets the value of a specific field in the email header.
func (eb *EmailBuilder) setValue(lineContent *string, filePath string) error {
	if lineContent == nil {
		return fmt.Errorf("line content cannot be nil")
	}

	if setter, ok := setterMap[eb.currentField]; ok {
		err := setter(lineContent)
		if err != nil {
			log.Println("error in setterFunc:", eb.currentField, " error:", err)
		}
	} else {
		return fmt.Errorf("no match found for field '%s' in file '%s'", eb.currentField, filePath)
	}

	return nil
}

// setAddresses
func setAddresses(addrsField *[]*mail.Address, addrsList *string) {

	temp := *addrsField
	pairs := strings.Split(*addrsList, ",")

	for _, pair := range pairs {
		addr := new(mail.Address)
		addr.Name = regex.GetName(pair)
		addr.Address = regex.GetEmailAddress(pair)
		temp = append(temp, addr)
	}
	*addrsField = temp
}
