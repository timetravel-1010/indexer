package program

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/timetravel-1010/indexer/cli/internal/parser"
)

type Credentials struct {
	User     string
	Password string
}

type HttpRequest struct {
	Creds   Credentials
	BaseURL string
	Port    string
	Index   string
	Type    string
}

type IndexDocument struct {
	Index string `json:"_index"`
}

type IndexAction struct {
	Index IndexDocument `json:"index"`
}

type Payload struct {
	Index        string            `json:"index"`
	DocumentData []parser.Document `json:"records"`
}

var client = &http.Client{}

// Upload
func Upload(re HttpRequest, payload *bytes.Buffer) error {
	u := fmt.Sprintf("http://%s:%s/api/_bulk", re.BaseURL, re.Port)
	req, err := http.NewRequest("POST", u, payload)
	if err != nil {
		return err
	}

	req.SetBasicAuth(re.Creds.User, re.Creds.Password)
	req.Header.Set("Content-Type", "application/x-ndjson")

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := getBodyResponse(res)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New(fmt.Sprintf("status code: %d - %s\n", res.StatusCode, body))
	}

	payload.Reset()

	return nil
}

func getBodyResponse(res *http.Response) (string, error) {
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
