package main

import (
	//"fmt"
	"net/http"
	//"html/template"
	//"bytes"
	//"database/sql"
	//_ "github.com/go-sql-driver/mysql"
	//_ "api/db"
	//users "api/users"
	//"github.com/gorilla/mux"
	//"github.com/gorilla/sessions"
	"api/context"
	"api/inv"
)

type WebPage struct {
	UserName 	string
	Email 		string
	NavBar 		string
	Vehicles    []inv.VehicleEntity
	Vehicle     inv.VehicleEntity
}

func GetSessionWebPage(w http.ResponseWriter, r *http.Request, appcontext *context.Context) WebPage {
	var h WebPage
	//appcontext := context.GetContext()
	session, _ := appcontext.Store.Get(r, "jaegersignup")
	//panic(session)
	var name string 
	var email string
	var navbar string 
	if len(session.Values) > 0 {
		name = session.Values["UserName"].(string)
		email = session.Values["Email"].(string)
		navbar = "navbar.html"
		//if len(name) == 0  {
		//	name = "Undefined Name"
		//}
		h = WebPage{UserName:name, Email:email, NavBar:navbar}
	} 
	return h
}