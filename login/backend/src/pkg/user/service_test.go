package user_test

import (
	"errors"
	"esystem/src/pkg/entity"
	"esystem/src/pkg/user"
	"esystem/src/pkg/user/mocks"
	"esystem/src/pkg/utils/convert"
	utilsMock "esystem/src/pkg/utils/encode/mocks"
	esystemError "esystem/src/pkg/utils/error"
	loggerMocks "esystem/src/pkg/utils/logger/mocks"
	"esystem/src/pkg/utils/mocker"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockService *mocks.UseCase
	repository  *mocks.Repository
	service     *user.Service
	utilService *utilsMock.UseCase
	logs        *loggerMocks.Logger

	mockRepoFindUserByRoleCode           *mocker.MockCall
	mockRepoFindUserByUserName           *mocker.MockCall
	mockRepoFindUserByRefId              *mocker.MockCall
	mockRepoUpsertUser                   *mocker.MockCall
	mockRepoUpdatePassword               *mocker.MockCall
	mockRepoSearch                       *mocker.MockCall
	mockServiceUpsertUser                *mocker.MockCall
	mockServiceRoleFindRoleByRoleId      *mocker.MockCall
	mockServiceUserRoleUpsertUserRole    *mocker.MockCall
	mockServiceUtilsGenerateFromPassword *mocker.MockCall
	mockServiceUtilsGenerateJWT          *mocker.MockCall
	mockServiceUtilsValidatePassword     *mocker.MockCall
	mockRepoFindExternalUserByUserName   *mocker.MockCall
	mockRepoUpdateActiveFlagByUserName   *mocker.MockCall
	mockRepoDeleteByID                   *mocker.MockCall
	mockRepoFindUserByID                 *mocker.MockCall
	mockServiceFindUserByID              *mocker.MockCall
	mockServiceUserRoleDeleteByUserID    *mocker.MockCall
)

func callRepoFindUserByUserName() *mock.Call {
	return repository.On("FindUserByUserName", mock.Anything)
}

func callRepoUpsertUser() *mock.Call {
	return repository.On("UpsertUser", mock.Anything)
}

func callServiceUpsertUser() *mock.Call {
	return mockService.On("UpsertUser", mock.Anything)
}

func callServiceUtilsGenerateFromPassword() *mock.Call {
	return utilService.On("GenerateFromPassword", mock.Anything)
}

func callServiceUtilsGenerateJWT() *mock.Call {
	return utilService.On("GenerateJWT", mock.Anything)
}

func callServiceUtilsValidatePassword() *mock.Call {
	return utilService.On("ValidatePassword", mock.Anything, mock.Anything)
}

func callRepoFindExternalUserByUserName() *mock.Call {
	return repository.On("FindExternalUserByUserName", mock.Anything)
}

func callRepoUpdateActiveFlagByUserName() *mock.Call {
	return repository.On("UpdateActiveFlagByUserName", mock.Anything)
}

func callRepoDeleteByID() *mock.Call {
	return repository.On("Delete", mock.Anything)
}

func callRepoFindUserByID() *mock.Call {
	return repository.On("FindUserByID", mock.Anything)
}

func callServiceFindByID() *mock.Call {
	return mockService.On("FindUserByID", mock.Anything)
}

func beforeEach() {
	mockService = &mocks.UseCase{}
	repository = &mocks.Repository{}

	utilService = &utilsMock.UseCase{}
	logs = &loggerMocks.Logger{}

	logs.On("Error", mock.Anything, mock.Anything, mock.Anything, mock.Anything)

	service = &user.Service{
		Repo:        repository,
		UtilService: utilService,
		Logs:        logs,
	}
	service.UseCase = mockService
}

func TestService_NewService(t *testing.T) {
	t.Run("should return struct user service when call new service", func(t *testing.T) {
		beforeEach()

		serviceParam := user.ServiceParam{
			Repo: repository,
			Logs: logs,
		}

		resService := user.NewService(serviceParam)

		expectedService := &user.Service{
			Repo: repository,
			Logs: logs,
		}
		expectedService.UseCase = expectedService

		assert.Equal(t, expectedService, resService)
	})
}

func TestService_ValidateLoginExternal(t *testing.T) {
	beforeEachValidateLoginExternal := func() {
		beforeEach()

		mockRepoFindUserByUserName = mocker.NewMockCall(callRepoFindUserByUserName)
		mockRepoFindUserByUserName.Return(nil, nil)
		mockServiceUtilsValidatePassword = mocker.NewMockCall(callServiceUtilsValidatePassword)
		mockServiceUtilsValidatePassword.Return(nil)
		mockServiceUtilsGenerateJWT = mocker.NewMockCall(callServiceUtilsGenerateJWT)
		mockServiceUtilsGenerateJWT.Return("", nil)
	}

	t.Run("should call repository find user by user name when call service validate login external", func(t *testing.T) {
		beforeEachValidateLoginExternal()
		mockUser := mockUser()
		login := mockLogin()
		mockRepoFindUserByUserName.Return(&mockUser, nil)

		service.ValidateLoginExternal(login)

		repository.AssertCalled(t, "FindUserByUserName", mock.Anything)
	})

	t.Run("should return internal server error when call repository find user by user name failed", func(t *testing.T) {
		beforeEachValidateLoginExternal()
		login := mockLogin()
		mockRepoFindUserByUserName.Return(nil, errors.New("Error"))

		_, _, err := service.ValidateLoginExternal(login)

		expectedError := convert.ValueToErrorCodePointer(esystemError.InternalServerError)
		assert.Equal(t, expectedError, err)
	})

	t.Run("should return data not found when call repository find user by user name not found data", func(t *testing.T) {
		beforeEachValidateLoginExternal()
		login := mockLogin()

		_, _, err := service.ValidateLoginExternal(login)

		expectedError := convert.ValueToErrorCodePointer(esystemError.DataNotFound)
		assert.Equal(t, expectedError, err)
	})

	t.Run("should call utils service validate password when call service validate login external", func(t *testing.T) {
		beforeEachValidateLoginExternal()
		mockUser := mockUser()
		login := mockLogin()
		mockRepoFindUserByUserName.Return(&mockUser, nil)

		service.ValidateLoginExternal(login)

		utilService.AssertCalled(t, "ValidatePassword", mock.Anything, mock.Anything)
	})

	t.Run("should return internal server error when call utils service validate password failed ", func(t *testing.T) {
		beforeEachValidateLoginExternal()
		mockUser := mockUser()
		login := mockLogin()
		mockRepoFindUserByUserName.Return(&mockUser, nil)
		mockServiceUtilsValidatePassword.Return(errors.New("Error"))

		_, _, err := service.ValidateLoginExternal(login)

		expectedError := convert.ValueToErrorCodePointer(esystemError.InvalidPassword)
		assert.Equal(t, expectedError, err)
	})

	t.Run("should call utils service generate JWT when call service validate login external ", func(t *testing.T) {
		beforeEachValidateLoginExternal()
		mockUser := mockUser()
		login := mockLogin()
		mockRepoFindUserByUserName.Return(&mockUser, nil)

		service.ValidateLoginExternal(login)

		utilService.AssertCalled(t, "GenerateJWT", mock.Anything)
	})

	t.Run("should return internal server error when call utils service generate JWT failed ", func(t *testing.T) {
		beforeEachValidateLoginExternal()
		mockUser := mockUser()
		login := mockLogin()
		mockRepoFindUserByUserName.Return(&mockUser, nil)
		mockServiceUtilsGenerateJWT.Return("", errors.New("Error"))

		_, _, err := service.ValidateLoginExternal(login)

		expectedError := convert.ValueToErrorCodePointer(esystemError.InternalServerError)
		assert.Equal(t, expectedError, err)
	})

	t.Run("should return user and JWT when call service find user by ref id success ", func(t *testing.T) {
		beforeEachValidateLoginExternal()
		mockUser := mockUser()
		login := mockLogin()
		mockRepoFindUserByUserName.Return(&mockUser, nil)
		mockServiceUtilsGenerateJWT.Return("test", nil)

		res, resGenerate, _ := service.ValidateLoginExternal(login)

		expectedRes := &mockUser
		expectedResGenerate := "test"
		assert.Equal(t, expectedRes, res)
		assert.Equal(t, expectedResGenerate, *resGenerate)
	})
}

func TestService_Add(t *testing.T) {
	beforeEachAdd := func() {
		beforeEach()

		mockRepoFindUserByUserName = mocker.NewMockCall(callRepoFindUserByUserName)
		mockRepoFindUserByUserName.Return(nil, nil)
		mockRepoUpsertUser = mocker.NewMockCall(callRepoUpsertUser)
		mockRepoUpsertUser.Return(nil, nil)
	}

	t.Run("should call repository find user by username when call service add", func(t *testing.T) {
		beforeEachAdd()
		var mockUser = mockUser()

		mockRepoUpsertUser.Return(&mockUser, nil)

		service.Add(mockUser)

		repository.AssertCalled(t, "FindUserByUserName", mock.Anything)
	})

	t.Run("should response internal server error when call repository find user by user name failed", func(t *testing.T) {
		beforeEachAdd()
		var mockUser = mockUser()
		mockRepoFindUserByUserName.Return(nil, errors.New("Error"))

		err := service.Add(mockUser)

		expectedError := convert.ValueToErrorCodePointer(esystemError.InternalServerError)
		assert.Equal(t, expectedError, err)
	})

	t.Run("should response exist user name when repository find user by user name result user name is duplicate ", func(t *testing.T) {
		beforeEachAdd()
		var mockUser = mockUser()
		mockRepoFindUserByUserName.Return(&mockUser, nil)

		err := service.Add(mockUser)

		expectedError := convert.ValueToErrorCodePointer(esystemError.ExistUserName)
		assert.Equal(t, expectedError, err)
	})

	t.Run("should call repository upsert user when call service find role by role id success", func(t *testing.T) {
		beforeEachAdd()
		var mockUser = mockUser()
		mockRepoUpsertUser.Return(&mockUser, nil)

		service.Add(mockUser)

		repository.AssertCalled(t, "UpsertUser", mockUser)
	})

	t.Run("should response internal server error when call repository upsert user failed", func(t *testing.T) {
		beforeEachAdd()
		var mockUser = mockUser()
		mockRepoUpsertUser.Return(nil, errors.New("Error"))
		err := service.Add(mockUser)

		expectedError := convert.ValueToErrorCodePointer(esystemError.InternalServerError)
		assert.Equal(t, expectedError, err)
	})

	t.Run("should return nil when call service add success", func(t *testing.T) {
		beforeEachAdd()
		var mockUser = mockUser()
		mockRepoUpsertUser.Return(&mockUser, nil)
		err := service.Add(mockUser)

		assert.Nil(t, err)
	})
}

func mockUser() entity.User {
	return entity.User{
		ID:              55,
		UserName:        "AAAHQ",
		Password:        "PASSWORD",
		CreatedByUserId: 1,
		UpdatedByUserId: 1,
	}
}

func mockLogin() entity.Login {
	return entity.Login{
		UserName: "UserName",
		Password: "Password",
	}
}
