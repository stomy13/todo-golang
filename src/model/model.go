package model

import (
	"database/sql"
	"log"
	"os"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// DBConnect returns *sql.DB
func DBConnect() (db *sql.DB) {
	dbDriver := "mysql"
    dbUser := "todo"
	dbPass := os.Getenv("TODO_DATABASE_PASSWORD") // 環境変数から取得
	dbHost := "tcp(db:3306)" // dockerのサービス名を指定
	dbName := "todo"
    dbOption := "?parseTime=true"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbHost+"/"+dbName+dbOption)
    if err != nil {
        log.Fatal(err)
    }
    return db
}