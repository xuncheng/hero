package mysql

import (
	"github.com/gocraft/dbr"
	_ "github.com/go-sql-driver/mysql"
)

var conn *dbr.Connection

func Connect(dsn string) error {
	var err error
	conn, err = dbr.Open("mysql", dsn, nil)

	return err
}

func Conn() *dbr.Connection {
	return conn
}

func Close() error {
	return conn.Close()
}