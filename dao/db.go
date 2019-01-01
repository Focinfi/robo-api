package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
)

const (
	mysqlDSN = "root@tcp(127.0.0.1:3306)/robo?autocommit=1&timeout=200ms&readTimeout=2s&writeTimeout=2s&parseTime=true"
)

var conn *dbr.Connection

func init() {
	var err error
	conn, err = dbr.Open("mysql", mysqlDSN, nil)
	if err != nil {
		panic("init mysql connection failed: " + err.Error())
	}
}

func newSess() *dbr.Session {
	return conn.NewSession(nil)
}
