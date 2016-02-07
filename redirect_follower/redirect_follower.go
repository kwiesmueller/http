package redirect_follower

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/bborbe/log"
)

const LIMIT = 10

var logger = log.DefaultLogger

type ExecuteRequest func(req *http.Request) (resp *http.Response, err error)

type RedirectFollower interface {
	ExecuteRequestAndFollow(req *http.Request) (resp *http.Response, err error)
}

type redirectFollower struct {
	executeRequest ExecuteRequest
}

func New(executeRequest ExecuteRequest) *redirectFollower {
	r := new(redirectFollower)
	r.executeRequest = executeRequest
	return r
}

func (r *redirectFollower) ExecuteRequestAndFollow(req *http.Request) (*http.Response, error) {
	return executeRequestAndFollow(r.executeRequest, req, 0)
}

func executeRequestAndFollow(executeRequest ExecuteRequest, req *http.Request, counter int) (*http.Response, error) {
	logger.Debugf("execute request to %s", req.URL)
	logger.Debugf("request %v\n", req)
	resp, err := executeRequest(req)
	if err != nil {
		return nil, err
	}
	logger.Debugf("response %v", resp)
	if resp.StatusCode/100 == 3 {
		logger.Debugf("redirect - statuscode: %d", resp.StatusCode)
		if counter > LIMIT {
			return nil, fmt.Errorf("redirect limit reached")
		}
		var reqCopy http.Request = *req
		var p *http.Request = &reqCopy
		var location []string = resp.Header["Location"]
		if len(location) != 1 {
			return nil, fmt.Errorf("redirect failed")
		}
		logger.Debugf("redirect to %s", location[0])
		p.URL, err = url.Parse(location[0])
		if err != nil {
			return nil, nil
		}
		p.Host = p.URL.Host
		return executeRequestAndFollow(executeRequest, p, counter+1)
	}

	return resp, nil
}
