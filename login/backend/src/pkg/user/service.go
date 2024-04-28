package user

import (
	"esystem/src/pkg/entity"
	"esystem/src/pkg/utils/convert"
	esystemError "esystem/src/pkg/utils/error"
	"esystem/src/pkg/utils/logger"
	"fmt"

	utils "esystem/src/pkg/utils/encode"
)

type UseCase interface {
	ValidateLoginExternal(logIn entity.Login) (*entity.User, *string, *esystemError.ErrorCode)
	Add(user entity.User) *esystemError.ErrorCode
}

type Service struct {
	UseCase
	Repo        Repository
	Logs        logger.Logger
	UtilService utils.UseCase
}

func NewService(serviceParam ServiceParam) *Service {
	service := &Service{
		Repo:        serviceParam.Repo,
		Logs:        serviceParam.Logs,
		UtilService: serviceParam.UtilService,
	}
	service.UseCase = service
	return service
}

func (service *Service) ValidateLoginExternal(logIn entity.Login) (*entity.User, *string, *esystemError.ErrorCode) {
	res, err := service.Repo.FindUserByUserName(logIn.UserName)
	if err != nil {
		return nil, nil, convert.ValueToErrorCodePointer(esystemError.InternalServerError)
	}

	if res == nil {
		return nil, nil, convert.ValueToErrorCodePointer(esystemError.DataNotFound)
	}

	errPwd := service.UtilService.ValidatePassword([]byte(res.Password), []byte(logIn.Password))
	if errPwd != nil {
		go service.Logs.Error("", "User_Service_ValidateLoginExternal_ValidatePassword_InvalidPassword", res.Password, errPwd)
		return nil, nil, convert.ValueToErrorCodePointer(esystemError.InvalidPassword)
	}

	resGenerateJWT, errGenerateToken := service.UtilService.GenerateJWT(res)
	if errGenerateToken != nil {
		go service.Logs.Error("", "User_service_login_generateJWT_internalServerError", errGenerateToken, logIn)
		return nil, nil, convert.ValueToErrorCodePointer(esystemError.InternalServerError)
	}

	return res, &resGenerateJWT, nil
}

func (service *Service) Add(user entity.User) *esystemError.ErrorCode {

	resFindUser, errFindUser := service.Repo.FindUserByUserName(user.UserName)
	if errFindUser != nil {
		return convert.ValueToErrorCodePointer(esystemError.InternalServerError)
	}

	pwd, errGenerateFromPassword := service.UtilService.GenerateFromPassword([]byte(user.Password))
	if errGenerateFromPassword != nil {
		go service.Logs.Error("", "User_Service_UpsertUser_GenerateFromPassword_InternalServerError", user.Password, errGenerateFromPassword)
		return convert.ValueToErrorCodePointer(esystemError.InternalServerError)
	}

	user.Password = pwd

	if resFindUser != nil {
		if resFindUser.UserName == user.UserName {
			return convert.ValueToErrorCodePointer(esystemError.ExistUserName)
		}
	}

	fmt.Println(user.UserName)

	err := service.Repo.UpsertUser(user)
	if err != nil {
		return convert.ValueToErrorCodePointer(esystemError.InternalServerError)
	}

	return nil
}
