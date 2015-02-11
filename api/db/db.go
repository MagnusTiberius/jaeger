package db

import (
	//"fmt"
	//"net/http"
	//"html/template"
	//"bytes"
	s "database/sql"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/gorilla/mux"
	//"appengine/cloudsql"
)

var db *s.DB

const CONN2 string = "root:konbanwa@tcp(127.0.0.1:3306)/jaeger"

func Exec( sql string ) s.Result {
	db, err := s.Open("mysql", CONN2)
	if err != nil {

	}
	res, err := db.Exec(sql)
	if err != nil {

	}
	return res
}

func Query( sql string ) *s.Rows {
	db, _ = s.Open("mysql", CONN2)
	defer db.Close()
	rows, _ := db.Query(sql)
	_ = rows
	//for rows.Next() {

	//}
	return rows
}