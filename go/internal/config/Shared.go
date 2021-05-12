package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type GlobalShared struct {
	Psqlconn *sql.DB
}

type PsqlConn struct {
	Host     string
	Port     int64
	User     string
	Password string
	Dbname   string
}

func InitShared(psql PsqlConn) GlobalShared {
	psqlconn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", psql.User, psql.Password, psql.Host, psql.Dbname)
	psqldb, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err.Error())
	}
	return GlobalShared{
		Psqlconn: psqldb,
	}
}
