package restclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Execute(request Request) error {

	var payload bytes.Buffer
	if request.payload != nil {
		payload = *bytes.NewBuffer(request.payload)
	}

	req, err := http.NewRequest(string(request.method), request.url, &payload)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	for key, value := range request.headers {
		req.Header.Set(key, value)
	}

	queryParams := req.URL.Query()
	for key, value := range request.queryParams {
		queryParams.Add(key, value)
	}

	req.URL.RawQuery = queryParams.Encode()

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		decoder := json.NewDecoder(resp.Body)
		decoder.DisallowUnknownFields()
		err = decoder.Decode(request.response)
		if err != nil {
			log.Fatal("Error decoding response. ", err)
		}
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("Error decoding response. ", err)
			return err
		}

		return fmt.Errorf("An error ocurred: %v", string(body))
	}

	return nil

}
