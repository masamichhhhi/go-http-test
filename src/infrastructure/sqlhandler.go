package infrastructure

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSQLHandler() *SqlHandler {
	conn, err := sql.Open("mysql", "localuser:localpass@tcp(db:3306)/go_sample")
	if err != nil {
		panic(err.Error)
	}

	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
