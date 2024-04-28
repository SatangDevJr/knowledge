package user

import (
	utils "esystem/src/pkg/utils/encode"
	"esystem/src/pkg/utils/logger"
)

type ServiceParam struct {
	Repo        Repository
	UtilService utils.UseCase
	Logs        logger.Logger
}
