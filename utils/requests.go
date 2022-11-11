package utils

import (
	"github.com/parnurzeal/gorequest"
	"go.uber.org/multierr"
	"net/http"
	"net/url"
	"time"
)

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	PATCH  Method = "PATCH"
	DELETE Method = "DELETE"
)

func Request(url url.URL, req interface{}, methods ...Method) (res *http.Response, body string, err error) {
	request := gorequest.New()
	u := url.String()

	method := GET
	if len(methods) > 0 {
		method = methods[0]
	}

	switch method {
	case GET:
		request.Get(u)
	case POST:
		request.Post(u)
	case PUT:
		request.Put(u)
	case PATCH:
		request.Patch(u)

	case DELETE:
		request.Delete(u)

	default:
		request.Get(u)

	}

	res, body, errs := request.Send(req).End()

	err = multierr.Combine(errs...)

	return
}

func RequestWithTimeout(url url.URL, req interface{}, timeout time.Duration, methods ...Method) (res *http.Response, body string, err error) {
	request := gorequest.New().Timeout(timeout)
	u := url.String()

	method := GET
	if len(methods) > 0 {
		method = methods[0]
	}

	switch method {
	case GET:
		request.Get(u)
	case POST:
		request.Post(u)
	case PUT:
		request.Put(u)
	case PATCH:
		request.Patch(u)

	case DELETE:
		request.Delete(u)

	default:
		request.Get(u)

	}

	res, body, errs := request.Send(req).End()

	err = multierr.Combine(errs...)

	return
}
