package bearer

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

func CreateBearerToken(name string, value string) string {
	logger.Debugf("create bearer from: %s:%s", name, value)
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", name, value)))
}

func CreateBearerHeader(name string, value string) string {
	return fmt.Sprintf("Bearer %s", CreateBearerToken(name, value))
}

func ParseBearerHttpRequest(req *http.Request) (string, string, error) {
	authorizations := req.Header["Authorization"]
	if len(authorizations) != 1 {
		return "", "", fmt.Errorf("header Authorization invalid")
	}
	return ParseBearerHeader(authorizations[0])
}

func ParseBearerHeader(header string) (string, string, error) {
	logger.Debugf("parse bearer: %s", header)
	if strings.Index(header, "Bearer ") != 0 {
		return "", "", fmt.Errorf("header Authorization invalid")
	}
	tokens := strings.SplitN(header, " ", 2)
	if len(tokens) != 2 {
		return "", "", fmt.Errorf("parse token from header failed")
	}
	return ParseBearerToken(tokens[1])
}

func ParseBearerToken(token string) (string, string, error) {
	logger.Debugf("parse token: %s", token)
	value, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", "", err
	}
	parse := strings.SplitN(string(value), ":", 2)
	if len(parse) != 2 {
		return "", "", fmt.Errorf("parse header failed")
	}
	return parse[0], parse[1], nil
}
