package luci

import (
	"bytes"
        _ "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type RequestData struct {
	Id     int
	Method string
	Params []string
}

func Request(url string, data RequestData) ([]byte, error) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	// TODO: replace hostname and token in url string
	// TODO: append RequestData to POST request
	resp, err := c.Post(url, "application/json", bytes.NewBuffer(nil))
	if err != nil {
		fmt.Printf("Error %s", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error %s", err)
		return nil, err
	}

	fmt.Printf("Body : %s", body)
	return body, err
}
