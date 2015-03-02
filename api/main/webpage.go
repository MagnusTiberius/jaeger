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
	//"appengine/blobstore"
	//"appengine/urlfetch"
	//"appengine"
	"net/url"
)

type WebPage struct {
	UserName 	string
	Email 		string
	NavBar 		string
	KeyIdString string
	Vehicles    []inv.VehicleEntity
	Vehicle     inv.VehicleEntity
	UploadURL   *url.URL
	RequestURI  string
}

func GetSessionWebPage(w http.ResponseWriter, r *http.Request, appcontext *context.Context) WebPage {
	var h WebPage
	//appcontext := context.GetContext()
	session, _ := appcontext.Store.Get(r, "jaegersignup")
	//panic(session)
	var name string 
	var email string
	var navbar string 
	var keyIdString string 
	if len(session.Values) > 0 {
		name = session.Values["UserName"].(string)
		email = session.Values["Email"].(string)
		navbar = "navbar.html"
		keyIdString = session.Values["KeyIdString"].(string) 
		//if len(name) == 0  {
		//	name = "Undefined Name"
		//}
		h = WebPage{UserName:name, Email:email, NavBar:navbar, KeyIdString:keyIdString}
	} 
	return h
}