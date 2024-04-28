package decode

import (
	"errors"
	esystemError "esystem/src/pkg/utils/error"
	"esystem/src/pkg/utils/logger"

	"github.com/golang-jwt/jwt"
)

type Service struct {
	Service UseCase
	Logs    logger.Logger
}

type UseCase interface {
	DecodeJWT(accessToken string) (interface{}, error)
}

type ServiceParam struct {
	Logs logger.Logger
}

func NewService(logs logger.Logger) *Service {
	return &Service{
		Logs: logs,
	}
}

func (service *Service) DecodeJWT(accessToken string) (interface{}, error) {

	atClaims := jwt.MapClaims{}

	jwt.ParseWithClaims(accessToken, atClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})

	if atClaims == nil {
		return "", errors.New(string(esystemError.EmptyToken))
	}

	return atClaims["sub"], nil
}
