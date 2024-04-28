package previewfile_test

import (
	"errors"
	"esystem/src/pkg/utils/convert"
	esystemError "esystem/src/pkg/utils/error"
	fileMock "esystem/src/pkg/utils/file/mocks"
	loggerMocks "esystem/src/pkg/utils/logger/mocks"
	"esystem/src/pkg/utils/mocker"
	previewFile "esystem/src/pkg/utils/previewfile"
	previewFileMock "esystem/src/pkg/utils/previewfile/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockService *previewFileMock.UseCase
	service     *previewFile.Service
	fileService *fileMock.FileService
	logs        *loggerMocks.Logger

	mockServiceFileDownload *mocker.MockCall
)

func beforeEach() {
	mockService = &previewFileMock.UseCase{}
	fileService = &fileMock.FileService{}

	logs = &loggerMocks.Logger{}

	logs.On("Error", mock.Anything, mock.Anything, mock.Anything, mock.Anything)

	service = &previewFile.Service{
		FileService: fileService,
		Logs:        logs,
	}
	service.UseCase = mockService
}

func callServiceFileDownload() *mock.Call {
	return fileService.On("Download", mock.Anything)
}

func TestService_NewService(t *testing.T) {
	t.Run("should return struct vendors service when call new service", func(t *testing.T) {
		beforeEach()

		serviceParam := previewFile.ServiceParam{
			FileService: fileService,
			Logs:        logs,
		}

		resService := previewFile.NewService(serviceParam)

		expectedService := &previewFile.Service{
			Logs:        logs,
			FileService: fileService,
		}
		expectedService.UseCase = expectedService

		assert.Equal(t, expectedService, resService)
	})
}

func TestService_PreviewFile(t *testing.T) {
	beforeEachPreviewFile := func() {
		beforeEach()

		mockServiceFileDownload = mocker.NewMockCall(callServiceFileDownload)
		mockServiceFileDownload.Return(nil, nil)
	}

	t.Run("should return error when call e-document failed", func(t *testing.T) {
		beforeEachPreviewFile()
		mockServiceFileDownload.Return(nil, errors.New("Error"))

		_, err := service.PreviewFile(int64(1), int64(1))

		expectedError := convert.ValueToErrorCodePointer(esystemError.InternalServerError)
		assert.Equal(t, expectedError, err)
	})

	t.Run("should return struct slice byte when call e-document success", func(t *testing.T) {
		beforeEachPreviewFile()

		resDownload := []byte("Hello World")
		mockServiceFileDownload.Return(resDownload, nil)

		res, err := service.PreviewFile(int64(1), int64(1))

		expected := resDownload
		assert.Equal(t, expected, res)
		assert.Nil(t, err)
	})

}
