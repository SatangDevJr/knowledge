package handler

import (
	"encoding/json"
	requestHeader "esystem/src/api/requestheader"
	"esystem/src/pkg/entity"
	"esystem/src/pkg/user"
	"esystem/src/pkg/utils/logger"
	"net/http"

	esystemError "esystem/src/pkg/utils/error"
)

type UserHandler struct {
	Service user.UseCase
	Logs    logger.Logger
}

func MakeUserHandler(handlerParam HandlerParam) *UserHandler {
	return &UserHandler{
		Service: handlerParam.Service,
		Logs:    handlerParam.Logs,
	}
}

func (handler *UserHandler) Add(response http.ResponseWriter, request *http.Request) {
	response.Header().Set(requestHeader.ContentType, requestHeader.ApplicationJson)

	var body entity.User

	errBody := json.NewDecoder(request.Body).Decode(&body)
	if errBody != nil || body.UserName == "" {
		go handler.Logs.Error(request.URL.Path, "user_handler_add_badRequest", request.Body, errBody)
		httpStatusCode, errMap := esystemError.MapMessageError(esystemError.BadRequest, "en")
		response.WriteHeader(httpStatusCode)
		json.NewEncoder(response).Encode(&errMap)
		return
	}

	err := handler.Service.Add(body)
	if err != nil {
		switch *err {
		case esystemError.InternalServerError:
			go handler.Logs.Error(request.URL.Path, "user_handler_add_internalServerError", request.Body, err)
			httpStatusCode, errMap := esystemError.MapMessageError(esystemError.InternalServerError, "en")
			response.WriteHeader(httpStatusCode)
			json.NewEncoder(response).Encode(&errMap)
			return
		case esystemError.ExistUserName:
			httpStatusCode, errMap := esystemError.MapMessageError(esystemError.ExistUserName, "th")
			response.WriteHeader(httpStatusCode)
			json.NewEncoder(response).Encode(&errMap)
			return
		case esystemError.DataNotFound:
			httpStatusCode, errMap := esystemError.MapMessageError(esystemError.DataNotFound, "en")
			response.WriteHeader(httpStatusCode)
			json.NewEncoder(response).Encode(&errMap)
			return
		}
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(map[string]interface{}{})
}
