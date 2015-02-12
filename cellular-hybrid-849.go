// Copyright 2011 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package cellularhybrid849

import (
	//"fmt"
	"net/http"
	"html/template"
	"bytes"
	//"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "api/db"
	users "api/users"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	templates = template.Must(template.New("").Delims("[[", "]]").ParseFiles(
		"index.html",
		"signup.html",
		"about.html",
		"contact.html",
		"error.html",
		"template1/signupgood.html",
		"template1/signupinvalid.html",
		"myaccount.html",
	))
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func init() {
	rtr := mux.NewRouter()
	//http.HandleFunc("/", handleHome)
	http.HandleFunc("/signup", handleSignUp)
	http.HandleFunc("/signin", handleSignIn)
	http.HandleFunc("/signup/good", handleSignUpGood)
	http.HandleFunc("/about", handleAbout)
	http.HandleFunc("/contact", handleContact)
	http.HandleFunc("/error", handleError)
	rtr.HandleFunc("/user/{name:[a-z]+}/myaccount", handleProfile).Methods("GET","POST")
	rtr.HandleFunc("/", handleHome).Methods("GET","POST")
	http.Handle("/", rtr)
	/*
	indexTmpl := template.New("").Delims("<<", ">>")
	indexTmpl,_ = indexTmpl.ParseFiles(
			"index.html",
			"signup.html",
			"about.html",
			"contact.html",
			"error.html",
		)	
	templates = template.Must(indexTmpl.ParseFiles(
			"index.html",
			"signup.html",
			"about.html",
			"contact.html",
			"error.html",
		))
		*/
}

func handleProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	_ = name
	b := &bytes.Buffer{}
	h := GetSessionWebPage(w,r)
	if err := templates.ExecuteTemplate(b, "myaccount.html", h); err != nil {
		//writeError(w, r, err)
		return
	}
	b.WriteTo(w)
}

type WebPage struct {
	UserName 	string
	Email 		string
}

func GetSessionWebPage(w http.ResponseWriter, r *http.Request) WebPage {
	var h WebPage
	session, _ := store.Get(r, "jaegersignup")
	var name string 
	var email string 
	if len(session.Values) > 0 {
		name = session.Values["UserName"].(string)
		email = session.Values["Email"].(string)
		//if len(name) == 0  {
		//	name = "Undefined Name"
		//}
		h = WebPage{UserName:name, Email:email}
	} 
	return h
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	h := GetSessionWebPage(w,r)
	b := &bytes.Buffer{}
	if err := templates.ExecuteTemplate(b, "index.html", h); err != nil {
		//writeError(w, r, err)
		return
	}
	b.WriteTo(w)
}


func handleSignUpGood(w http.ResponseWriter, r *http.Request) {
	b := &bytes.Buffer{}
	if err := templates.ExecuteTemplate(b, "template1/signupgood.html", nil); err != nil {
		//writeError(w, r, err)
		return
	}
	b.WriteTo(w)

}



func handleSignIn(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		email := r.FormValue("email")
		pwd := r.FormValue("passwd")
		if len(email) == 0 {
			panic("email is empty")
		}
		if len(pwd) == 0 {
			panic("pwd is empty")
		}
		u := users.NewUser()
		u.Email = email
		u.Password = pwd

		ok, usr := users.SignIn(u, r)

		if ok {

			session, _ := store.Get(r, "jaegersignup")
			session.Values["Email"] = usr.Email
			session.Values["UserName"] = usr.UserName
			session.Save(r, w)
			http.Redirect(w,r,"/",301)
		}
	}
}

func handleSignUp(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		
		email := r.FormValue("email")
		pwd := r.FormValue("passwd")
		uname := r.FormValue("username") 

		if len(email) == 0 {
			panic("email is empty")
		}
		if len(pwd) == 0 {
			panic("pwd is empty")
		}
		if len(uname) == 0 {
			panic("uname is empty")
		}
		u := users.NewUser()
		u.Email = email
		u.Password = pwd
		u.UserName = uname
		u.FirstName = r.FormValue("firstname") 
		u.LastName = r.FormValue("lastname") 
		if users.SignUp(u, r) {
			b := &bytes.Buffer{}
			if err := templates.ExecuteTemplate(b, "signupgood.html", nil); err != nil {
				panic(err)
				//writeError(w, r, err)
				return
			}
			b.WriteTo(w)
		} else {
			b := &bytes.Buffer{}
			if err := templates.ExecuteTemplate(b, "signupinvalid.html", nil); err != nil {
				panic(err)
				//writeError(w, r, err)
				return
			}
			b.WriteTo(w)
		}
	}

	if r.Method == "GET" {
		b := &bytes.Buffer{}
		if err := templates.ExecuteTemplate(b, "signup.html", nil); err != nil {
			//writeError(w, r, err)
			return
		}
		b.WriteTo(w)
	}
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "<html><body>Hello, World! 세상아 안녕!</body></html>")
	
	b := &bytes.Buffer{}
	if err := templates.ExecuteTemplate(b, "about.html", nil); err != nil {
		//writeError(w, r, err)
		return
	}
	b.WriteTo(w)
}

func handleContact(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "<html><body>Hello, World! 세상아 안녕!</body></html>")
	
	b := &bytes.Buffer{}
	if err := templates.ExecuteTemplate(b, "contact.html", nil); err != nil {
		//writeError(w, r, err)
		return
	}
	b.WriteTo(w)
}

func handleError(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "<html><body>Hello, World! 세상아 안녕!</body></html>")
	
	b := &bytes.Buffer{}
	if err := templates.ExecuteTemplate(b, "error.html", nil); err != nil {
		//writeError(w, r, err)
		return
	}
	b.WriteTo(w)
}
