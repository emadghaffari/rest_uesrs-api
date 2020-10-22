package userdb

import (
	"database/sql"
	"fmt"
	"log"

	// mysql package for connection to mysqlDB
	"github.com/emadghaffari/res_errors/logger"
	"github.com/go-sql-driver/mysql"
)

var (
	// Client connection
	Client *sql.DB
)

func init() {
	var err error
	username := "root"
	password := "password"
	host := "db"
	schema := "golang"
	datasource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)
	Client, err = sql.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	mysql.SetLogger(logger.GetLogger())
	log.Println("database successfuly confiqured.")
}
