package file

import (
	"esystem/src/pkg/utils/logger"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	DocumentFolderId string
	DocumentDomainId string
	DocumentURL      string
)

type FileService interface {
	Download(download Download) ([]byte, error)
}

type Service struct {
	Service FileService
	Logs    logger.Logger
}

type FileInfoE_document struct {
	FolderFileId int    `json:"folderFileId"`
	FileName     string `json:"fileName"`
}

type Download struct {
	FolderFileId int64 `json:"folderId"`
	DomainId     int64 `json:"domainId"`
}

type ServiceParam struct {
	Logs logger.Logger
}

func NewService(logs logger.Logger) *Service {
	return &Service{
		Logs: logs,
	}
}

func (service *Service) Download(download Download) ([]byte, error) {
	url := fmt.Sprintf(`%[1]s?folderFileId=%[2]d&domainId=%[3]d`, DocumentURL, download.FolderFileId, download.DomainId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	resBody, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		return nil, errRead
	}

	return resBody, nil
}
