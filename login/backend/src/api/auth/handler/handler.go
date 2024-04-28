package handler

import (
	"encoding/json"
	requestHeader "esystem/src/api/requestheader"
	"esystem/src/pkg/auth"
	"esystem/src/pkg/entity"
	esystemError "esystem/src/pkg/utils/error"
	"esystem/src/pkg/utils/logger"
	"net/http"
)

type AuthHandler struct {
	Service auth.UseCase
	Logs    logger.Logger
}

func MakeAuthHandler(handlerParam HandlerParam) *AuthHandler {
	return &AuthHandler{
		Service: handlerParam.Service,
		Logs:    handlerParam.Logs,
	}
}

func (handler *AuthHandler) Login(response http.ResponseWriter, request *http.Request) {
	response.Header().Set(requestHeader.ContentType, requestHeader.ApplicationJson)
	var body entity.Login

	errBody := json.NewDecoder(request.Body).Decode(&body)
	if errBody != nil {
		go handler.Logs.Error(request.URL.Path, "auth_handler_login_body_badRequest", request.Body, errBody)
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(esystemError.NewError(esystemError.BadRequest, esystemError.BadRequestMessage))
		return
	}

	res, err := handler.Service.Login(body)

	if err != nil {
		switch *err {
		case esystemError.InternalServerError:
			go handler.Logs.Error(request.URL.Path, "auth_handler_login_internalServerError", body, err)
			statusCode, err := esystemError.MapMessageError(esystemError.InternalServerError, "en")
			response.WriteHeader(statusCode)
			json.NewEncoder(response).Encode(err)
			return
		case esystemError.DataNotFound:
			go handler.Logs.Error(request.URL.Path, "auth_handler_login_dataNotFound", body, err)
			statusCode, err := esystemError.MapMessageError(esystemError.DataNotFound, "en")
			response.WriteHeader(statusCode)
			json.NewEncoder(response).Encode(err)
			return
		case esystemError.InvalidPassword:
			go handler.Logs.Error(request.URL.Path, "auth_handler_login_invalidPassword", body, err)
			statusCode, err := esystemError.MapMessageError(esystemError.InvalidPassword, "en")
			response.WriteHeader(statusCode)
			json.NewEncoder(response).Encode(err)
			return
		case esystemError.UserActiveFailed:
			go handler.Logs.Error(request.URL.Path, "auth_handler_login_userActiveFailed", body, err)
			statusCode, err := esystemError.MapMessageError(esystemError.UserActiveFailed, "en")
			response.WriteHeader(statusCode)
			json.NewEncoder(response).Encode(err)
			return
		}
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(&res)
}
