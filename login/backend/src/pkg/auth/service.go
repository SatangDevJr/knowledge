package auth

import (
	"esystem/src/pkg/entity"
	"esystem/src/pkg/user"
	"esystem/src/pkg/utils/logger"

	"esystem/src/pkg/utils/convert"
	utilsDecode "esystem/src/pkg/utils/decode"
	esystemError "esystem/src/pkg/utils/error"
)

type UseCase interface {
	Login(login entity.Login) (entity.UserInfo, *esystemError.ErrorCode)
}

type Service struct {
	UseCase
	UtilDecodeService utilsDecode.UseCase
	UserService       user.UseCase
	Repo              Repository
	Logs              logger.Logger
}

func NewService(serviceParam ServiceParam) *Service {
	service := &Service{
		UtilDecodeService: serviceParam.UtilDecodeService,
		UserService:       serviceParam.UserService,
		Repo:              serviceParam.Repo,
		Logs:              serviceParam.Logs,
	}
	service.UseCase = service
	return service
}

func (service *Service) Login(login entity.Login) (entity.UserInfo, *esystemError.ErrorCode) {
	var user *entity.User
	token := ""

	resLoginExternal, resToken, errLoginExternal := service.UserService.ValidateLoginExternal(login)
	if errLoginExternal != nil {
		switch *errLoginExternal {
		case esystemError.InternalServerError:
			go service.Logs.Error("", "auth_service_login_loginExternal_internalServerError", errLoginExternal, login)
			return entity.UserInfo{}, errLoginExternal
		case esystemError.DataNotFound:
			go service.Logs.Error("", "auth_service_login_loginExternal_dataNotFound", errLoginExternal, login)
			return entity.UserInfo{}, errLoginExternal
		case esystemError.InvalidPassword:
			go service.Logs.Error("", "auth_service_login_loginExternal_invalidPassword", errLoginExternal, login)
			return entity.UserInfo{}, errLoginExternal
		}
	}

	user = resLoginExternal
	token = convert.StringPointerToValue(resToken)

	if user == nil {
		go service.Logs.Error("", "auth_service_login_user", "User is null", login)
		return entity.UserInfo{}, convert.ValueToErrorCodePointer(esystemError.InternalServerError)
	}

	userInfo := entity.UserInfo{
		User:  user,
		Token: token,
	}

	return userInfo, nil
}
