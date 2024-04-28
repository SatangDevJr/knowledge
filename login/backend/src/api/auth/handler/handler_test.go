package handler_test

import (
	"bytes"
	"encoding/json"
	"esystem/src/api/auth/handler"
	requestHeader "esystem/src/api/requestheader"
	"esystem/src/pkg/auth/mocks"
	"esystem/src/pkg/entity"
	esystemError "esystem/src/pkg/utils/error"
	loggerMocks "esystem/src/pkg/utils/logger/mocks"
	"esystem/src/pkg/utils/mocker"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	uri         string
	service     *mocks.UseCase
	authHandler *handler.AuthHandler
	logs        *loggerMocks.Logger
	recorder    *httptest.ResponseRecorder
	request     *http.Request
	router      *mux.Router

	mockServiceLogin *mocker.MockCall
)

func callServiceLogin() *mock.Call {
	return service.On("Login", mock.Anything)
}

func beforeEach() {
	uri = "/auth"
	service = &mocks.UseCase{}
	logs = &loggerMocks.Logger{}

	logs.On("Error", mock.Anything, mock.Anything, mock.Anything, mock.Anything)

	authHandler = &handler.AuthHandler{
		Service: service,
		Logs:    logs,
	}
}

func TestHandler_MakeAuthHandler(t *testing.T) {

	t.Run("should return struct auth handler when call make auth handler", func(t *testing.T) {
		beforeEach()
		handlerParam := handler.HandlerParam{
			Service: service,
			Logs:    logs,
		}

		handlerMakePartner := handler.MakeAuthHandler(handlerParam)

		expectedResult := &handler.AuthHandler{
			Service: handlerParam.Service,
			Logs:    handlerParam.Logs,
		}

		assert.Equal(t, expectedResult, handlerMakePartner)
	})

}

func TestHandler_Login(t *testing.T) {
	beforeEachLogin := func() {
		beforeEach()
		router = mux.NewRouter()
		router.HandleFunc(uri, authHandler.Login)
		recorder = httptest.NewRecorder()
		request = httptest.NewRequest(http.MethodPost, uri, nil)

		mockServiceLogin = mocker.NewMockCall(callServiceLogin)
		mockServiceLogin.Return(entity.UserInfo{}, nil)
	}

	t.Run("should response bad request when request service login with body wrong format", func(t *testing.T) {
		beforeEachLogin()
		requestBody := ``
		request = httptest.NewRequest(http.MethodPost, uri, bytes.NewBuffer([]byte(requestBody)))

		router.ServeHTTP(recorder, request)

		var responseBody esystemError.Error
		json.NewDecoder(recorder.Body).Decode(&responseBody)
		err := esystemError.NewError(esystemError.BadRequest, esystemError.BadRequestMessage)
		assert.Equal(t, *err, responseBody)
		assert.Equal(t, requestHeader.ApplicationJson, recorder.Header().Get(requestHeader.ContentType))
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("should call service login when request service", func(t *testing.T) {
		beforeEachLogin()
		parameter := mockUserLogin()
		jsonMock, _ := json.Marshal(parameter)
		requestBody := jsonMock
		request = httptest.NewRequest(http.MethodPost, uri, bytes.NewBuffer([]byte(requestBody)))

		router.ServeHTTP(recorder, request)

		service.AssertCalled(t, "Login", parameter)
	})

	t.Run("should response internal server error when service login fail", func(t *testing.T) {
		beforeEachLogin()
		parameter := mockUserLogin()
		jsonMock, _ := json.Marshal(parameter)
		requestBody := jsonMock
		request = httptest.NewRequest(http.MethodPost, uri, bytes.NewBuffer([]byte(requestBody)))
		err := esystemError.InternalServerError
		mockServiceLogin.Return(entity.UserInfo{}, &err)

		router.ServeHTTP(recorder, request)

		expectedStatusCode, expectedError := esystemError.MapMessageError(esystemError.InternalServerError, "en")
		var body esystemError.Error
		json.NewDecoder(recorder.Body).Decode(&body)
		assert.Equal(t, expectedError, body)
		assert.Equal(t, expectedStatusCode, recorder.Code)
		assert.Equal(t, requestHeader.ApplicationJson, recorder.Header().Get(requestHeader.ContentType))
	})

	t.Run("should response data not found when service login fail", func(t *testing.T) {
		beforeEachLogin()
		parameter := mockUserLogin()
		jsonMock, _ := json.Marshal(parameter)
		requestBody := jsonMock
		request = httptest.NewRequest(http.MethodPost, uri, bytes.NewBuffer([]byte(requestBody)))
		err := esystemError.DataNotFound
		mockServiceLogin.Return(entity.UserInfo{}, &err)

		router.ServeHTTP(recorder, request)

		expectedStatusCode, expectedError := esystemError.MapMessageError(esystemError.DataNotFound, "en")
		var body esystemError.Error
		json.NewDecoder(recorder.Body).Decode(&body)
		assert.Equal(t, expectedError, body)
		assert.Equal(t, expectedStatusCode, recorder.Code)
		assert.Equal(t, requestHeader.ApplicationJson, recorder.Header().Get(requestHeader.ContentType))
	})

	t.Run("should response invalid password when service login fail", func(t *testing.T) {
		beforeEachLogin()
		parameter := mockUserLogin()
		jsonMock, _ := json.Marshal(parameter)
		requestBody := jsonMock
		request = httptest.NewRequest(http.MethodPost, uri, bytes.NewBuffer([]byte(requestBody)))
		err := esystemError.InvalidPassword
		mockServiceLogin.Return(entity.UserInfo{}, &err)

		router.ServeHTTP(recorder, request)

		expectedStatusCode, expectedError := esystemError.MapMessageError(esystemError.InvalidPassword, "en")
		var body esystemError.Error
		json.NewDecoder(recorder.Body).Decode(&body)
		assert.Equal(t, expectedError, body)
		assert.Equal(t, expectedStatusCode, recorder.Code)
		assert.Equal(t, requestHeader.ApplicationJson, recorder.Header().Get(requestHeader.ContentType))
	})

	t.Run("should response user active failed when service login fail", func(t *testing.T) {
		beforeEachLogin()
		parameter := mockUserLogin()
		jsonMock, _ := json.Marshal(parameter)
		requestBody := jsonMock
		request = httptest.NewRequest(http.MethodPost, uri, bytes.NewBuffer([]byte(requestBody)))
		err := esystemError.UserActiveFailed
		mockServiceLogin.Return(entity.UserInfo{}, &err)

		router.ServeHTTP(recorder, request)

		expectedStatusCode, expectedError := esystemError.MapMessageError(esystemError.UserActiveFailed, "en")
		var body esystemError.Error
		json.NewDecoder(recorder.Body).Decode(&body)
		assert.Equal(t, expectedError, body)
		assert.Equal(t, expectedStatusCode, recorder.Code)
		assert.Equal(t, requestHeader.ApplicationJson, recorder.Header().Get(requestHeader.ContentType))
	})

	t.Run("should return service login when call service login success", func(t *testing.T) {
		beforeEachLogin()
		parameter := mockUserLogin()
		jsonMock, _ := json.Marshal(parameter)
		requestBody := jsonMock
		request = httptest.NewRequest(http.MethodPost, uri, bytes.NewBuffer([]byte(requestBody)))

		resLogin := entity.UserInfo{
			User: &entity.User{
				ID:       1,
				UserName: "ATPHQ",
			},
		}
		mockServiceLogin.Return(resLogin, nil)

		router.ServeHTTP(recorder, request)

		var body entity.UserInfo
		json.NewDecoder(recorder.Body).Decode(&body)
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, resLogin, body)
		assert.Equal(t, requestHeader.ApplicationJson, recorder.Header().Get(requestHeader.ContentType))
	})
}

func mockUserLogin() entity.Login {
	return entity.Login{
		UserName: "ATPHQ",
		Password: "123456",
	}
}
