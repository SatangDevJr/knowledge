package previewfile

import (
	"esystem/src/pkg/utils/file"
	"esystem/src/pkg/utils/logger"
)

type ServiceParam struct {
	Logs        logger.Logger
	FileService file.FileService
}
