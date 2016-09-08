package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	http_requestbuilder "github.com/bborbe/http/requestbuilder"
	"github.com/golang/glog"
)

type executeRequest func(req *http.Request) (resp *http.Response, err error)

type rest struct {
	httpRequestBuilderProvider http_requestbuilder.HTTPRequestBuilderProvider
	executeRequest             executeRequest
}

type Rest interface {
	Call(url string, method string, request interface{}, response interface{}, header http.Header) error
}

func New(
	executeRequest executeRequest,
	httpRequestBuilderProvider http_requestbuilder.HTTPRequestBuilderProvider,
) *rest {
	r := new(rest)
	r.httpRequestBuilderProvider = httpRequestBuilderProvider
	r.executeRequest = executeRequest
	return r
}

func (r *rest) Call(url string, method string, request interface{}, response interface{}, headers http.Header) error {
	glog.V(2).Infof("call %s on path %s", method, url)
	start := time.Now()
	defer glog.V(2).Infof("create completed in %dms", time.Now().Sub(start)/time.Millisecond)
	glog.V(2).Infof("send message to %s", url)
	requestbuilder := r.httpRequestBuilderProvider.NewHTTPRequestBuilder(url)
	for key, values := range headers {
		requestbuilder.AddHeader(key, values...)
	}
	requestbuilder.SetMethod(method)
	requestbuilder.AddContentType("application/json")
	if request != nil {
		content, err := json.Marshal(request)
		if err != nil {
			glog.V(2).Infof("marhal request failed: %v", err)
			return err
		}
		glog.V(2).Infof("send request to %s: %s", url, string(content))
		requestbuilder.SetBody(bytes.NewBuffer(content))
	}
	req, err := requestbuilder.Build()
	if err != nil {
		glog.V(2).Infof("build request failed: %v", err)
		return err
	}
	resp, err := r.executeRequest(req)
	if err != nil {
		glog.V(2).Infof("execute request failed: %v", err)
		return err
	}
	if resp.StatusCode/100 != 2 {
		glog.V(2).Infof("status %d not 2xx", resp.StatusCode)
		return fmt.Errorf("request to %s failed with status: %d", url, resp.StatusCode)
	}
	if response != nil {
		if err = json.NewDecoder(resp.Body).Decode(response); err != nil {
			glog.V(2).Infof("decode response failed: %v", err)
			return err
		}
	}
	glog.V(2).Infof("rest call successful")
	return nil
}
