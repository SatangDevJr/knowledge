package handler

import (
	"esystem/src/pkg/auth"
	"esystem/src/pkg/utils/logger"
)

type HandlerParam struct {
	Service auth.UseCase
	Logs    logger.Logger
}
