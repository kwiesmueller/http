package util

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func ResponseToString(response *http.Response) (string, error) {
	content, err := ResponseToByteArray(response)
	return string(content), err
}

func ResponseToByteArray(response *http.Response) ([]byte, error) {
	body := response.Body
	return ioutil.ReadAll(body)
}

var contentTypeToExt = map[string]string{
	"image/jpeg": "jpg",
	"image/png":  "png",
}

func FindFileExtension(response *http.Response) string {
	if response.Request != nil && response.Request.URL != nil{
		path := response.Request.URL.Path
		pos := strings.LastIndex(path, ".")
		if pos > 0 {
			return path[pos+1:]
		}
	}
	if response.Header != nil {
		contentType := response.Header.Get("Content-Type")
		return contentTypeToExt[contentType]

	}
	return ""
}
