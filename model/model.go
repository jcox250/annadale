package model

import (
	"database/sql"
	"fmt"
	"io"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() (*sql.DB, error) {
	connStr := "root:root@(localhost:3306)/flexbdna"
	conn, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}
	return conn, nil
}

func SafeClose(c io.Closer, err *error) {
	if cerr := c.Close(); cerr != nil && err != nil {
		*err = cerr
	}
}
