package data

import (
	"database/sql"
	"log"
)

var Db *sql.DB

// 初始化, 打开DB连接
func init() {
	var err error
	Db, err := sql.Open("postgres", "dbname=chitchat sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}
