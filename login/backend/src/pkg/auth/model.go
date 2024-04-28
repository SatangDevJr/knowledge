package auth

import (
	"esystem/src/pkg/user"
	utilsDecode "esystem/src/pkg/utils/decode"
	"esystem/src/pkg/utils/logger"
)

type ServiceParam struct {
	UtilDecodeService utilsDecode.UseCase
	UserService       user.UseCase
	Repo              Repository
	Logs              logger.Logger
}
