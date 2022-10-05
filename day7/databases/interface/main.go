package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	sql.Register("sqlite3", &SQLiteDruver{})
}

//https://github.com/mattn/go-sqlite3 driver
//var d = Driver{proto: "mymysql", raddr: "127.0.0.1:3306"}
//func init() {
//	Register("SET NAMES utf8")
//	sql.Register("mymysql", &d)
//}