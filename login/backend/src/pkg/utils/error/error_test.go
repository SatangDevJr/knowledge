package error_test

import (
	esystem "esystem/src/pkg/utils/error"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewError(t *testing.T) {

	t.Run("should return error struct when new error", func(t *testing.T) {
		err := esystem.NewError(esystem.InternalServerError, esystem.InternalServerErrorMessage)

		expectedError := &esystem.Error{Code: esystem.InternalServerError, Message: esystem.InternalServerErrorMessage}
		assert.Equal(t, expectedError, err)
	})

	t.Run("should return error with data when error has data", func(t *testing.T) {
		err := esystem.NewError(esystem.InternalServerError, esystem.InternalServerErrorMessage, "some data")

		expectedError := &esystem.Error{
			Code:    esystem.InternalServerError,
			Message: esystem.InternalServerErrorMessage,
			Data:    "some data",
		}
		assert.Equal(t, expectedError, err)
	})

}
