package rshttp

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/realsangil/apimonitor/pkg/rsjson"
	"github.com/realsangil/apimonitor/pkg/rslog"
	"github.com/realsangil/apimonitor/pkg/rsvalid"
)

type Header rsjson.MapJson

func (header Header) GetHttpHeader() http.Header {
	if rsvalid.IsZero(header) {
		return nil
	}
	var httpHeader http.Header
	for k, v := range header {
		httpHeader.Add(k, convertInterfaceToString(v))
	}
	return httpHeader
}

type Query rsjson.MapJson

func (query Query) GetHttpQuery() url.Values {
	if rsvalid.IsZero(query) {
		return nil
	}
	var httpQuery url.Values
	for k, v := range query {
		httpQuery.Add(k, convertInterfaceToString(v))
	}
	return httpQuery
}

func (query Query) GetHttpQueryString() string {
	q := query.GetHttpQuery()
	if q == nil {
		return ""
	}
	return q.Encode()
}

type Request struct {
	RawUrl string
	Header Header
	Query  Query
	Body   map[string]interface{}
}

func (request Request) GetUrl() string {
	parsedUrl, err := url.Parse(request.RawUrl)
	if err != nil {
		rslog.Error(err)
		return ""
	}
	query := request.Query.GetHttpQuery()
	for k, v := range query {
		parsedUrl.Query().Add(k, v[0])
	}
	parsedUrl.RawQuery = parsedUrl.Query().Encode()
	rslog.Infof("request_url='%s'", parsedUrl.String())
	return parsedUrl.String()
}

type Response interface {
	GetStatusCode() int
	GetResponseTime() int64
	GetBody() interface{}
}

type HttpResponse struct {
	StatusCode   int
	ResponseTime int64
	Body         map[string]interface{}
}

func (d HttpResponse) GetStatusCode() int {
	return d.StatusCode
}

func (d HttpResponse) GetResponseTime() int64 {
	return d.ResponseTime
}

func (d HttpResponse) GetBody() interface{} {
	return d.Body
}

func convertInterfaceToString(i interface{}) string {
	return fmt.Sprintf("%v", i)
}
