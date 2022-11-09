package utiliity

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func ExternalPostAPI(url string, reqBody *bytes.Reader) string {

	resp, err := http.Post(url, "application/json", reqBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	return string(body)
}
