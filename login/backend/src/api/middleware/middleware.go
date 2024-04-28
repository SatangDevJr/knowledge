package middleware

import (
	"encoding/json"
	"esystem/src/api/requestheader"
	"net/http"

	// "esystem/src/pkg/auth"
	ptpspmerror "esystem/src/pkg/utils/error"
	"esystem/src/pkg/utils/logger"
	"esystem/src/pkg/utils/recover"
)

type contextKey string

const (
	JWTClaimsContextKey contextKey = "JWTClaims"
)

type Middleware struct {
	// AuthService auth.UseCase
	Logs logger.Logger
}

func NewMiddleware(
	// authService auth.UseCase,
	logs logger.Logger,
) *Middleware {
	return &Middleware{
		// AuthService: authService,
		Logs: logs,
	}
}

func (m Middleware) Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if recover.Recover(m.Logs) {
				w.Header().Set(requestheader.ContentType, requestheader.ApplicationJson)
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(ptpspmerror.NewError(ptpspmerror.InternalServerError, ptpspmerror.InternalServerErrorMessage))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
