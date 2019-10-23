package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func SendPostRequest(url string, params map[string]string) (string, error) {
	paramBytes, err := json.Marshal(params)
	if err != nil {
		return "", err
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(paramBytes))
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	return string(body), err
}
