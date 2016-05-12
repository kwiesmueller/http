package header

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

func CreateAuthorizationToken(name string, value string) string {
	logger.Debugf("create bearer from: %s:%s", name, value)
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", name, value)))
}

func ParseAuthorizationToken(token string) (string, string, error) {
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

func CreateAuthorizationBearerHeader(name string, value string) string {
	return fmt.Sprintf("Bearer %s", CreateAuthorizationToken(name, value))
}

func ParseAuthorizationBearerHttpRequest(req *http.Request) (string, string, error) {
	return ParseAuthorizationHttpRequest("Bearer", req)
}

func ParseAuthorizationBasisHttpRequest(req *http.Request) (string, string, error) {
	return ParseAuthorizationHttpRequest("Basic", req)
}

func ParseAuthorizationHttpRequest(authtype string, req *http.Request) (string, string, error) {
	authorizations := req.Header["Authorization"]
	if len(authorizations) != 1 {
		return "", "", fmt.Errorf("header Authorization invalid")
	}
	return ParseAuthorizationHeader("Bearer", authorizations[0])
}

func ParseAuthorizationHeader(authtype string, header string) (string, string, error) {
	logger.Debugf("parse %s: %s", authtype, header)
	if strings.Index(header, fmt.Sprintf("%s ", authtype)) != 0 {
		return "", "", fmt.Errorf("header Authorization invalid")
	}
	tokens := strings.SplitN(header, " ", 2)
	if len(tokens) != 2 {
		return "", "", fmt.Errorf("parse token from header failed")
	}
	return ParseAuthorizationToken(tokens[1])
}
