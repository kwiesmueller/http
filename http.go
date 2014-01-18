package http

import "net/http"

func ResponseToString(response *http.Response) string {
	return string(ResponseToByteArray(response))
}

func ResponseToByteArray(response *http.Response) []byte {
	body := response.Body
	result := make([]byte, 0)
	b := make([]byte, 8)
	for {
		n, err := body.Read(b)
		if err != nil || n == 0 {
			return result
		}
		result = append(result, b[:n]...)
	}
	return result
}
