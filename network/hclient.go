package network

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetParamFromPost(c echo.Context, paramName string) (string, error) {
	param := c.FormValue(paramName)
	if strings.TrimSpace(param) == "" {
		return "", errors.New(paramName + " cannot be empty")
	}

	return param, nil
}

func SendPostRequest(url string, params map[string]string) (string, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return "", err
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	str, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	return string(str), err
}
