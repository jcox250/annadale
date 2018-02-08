package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// SQLHandler foo
type SQLHandler struct {
	Conn *sql.DB
}

// NewSQLHandler foo
func NewSQLHandler(dbtype, connStr string) *SQLHandler {
	conn, err := sql.Open(dbtype, connStr)
	if err != nil {
		log.Fatalf("could not connect to db: %s", err)
	}
	sh := new(SQLHandler)
	sh.Conn = conn
	return sh
}

// Query executes a sql statement
func (s *SQLHandler) Query(stmt string) (*sql.Rows, error) {
	rows, err := s.Conn.Query(stmt)
	if err != nil {
		return rows, fmt.Errorf("could not execute %s: %s", stmt, err)
	}
	return rows, nil
}
