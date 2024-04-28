package handler

import (
	"esystem/src/pkg/user"
	"esystem/src/pkg/utils/logger"
)

type HandlerParam struct {
	Service user.UseCase
	Logs    logger.Logger
}
