package utils

import (
	"esystem/src/pkg/entity"
	"esystem/src/pkg/utils/logger"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Service UseCase
	Logs    logger.Logger
}

type UseCase interface {
	GenerateFromPassword(pwd []byte) (string, error)
	ValidatePassword(pwdHash []byte, pwd []byte) error
	GenerateJWT(user *entity.User) (string, error)
}

type ServiceParam struct {
	Logs logger.Logger
}

func NewService(logs logger.Logger) *Service {
	return &Service{
		Logs: logs,
	}
}

func (service *Service) GenerateFromPassword(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (service *Service) ValidatePassword(pwdHash []byte, pwd []byte) error {
	err := bcrypt.CompareHashAndPassword(pwdHash, pwd)
	if err != nil {
		return err
	}

	return nil
}

func (service *Service) GenerateJWT(user *entity.User) (string, error) {
	var mySigningKey = []byte("Testing2022")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = user.UserName
	claims["userid"] = user.ID
	claims["exp"] = time.Now().Add(time.Minute * 300).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
