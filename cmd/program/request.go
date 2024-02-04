package program

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
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
	Index        string     `json:"index"`
	DocumentData []Document `json:"records"`
}

var client = &http.Client{}

// Upload
func Upload(re HttpRequest, payload *bytes.Buffer) {
	//u := fmt.Sprintf("http://%s:%s/api/_bulkv2/", re.BaseURL, re.Port) //, re.Index, re.Type)

	u := fmt.Sprintf("http://%s:%s/api/_bulk", re.BaseURL, re.Port) //, re.Index, re.Type)
	req, err := http.NewRequest("POST", u, payload)

	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(re.Creds.User, re.Creds.Password)
	//req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Type", "application/x-ndjson")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
	payload.Reset()
}
