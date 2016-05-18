package util

import (
	"encoding/json"
	"net/http"
	"net/http/httputil"

	"fmt"
	"io/ioutil"
	"strings"

	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

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

func FindFileExtension(response *http.Response) (string, error) {
	var ext string
	if response.Request != nil && response.Request.URL != nil {
		path := response.Request.URL.Path
		pos := strings.LastIndex(path, ".")
		if pos > 0 {
			ext = path[pos+1:]
		}
	}
	if response.Header != nil {
		contentType := response.Header.Get("Content-Type")
		ext = contentTypeToExt[contentType]
	}
	if len(ext) == 0 {
		return "", fmt.Errorf("find extension failed")
	}
	return ext, nil
}

// PrintDump prints dump of request, optionally writing it in the response
func PrintDump(w http.ResponseWriter, r *http.Request, write bool) {
	dump, _ := httputil.DumpRequest(r, true)
	logger.Debugf("%v", string(dump))
	if write == true {
		w.Write(dump)
	}
}

// Decode into a ma[string]interface{} the JSON in the POST Request
func DecodePostJSON(r *http.Request, logging bool) (map[string]interface{}, error) {
	var err error
	var payLoad map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&payLoad)
	if logging == true {
		logger.Debugf("Parsed body:%v", payLoad)
	}
	return payLoad, err
}
