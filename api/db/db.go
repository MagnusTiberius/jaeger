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

const CONN1 string = "user1:konbanwa(173.194.249.186:3306)/jaegerterrace"
const CONN2 string = "root:konbanwa@tcp(127.0.0.1:3306)/jaeger"
const CONN3 string = "cloudsql:cellular-hybrid-849:jaeger*jaegerterrace/user1/konbanwa"

func Exec( sql string ) s.Result {
	db, err := s.Open("mysql", CONN1)
	if err != nil {

	}
	res, err := db.Exec(sql)
	if err != nil {

	}
	return res
}

func Query( sql string ) *s.Rows {
	db, _ = s.Open("mysql", CONN1)
	defer db.Close()
	rows, _ := db.Query(sql)
	_ = rows
	//for rows.Next() {

	//}
	return rows
}