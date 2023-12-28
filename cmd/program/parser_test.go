package program

import (
	"strings"
	"testing"
)

func TestParseEmptyEmail(t *testing.T) {
	fileName := "../../tests/empty.txt"
	em, err := Parse(fileName)

	if err != nil {
		t.Fatalf("Error parsing an empty email file! %v", err)
	}
	if em != nil {
		t.Fatalf("got %q, expected nil", em)
	}
}

func TestParaseFullBody(t *testing.T) {
	fileName := "../../tests/email1.txt"
	ex := "Investor, we will refund your money."
	em, err := Parse(fileName)

	if err != nil {
		t.Fatalf("Error parsing the file, %v", err)
	}
	lines := strings.Split(em.Body, "\n")
	lastLine := lines[len(lines)-1]
	if strings.Compare(lastLine, ex) != 0 {
		t.Fatalf("got %q, expected %s", lastLine, ex)
	}
}
