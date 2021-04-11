package mysql

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"activity_service/activity"
)

type mysqlRepo struct {
	db    *sql.DB
	table string
}

var (
	NoRowsAffected = errors.New("no_rows_affected")
)

func NewRepo(dsn, table string, pool int) (activity.Repo, error) {
	var (
		db    *sql.DB
		err error
	)

	for i := 0; i < 5; i++ {
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			time.Sleep(2*time.Second)
			continue
		}
		err = db.Ping()
		if err != nil {
			time.Sleep(2*time.Second)
			continue
		}
		err = nil
		break
	}

	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3) // Idle connection time is set to 3 minutes
	db.SetMaxOpenConns(pool)
	db.SetMaxIdleConns(pool)

	return &mysqlRepo{
		db:    db,
		table: table,
	}, nil
}
