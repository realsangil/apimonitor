package services

import (
	"reflect"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/realsangil/apimonitor/models"
	"github.com/realsangil/apimonitor/pkg/amerr"
	"github.com/realsangil/apimonitor/pkg/rsdb"
	mocks2 "github.com/realsangil/apimonitor/pkg/rsdb/mocks"
	"github.com/realsangil/apimonitor/pkg/rserrors"
	"github.com/realsangil/apimonitor/pkg/rsmodels"
	"github.com/realsangil/apimonitor/pkg/testutils"
	"github.com/realsangil/apimonitor/repositories"
	"github.com/realsangil/apimonitor/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type webServiceMockFunc func(mockTx *mocks2.Connection, mockWebServiceRepository *mocks.WebServiceRepository)

var monkeyWebServiceServiceGetWebServiceById = func(webService *models.WebService, err *amerr.ErrorWithLanguage) {
	monkey.UnpatchInstanceMethod(reflect.TypeOf(&WebServiceServiceImpl{}), "GetWebServiceById")
	monkey.PatchInstanceMethod(reflect.TypeOf(&WebServiceServiceImpl{}), "GetWebServiceById", func(service *WebServiceServiceImpl, ws *models.WebService) *amerr.ErrorWithLanguage {
		*ws = *webService
		return err
	})
}

func TestWebServiceServiceImpl_CreateWebService(t *testing.T) {
	testutils.MonkeyAll()

	validatedRequest := models.WebServiceRequest{
		Host:        "https://realsangil.github.io",
		Description: "sangil's dev blog",
		Favicon:     "",
	}

	validatedWebServiceWithoutId := &models.WebService{
		DefaultValidateChecker: rsmodels.ValidatedDefaultValidateChecker,
		Host:                   "realsangil.github.io",
		Schema:                 "https",
		Description:            "sangil's dev blog",
		Favicon:                "",
		CreatedAt:              time.Now(),
		ModifiedAt:             time.Now(),
	}

	validatedWebService := &models.WebService{
		DefaultValidateChecker: rsmodels.ValidatedDefaultValidateChecker,
		Id:                     1,
		Host:                   "realsangil.github.io",
		Schema:                 "https",
		Description:            "sangil's dev blog",
		Favicon:                "",
		CreatedAt:              time.Now(),
		ModifiedAt:             time.Now(),
	}

	mockFuncWithError := func(err error) webServiceMockFunc {
		return func(mockTx *mocks2.Connection, mockWebServiceRepository *mocks.WebServiceRepository) {
			m := mockWebServiceRepository.On("Create", rsdb.GetConnection(), validatedWebServiceWithoutId)
			if err == nil {
				m.Run(func(args mock.Arguments) {
					arg := args.Get(1).(*models.WebService)
					arg.Id = 1
				})
			}
			m.Return(err)
		}
	}

	type args struct {
		request models.WebServiceRequest
	}
	tests := []struct {
		name     string
		args     args
		mockFunc webServiceMockFunc
		want     *models.WebService
		wantErr  *amerr.ErrorWithLanguage
	}{
		{
			name: "pass_https_host",
			args: args{
				request: validatedRequest,
			},
			mockFunc: mockFuncWithError(nil),
			want:     validatedWebService,
			wantErr:  nil,
		},
		{
			name: "invalid request",
			args: args{
				// request: validatedRequest,
			},
			mockFunc: mockFuncWithError(nil),
			want:     nil,
			wantErr:  amerr.GetErrInternalServer(),
		},
		{
			name: "invalid_host",
			args: args{

				request: models.WebServiceRequest{
					Host:        "ftp://realsangil.github.io",
					Description: "sangil's dev blog",
					Favicon:     "",
				},
			},
			mockFunc: func(mockTx *mocks2.Connection, mockWebServiceRepository *mocks.WebServiceRepository) {
			},
			want:    nil,
			wantErr: amerr.GetErrorsFromCode(amerr.ErrBadRequest),
		},
		{
			name: "duplicated_web_service",
			args: args{
				request: validatedRequest,
			},
			mockFunc: mockFuncWithError(rsdb.ErrDuplicateData),
			want:     nil,
			wantErr:  amerr.GetErrorsFromCode(amerr.ErrDuplicatedWebService),
		},
		{
			name: "data too long",
			args: args{

				request: validatedRequest,
			},
			mockFunc: mockFuncWithError(rsdb.ErrInvalidData),
			want:     nil,
			wantErr:  amerr.GetErrorsFromCode(amerr.ErrBadRequest),
		},
		{
			name: "unexpected error",
			args: args{

				request: validatedRequest,
			},
			mockFunc: mockFuncWithError(rserrors.ErrUnexpected),
			want:     nil,
			wantErr:  amerr.GetErrorsFromCode(amerr.ErrInternalServer),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockWebServiceRepository := &mocks.WebServiceRepository{}
			service := &WebServiceServiceImpl{
				webServiceRepository: mockWebServiceRepository,
			}

			mockConn := &mocks2.Connection{}
			testutils.MonkeyGetConnection(mockConn)
			tt.mockFunc(mockConn, mockWebServiceRepository)
			got, gotErr := service.CreateWebService(tt.args.request)
			assert.Equal(t, tt.wantErr, gotErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNewWebServiceService(t *testing.T) {
	testutils.MonkeyAll()

	mockWebServiceRepository := &mocks.WebServiceRepository{}

	type args struct {
		webServiceRepository repositories.WebServiceRepository
	}
	tests := []struct {
		name    string
		args    args
		want    WebServiceService
		wantErr bool
	}{
		{
			name: "pass",
			args: args{
				webServiceRepository: mockWebServiceRepository,
			},
			want: &WebServiceServiceImpl{
				webServiceRepository: mockWebServiceRepository,
			},
			wantErr: false,
		},
		{
			name: "invalid WebServiceRepository",
			args: args{
				webServiceRepository: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid WebServiceRepository",
			args: args{
				webServiceRepository: repositories.WebServiceRepository(nil),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewWebServiceService(tt.args.webServiceRepository)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWebServiceService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWebServiceService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWebServiceServiceImpl_GetWebServiceById(t *testing.T) {
	testutils.MonkeyAll()

	webServiceWithId := &models.WebService{Id: 1}
	webService := models.WebService{
		DefaultValidateChecker: rsmodels.ValidatedDefaultValidateChecker,
		Id:                     1,
		Host:                   "realsangil.github.io",
		Schema:                 "https",
		Description:            "sangil's dev blog",
		Favicon:                "",
		CreatedAt:              time.Now(),
		ModifiedAt:             time.Now(),
	}

	type args struct {
		webService *models.WebService
	}
	tests := []struct {
		name     string
		args     args
		mockFunc webServiceMockFunc
		wantErr  *amerr.ErrorWithLanguage
	}{
		{
			name: "pass",
			args: args{

				webService: webServiceWithId,
			},
			mockFunc: func(mockTx *mocks2.Connection, mockWebServiceRepository *mocks.WebServiceRepository) {
				mockWebServiceRepository.On("GetById", rsdb.GetConnection(), webServiceWithId).Run(func(args mock.Arguments) {
					arg := args.Get(1).(*models.WebService)
					*arg = webService
				}).Return(nil)
			},
			wantErr: nil,
		},
		{
			name: "web service not found",
			args: args{

				webService: webServiceWithId,
			},
			mockFunc: func(mockTx *mocks2.Connection, mockWebServiceRepository *mocks.WebServiceRepository) {
				mockWebServiceRepository.On("GetById", rsdb.GetConnection(), webServiceWithId).Return(rsdb.ErrRecordNotFound)
			},
			wantErr: amerr.GetErrorsFromCode(amerr.ErrWebServiceNotFound),
		},
		{
			name: "repository unexpected error",
			args: args{

				webService: webServiceWithId,
			},
			mockFunc: func(mockTx *mocks2.Connection, mockWebServiceRepository *mocks.WebServiceRepository) {
				mockWebServiceRepository.On("GetById", rsdb.GetConnection(), webServiceWithId).Return(rserrors.ErrUnexpected)
			},
			wantErr: amerr.GetErrInternalServer(),
		},
		{
			name: "invalid parameter",
			args: args{
				webService: nil,
			},
			mockFunc: func(mockTx *mocks2.Connection, mockWebServiceRepository *mocks.WebServiceRepository) {
			},
			wantErr: amerr.GetErrInternalServer(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockWebServiceRepository := &mocks.WebServiceRepository{}
			service := &WebServiceServiceImpl{
				webServiceRepository: mockWebServiceRepository,
			}
			mockConn := &mocks2.Connection{}
			testutils.MonkeyGetConnection(mockConn)
			tt.mockFunc(mockConn, mockWebServiceRepository)
			gotErr := service.GetWebServiceById(tt.args.webService)
			assert.Equal(t, tt.wantErr, gotErr)
		})
	}
}

func TestWebServiceServiceImpl_DeleteWebServiceById(t *testing.T) {
	testutils.MonkeyAll()

	webServiceWithId := &models.WebService{Id: 1}
	webService := &models.WebService{
		Id:          1,
		Host:        "realsangil.github.io",
		Schema:      "https",
		Description: "sangil's dev blog",
		Favicon:     "",
		CreatedAt:   time.Now(),
		ModifiedAt:  time.Now(),
	}

	type args struct {
		webService *models.WebService
	}
	tests := []struct {
		name     string
		args     args
		mockFunc webServiceMockFunc
		wantErr  *amerr.ErrorWithLanguage
	}{
		{
			name: "pass",
			args: args{

				webService: webServiceWithId,
			},
			mockFunc: func(mockTx *mocks2.Connection, mockWebServiceRepository *mocks.WebServiceRepository) {
				monkeyWebServiceServiceGetWebServiceById(webService, nil)
				mockWebServiceRepository.On("DeleteById", rsdb.GetConnection(), webService).Return(nil)
			},
			wantErr: nil,
		},
		{
			name: "invalid parameter",
			args: args{
				// test: webServiceWithId,
			},
			mockFunc: func(mockTx *mocks2.Connection, mockWebServiceRepository *mocks.WebServiceRepository) {
			},
			wantErr: amerr.GetErrInternalServer(),
		},
		{
			name: "[WebServiceService.GetWebServiceById] test not found",
			args: args{

				webService: webServiceWithId,
			},
			mockFunc: func(mockTx *mocks2.Connection, mockWebServiceRepository *mocks.WebServiceRepository) {
				monkeyWebServiceServiceGetWebServiceById(webService, amerr.GetErrorsFromCode(amerr.ErrWebServiceNotFound))
			},
			wantErr: amerr.GetErrorsFromCode(amerr.ErrWebServiceNotFound),
		},
		{
			name: "[WebServiceService.GetWebServiceById] unexpected error",
			args: args{

				webService: webServiceWithId,
			},
			mockFunc: func(mockTx *mocks2.Connection, mockWebServiceRepository *mocks.WebServiceRepository) {
				monkeyWebServiceServiceGetWebServiceById(webService, amerr.GetErrInternalServer())
			},
			wantErr: amerr.GetErrInternalServer(),
		},
		{
			name: "[WebServiceRepository.DeleteById] test not found",
			args: args{
				webService: webServiceWithId,
			},
			mockFunc: func(mockTx *mocks2.Connection, mockWebServiceRepository *mocks.WebServiceRepository) {
				monkeyWebServiceServiceGetWebServiceById(webService, nil)
				mockWebServiceRepository.On("DeleteById", rsdb.GetConnection(), webService).Return(rsdb.ErrRecordNotFound)
			},
			wantErr: amerr.GetErrorsFromCode(amerr.ErrWebServiceNotFound),
		},
		{
			name: "[WebServiceRepository.DeleteById] unexpected error",
			args: args{
				webService: webServiceWithId,
			},
			mockFunc: func(mockTx *mocks2.Connection, mockWebServiceRepository *mocks.WebServiceRepository) {
				monkeyWebServiceServiceGetWebServiceById(webService, nil)
				mockWebServiceRepository.On("DeleteById", rsdb.GetConnection(), webService).Return(rserrors.ErrUnexpected)
			},
			wantErr: amerr.GetErrInternalServer(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockWebServiceRepository := &mocks.WebServiceRepository{}

			service := &WebServiceServiceImpl{
				webServiceRepository: mockWebServiceRepository,
			}
			mockConn := &mocks2.Connection{}
			testutils.MonkeyGetConnection(mockConn)
			tt.mockFunc(mockConn, mockWebServiceRepository)
			gotErr := service.DeleteWebServiceById(tt.args.webService)
			assert.Equal(t, tt.wantErr, gotErr)
		})
	}
}

func TestWebServiceServiceImpl_UpdateWebServiceById(t *testing.T) {
	testutils.MonkeyAll()

	webServiceWithId := &models.WebService{Id: 1}
	webService := &models.WebService{
		Id:          1,
		Host:        "realsangil.github.io",
		Schema:      "https",
		Description: "sangil's dev blog",
		Favicon:     "",
		CreatedAt:   time.Now(),
		ModifiedAt:  time.Now(),
	}
	request := models.WebServiceRequest{
		Host:        "https://www.lalaworks.com",
		Description: "lalaworks website",
		Favicon:     "",
	}
	updatedWebService := &models.WebService{
		Id:          1,
		Host:        "www.lalaworks.com",
		Schema:      "https",
		Description: "lalaworks website",
		Favicon:     "",
		CreatedAt:   time.Now(),
		ModifiedAt:  time.Now(),
	}

	type args struct {
		transaction rsdb.Connection
		webService  *models.WebService
		request     models.WebServiceRequest
	}
	tests := []struct {
		name     string
		args     args
		mockFunc webServiceMockFunc
		wantErr  *amerr.ErrorWithLanguage
	}{
		{
			name: "pass",
			args: args{

				webService: webServiceWithId,
				request:    request,
			},
			mockFunc: func(mockTx *mocks2.Connection, mockWebServiceRepository *mocks.WebServiceRepository) {
				monkeyWebServiceServiceGetWebServiceById(webService, nil)
				mockWebServiceRepository.On("Save", rsdb.GetConnection(), updatedWebService).Return(nil)
			},
			wantErr: nil,
		},
		{
			name: "[WebServiceService.GetWebServiceById] test not found",
			args: args{

				webService: webServiceWithId,
				request:    request,
			},
			mockFunc: func(mockTx *mocks2.Connection, mockWebServiceRepository *mocks.WebServiceRepository) {
				monkeyWebServiceServiceGetWebServiceById(webService, amerr.GetErrorsFromCode(amerr.ErrWebServiceNotFound))
			},
			wantErr: amerr.GetErrorsFromCode(amerr.ErrWebServiceNotFound),
		},
		{
			name: "[WebServiceService.GetWebServiceById] unexpected error",
			args: args{

				webService: webServiceWithId,
				request:    request,
			},
			mockFunc: func(mockTx *mocks2.Connection, mockWebServiceRepository *mocks.WebServiceRepository) {
				monkeyWebServiceServiceGetWebServiceById(webService, amerr.GetErrInternalServer())
			},
			wantErr: amerr.GetErrInternalServer(),
		},
		{
			name: "[WebService.UpdateFromRequest] invalid request",
			args: args{
				webService: webServiceWithId,
				request:    models.WebServiceRequest{},
			},
			mockFunc: func(mockTx *mocks2.Connection, mockWebServiceRepository *mocks.WebServiceRepository) {
				monkeyWebServiceServiceGetWebServiceById(webService, nil)
			},
			wantErr: amerr.GetErrorsFromCode(amerr.ErrBadRequest),
		},
		{
			name: "[WebServiceRepository.Save] not found",
			args: args{

				webService: webServiceWithId,
				request:    request,
			},
			mockFunc: func(mockTx *mocks2.Connection, mockWebServiceRepository *mocks.WebServiceRepository) {
				monkeyWebServiceServiceGetWebServiceById(webService, nil)
				mockWebServiceRepository.On("Save", rsdb.GetConnection(), updatedWebService).Return(rsdb.ErrRecordNotFound)
			},
			wantErr: amerr.GetErrorsFromCode(amerr.ErrWebServiceNotFound),
		},
		{
			name: "[WebServiceRepository.Save] unexpected error",
			args: args{

				webService: webServiceWithId,
				request:    request,
			},
			mockFunc: func(mockTx *mocks2.Connection, mockWebServiceRepository *mocks.WebServiceRepository) {
				monkeyWebServiceServiceGetWebServiceById(webService, nil)
				mockWebServiceRepository.On("Save", rsdb.GetConnection(), updatedWebService).Return(rserrors.ErrUnexpected)
			},
			wantErr: amerr.GetErrInternalServer(),
		},
		{
			name: "invalid parameter",
			args: args{
				// test: webServiceWithId,
				request: request,
			},
			mockFunc: func(mockTx *mocks2.Connection, mockWebServiceRepository *mocks.WebServiceRepository) {
			},
			wantErr: amerr.GetErrInternalServer(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockWebServiceRepository := &mocks.WebServiceRepository{}

			service := &WebServiceServiceImpl{
				webServiceRepository: mockWebServiceRepository,
			}

			mockConn := &mocks2.Connection{}
			testutils.MonkeyGetConnection(mockConn)
			tt.mockFunc(mockConn, mockWebServiceRepository)
			err := service.UpdateWebServiceById(tt.args.webService, tt.args.request)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestWebServiceServiceImpl_GetWebServiceList(t *testing.T) {
	testutils.MonkeyAll()

	request := models.WebServiceListRequest{
		Page:          1,
		NumItem:       20,
		SearchKeyword: "",
	}

	paginatedList := []*models.WebService{
		{
			Id:          1,
			Host:        "realsangil.github.io",
			Schema:      "https",
			Description: "sangil's dev blog",
			Favicon:     "",
			CreatedAt:   time.Now(),
			ModifiedAt:  time.Now(),
		},
	}

	type args struct {
		transaction rsdb.Connection
		request     models.WebServiceListRequest
	}
	tests := []struct {
		name     string
		args     args
		mockFunc webServiceMockFunc
		want     *rsmodels.PaginatedList
		wantErr  *amerr.ErrorWithLanguage
	}{
		{
			name: "pass",
			args: args{

				request: request,
			},
			mockFunc: func(mockTx *mocks2.Connection, mockWebServiceRepository *mocks.WebServiceRepository) {
				mockWebServiceRepository.On(
					"List",
					rsdb.GetConnection(),
					&[]*models.WebService{},
					rsdb.ListFilter{
						NumItem:    20,
						Page:       1,
						Conditions: map[string]interface{}{},
					}, rsdb.Orders{
						rsdb.Order{
							Field: "host",
							IsASC: true,
						},
					}).Run(func(args mock.Arguments) {
					arg := args.Get(1).(*[]*models.WebService)
					*arg = paginatedList
				}).Return(1, nil)
			},
			want: &rsmodels.PaginatedList{
				CurrentPage: 1,
				TotalCount:  1,
				NumItem:     20,
				Items:       paginatedList,
			},
			wantErr: nil,
		},
		{
			name: "[WebServiceService.List] unexpected error",
			args: args{

				request: request,
			},
			mockFunc: func(mockTx *mocks2.Connection, mockWebServiceRepository *mocks.WebServiceRepository) {
				mockWebServiceRepository.On(
					"List",
					rsdb.GetConnection(),
					&[]*models.WebService{},
					rsdb.ListFilter{
						NumItem:    20,
						Page:       1,
						Conditions: map[string]interface{}{},
					}, rsdb.Orders{
						rsdb.Order{
							Field: "host",
							IsASC: true,
						},
					}).Return(0, rserrors.ErrUnexpected)
			},
			want:    nil,
			wantErr: amerr.GetErrInternalServer(),
		},
		{
			name: "invalid parameter",
			args: args{
				// request: request,
			},
			mockFunc: func(mockTx *mocks2.Connection, mockWebServiceRepository *mocks.WebServiceRepository) {
			},
			want:    nil,
			wantErr: amerr.GetErrInternalServer(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockWebServiceRepository := &mocks.WebServiceRepository{}

			service := &WebServiceServiceImpl{
				webServiceRepository: mockWebServiceRepository,
			}

			mockConn := &mocks2.Connection{}
			testutils.MonkeyGetConnection(mockConn)
			tt.mockFunc(mockConn, mockWebServiceRepository)

			got, err := service.GetWebServiceList(tt.args.request)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
