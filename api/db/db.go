package db

import (
	//"fmt"
	//"net/http"
	//"html/template"
	//"bytes"
	s "database/sql"
	_ "api/mysqlmaster"
	//"github.com/gorilla/mux"
)

var db *s.DB

func Exec( sql string ) s.Result {
	db, err := s.Open("mysql", "root:konbanwa@tcp(127.0.0.1:3306)/jaeger")
	if err != nil {

	}
	res, err := db.Exec(sql)
	if err != nil {

	}
	return res
}

func Query( sql string ) s.Result {
	db, _ = s.Open("mysql", "root:konbanwa@tcp(127.0.0.1:3306)/jaeger")
	row, _ := db.Exec(sql)
	return row
}