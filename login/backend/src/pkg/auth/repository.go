package auth

import (
	"database/sql"
	"esystem/src/pkg/utils/logger"
)

type Repository interface {
}

type SqlRepository struct {
	Collection string
	Session    *sql.DB
	Logs       logger.Logger
}

func NewRepository(collection string, session *sql.DB, logs logger.Logger) *SqlRepository {
	return &SqlRepository{
		Collection: collection,
		Session:    session,
		Logs:       logs,
	}
}
