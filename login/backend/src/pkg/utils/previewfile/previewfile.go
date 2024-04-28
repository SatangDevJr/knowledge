package previewfile

import (
	"esystem/src/pkg/utils/convert"
	esystemError "esystem/src/pkg/utils/error"
	"esystem/src/pkg/utils/file"
	"esystem/src/pkg/utils/logger"
)

var (
	DocumentFolderId string
	DocumentDomainId string
)

type UseCase interface {
	PreviewFile(fileId int64, domainId int64) ([]byte, *esystemError.ErrorCode)
}

type Service struct {
	UseCase
	Logs        logger.Logger
	FileService file.FileService
}

func NewService(serviceParam ServiceParam) *Service {
	service := &Service{
		FileService: serviceParam.FileService,
		Logs:        serviceParam.Logs,
	}
	service.UseCase = service
	return service
}

func (service *Service) PreviewFile(fileId int64, domainId int64) ([]byte, *esystemError.ErrorCode) {
	download := file.Download{
		FolderFileId: fileId,
		DomainId:     domainId,
	}

	res, errE_document := service.FileService.Download(download)
	if errE_document != nil {
		go service.Logs.Error("", "previewFile_SaveDraft_File", "", "Can not Download")
		return nil, convert.ValueToErrorCodePointer(esystemError.InternalServerError)
	}

	return res, nil
}
