package repository

import "database/sql"

type DBHandler interface {
	Query(string) (*sql.Rows, error)
	Execute(string, ...interface{}) (int64, error)
}
