package models

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/pkg/errors"

	"github.com/realsangil/apimonitor/pkg/rserrors"
	"github.com/realsangil/apimonitor/pkg/rshttp"
	"github.com/realsangil/apimonitor/pkg/rsjson"
	"github.com/realsangil/apimonitor/pkg/rsvalid"
)

type WebServiceTest struct {
	DefaultValidateChecker
	Id           int64               `json:"id"`
	WebServiceId int64               `json:"-"`
	Path         rshttp.EndpointPath `json:"path"`
	HttpMethod   rshttp.Method       `json:"http_method"`
	ContentType  rshttp.ContentType  `json:"content_type"`
	Desc         string              `json:"desc" gorm:"Type:TEXT"`
	RequestData  rsjson.MapJson      `json:"request_data" gorm:"Type:JSON"`
	Header       rsjson.MapJson      `json:"header" gorm:"Type:JSON"`
	QueryParam   rsjson.MapJson      `json:"query_param" gorm:"Type:JSON"`
	Created      time.Time           `json:"created"`
	LastModified time.Time           `json:"last_modified"`
}

func NewWebServiceTest(webService *WebService, request WebServiceTestRequest) (*WebServiceTest, error) {
	if rsvalid.IsZero(webService, request) {
		return nil, errors.Wrap(rserrors.ErrInvalidParameter, "Endpoint")
	}
	webServiceTest := &WebServiceTest{
		WebServiceId: webService.Id,
		Created:      time.Now(),
	}
	if err := webServiceTest.UpdateFromRequest(request); err != nil {
		return nil, errors.WithStack(err)
	}
	return webServiceTest, nil
}

func (webServiceTest *WebServiceTest) UpdateFromRequest(request WebServiceTestRequest) error {
	webServiceTest.Path = request.Path
	webServiceTest.HttpMethod = request.HttpMethod
	webServiceTest.ContentType = request.ContentType
	webServiceTest.Desc = request.Desc
	webServiceTest.RequestData = request.RequestData
	webServiceTest.Header = request.Header
	webServiceTest.QueryParam = request.Header
	webServiceTest.LastModified = time.Now()
	return webServiceTest.Validate()
}

func (webServiceTest *WebServiceTest) Validate() error {
	if rsvalid.IsZero(
		webServiceTest.WebServiceId,
		webServiceTest.Path,
		webServiceTest.HttpMethod,
		webServiceTest.ContentType,
		webServiceTest.Created,
		webServiceTest.LastModified,
	) {
		return errors.Wrap(rserrors.ErrInvalidParameter, "Endpoint")
	}
	if err := webServiceTest.Path.Validate(); err != nil {
		return errors.WithStack(err)
	}
	if err := webServiceTest.HttpMethod.Validate(); err != nil {
		return errors.WithStack(err)
	}
	if err := webServiceTest.ContentType.Validate(); err != nil {
		return err
	}
	webServiceTest.SetValidated()
	return nil
}

type WebServiceTestRequest struct {
	Path        rshttp.EndpointPath `json:"path"`
	HttpMethod  rshttp.Method       `json:"http_method"`
	ContentType rshttp.ContentType  `json:"content_type"`
	Desc        string              `json:"desc"`
	RequestData rsjson.MapJson      `json:"request_data" gorm:"Type:JSON"`
	Header      rsjson.MapJson      `json:"header" gorm:"Type:JSON"`
	QueryParam  rsjson.MapJson      `json:"query_param" gorm:"Type:JSON"`
}

func (request WebServiceTestRequest) Validate() error {
	if rsvalid.IsZero(
		request.Path,
		request.HttpMethod,
		request.ContentType,
	) {
		return errors.Wrap(rserrors.ErrInvalidParameter, "EndpointRequest")
	}
	if err := request.Path.Validate(); err != nil {
		return errors.WithStack(err)
	}
	if err := request.HttpMethod.Validate(); err != nil {
		return errors.WithStack(err)
	}
	if err := request.ContentType.Validate(); err != nil {
		return err
	}
	return nil
}

type WebServiceTestListItem struct {
	Id           int64               `json:"id"`
	WebServiceId int64               `json:"-"`
	WebService   *WebService         `json:"web_service" gorm:"foreignkey:WebServiceId"`
	Path         rshttp.EndpointPath `json:"path"`
	HttpMethod   rshttp.Method       `json:"http_method"`
	Desc         string              `json:"desc"`
	Created      time.Time           `json:"created"`
	LastModified time.Time           `json:"last_modified"`
}

func (webServiceTestListItem WebServiceTestListItem) MarshalJSON() ([]byte, error) {
	endpointUrl := &url.URL{
		Scheme: webServiceTestListItem.WebService.HttpSchema,
		Host:   webServiceTestListItem.WebService.Host,
		Path:   webServiceTestListItem.Path.String(),
	}
	return json.Marshal(struct {
		Id           int64               `json:"id"`
		Path         rshttp.EndpointPath `json:"path"`
		Url          string              `json:"url"`
		HttpMethod   rshttp.Method       `json:"http_method"`
		Desc         string              `json:"desc"`
		Created      time.Time           `json:"created"`
		LastModified time.Time           `json:"last_modified"`
	}{
		Id:           webServiceTestListItem.Id,
		Path:         webServiceTestListItem.Path,
		Url:          endpointUrl.String(),
		HttpMethod:   webServiceTestListItem.HttpMethod,
		Desc:         webServiceTestListItem.Desc,
		Created:      webServiceTestListItem.Created,
		LastModified: webServiceTestListItem.LastModified,
	})
}

func (webServiceTestListItem WebServiceTestListItem) TableName() string {
	return "endpoints"
}

func (webServiceTest WebServiceTest) Run() (WebServiceTestResult, error) {

	panic("not implement")
}

type WebServiceTestListRequest struct {
	Page         int
	NumItem      int
	WebServiceId int64
}
