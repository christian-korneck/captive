package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type htdoc struct {
	Head struct {
		Title string `xml:"TITLE"`
	} `xml:"HEAD"`
	Body string `xml:"BODY"`
}

func offl() {
	fmt.Println("offline")
	os.Exit(1)
}

func onl() {
	fmt.Println("online")
	os.Exit(0)
}

//Iscaptive checks the captive status
func Iscaptive(client HTTPClient) (bool, error) {
	const SuccessTitle = "Success"
	const SuccessBody = "Success"
	const url = "http://captive.apple.com"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return false, err
	}

	res, err := client.Do(req)
	if err != nil {
		return false, err
	}

	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return false, err
	}

	h := htdoc{}
	err = xml.Unmarshal(b, &h)

	if err != nil {
		return false, err
	}

	if h.Head.Title == SuccessTitle && h.Body == SuccessBody {
		return true, nil
	} else {
		return false, err
	}
}

func main() {

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	c, err := Iscaptive(client)

	if err != nil || c != true {
		offl()
	} else {
		onl()
	}

}
