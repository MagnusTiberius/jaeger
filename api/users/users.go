package users

import (
	//"fmt"
	//"net/http"
	//"html/template"
	//"bytes"
	//s "database/sql"
	//mysql "api/mysqlmaster"
	//dbi "api/db"
	//"github.com/gorilla/mux"
	"appengine/datastore"
	//"appengine"
	//"github.com/gorilla/sessions"
	
)

type User struct {
	Email  		string
	Password 	string
	UserName	string
	FirstName	string
	LastName	string
	UserID		int
	KeyIdString string
	KeyIdInt	int64
	Key 		*datastore.Key
}


func NewUser() *User {
	v := new(User)
	return v
}

