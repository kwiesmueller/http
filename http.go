package http

import (
	"io/ioutil"
	"net/http"
)

func ResponseToString(response *http.Response) (string, error) {
	content, err := ResponseToByteArray(response)
	return string(content), err
}

func ResponseToByteArray(response *http.Response) ([]byte, error) {
	body := response.Body
	return ioutil.ReadAll(body)
}
