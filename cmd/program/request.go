package program

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/timetravel-1010/indexer/internal/parser"
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
	//u := fmt.Sprintf("http://%s:%s/api/_bulkv2/", re.BaseURL, re.Port) //, re.Index, re.Type)

	u := fmt.Sprintf("http://%s:%s/api/_bulk", re.BaseURL, re.Port) //, re.Index, re.Type)
	req, err := http.NewRequest("POST", u, payload)
	if err != nil {
		return err
	}

	req.SetBasicAuth(re.Creds.User, re.Creds.Password)
	//req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Type", "application/x-ndjson")
	req.Header.Set("User-Agent", "laptop")

	resp, err := client.Do(req)
	body, err := getBodyResponse(resp)
	if err != nil {
		return err
	}

	log.Println(resp.StatusCode)
	log.Println(body)
	if resp.StatusCode == http.StatusInternalServerError {
		return errors.New("internal server error")
	}

	payload.Reset()

	return nil
}

func getBodyResponse(resp *http.Response) (string, error) {

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
