package user

import (
	"context"
	"database/sql"
	"esystem/src/pkg/entity"
	"esystem/src/pkg/utils/logger"
	sqlQuery "esystem/src/pkg/utils/sqlquery"
	"fmt"

	esystemError "esystem/src/pkg/utils/error"

	sqlStruct "github.com/kisielk/sqlstruct"
)

type Repository interface {
	UpsertUser(user entity.User) error
	FindUserByUserName(username string) (*entity.User, error)
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

func (repo *SqlRepository) UpsertUser(user entity.User) error {
	ctx := context.Background()
	session, sessionErr := repo.Session.Conn(ctx)
	if sessionErr != nil {
		return sessionErr
	}
	defer session.Close()

	sql := fmt.Sprintf(`
	INSERT INTO %[1]s
	(
		%[2]s
	)
	VALUES
	(
		%[3]s
	)
	`,
		repo.Collection,
		sqlQuery.GenerateQueryColumnNames(entity.User{}, []string{"ID"}),
		sqlQuery.GenerateQueryColumnValues(user, []string{"ID"}),
	)

	_, err := session.ExecContext(ctx, sql)
	if err != nil {
		go repo.Logs.Error("", "User_Repo_Insert", user,
			esystemError.NewError(esystemError.TechnicalError, err.Error()))
		return err
	}

	return nil
}

func (repo *SqlRepository) FindUserByUserName(username string) (*entity.User, error) {
	ctx := context.Background()
	session, sessionErr := repo.Session.Conn(ctx)
	if sessionErr != nil {
		fmt.Println("sessionErr", sessionErr)
		return nil, sessionErr
	}
	defer session.Close()

	sql := fmt.Sprintf(`
		SELECT %[1]s 
		FROM %[2]s 
		WHERE Delflag = 0 AND UserName = N'%[3]s'	`,
		sqlQuery.GenerateQueryColumnNames(entity.User{}, []string{}),
		repo.Collection,
		username,
	)

	rows, err := session.QueryContext(ctx, sql)
	message := fmt.Sprintf("{Username: %[1]v}", username)

	if err != nil {
		go repo.Logs.Error("", "User_Repo_FindUserByUserName", message,
			esystemError.NewError(esystemError.TechnicalError, err.Error()))
		return nil, err
	}

	var entity entity.User
	for rows.Next() {
		if err := sqlStruct.Scan(&entity, rows); err != nil {
			go repo.Logs.Error("", "User_Repo_FindUserByUserName_Select", message, esystemError.NewError(esystemError.TechnicalError, err.Error()))
			return nil, err
		}
	}

	return &entity, nil
}
