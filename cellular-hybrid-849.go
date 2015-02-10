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
	_ "api/mysqlmaster"
	_ "api/db"
	users "api/users"
	"github.com/gorilla/mux"
)

var (
	templates = template.Must(template.New("").Delims("<<", ">>").ParseFiles(
		"index.html",
		"signup.html",
		"about.html",
		"contact.html",
		"error.html",
	))
)

func init() {
	rtr := mux.NewRouter()
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/signup", handleSignUp)
	http.HandleFunc("/about", handleAbout)
	http.HandleFunc("/contact", handleContact)
	http.HandleFunc("/error", handleError)
	rtr.HandleFunc("/user/{name:[a-z]+}/profile", handleProfile).Methods("GET")
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
	if err := templates.ExecuteTemplate(b, "index.html", nil); err != nil {
		//writeError(w, r, err)
		return
	}
	b.WriteTo(w)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "<html><body>Hello, World! 세상아 안녕!</body></html>")
	
	b := &bytes.Buffer{}
	if err := templates.ExecuteTemplate(b, "index.html", nil); err != nil {
		//writeError(w, r, err)
		return
	}
	b.WriteTo(w)
}

func handleSignUp(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		
		email := r.FormValue("email")
		pwd := r.FormValue("passwd")
		uname := r.FormValue("lastname") 
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
		users.SignUp(u)
		
		b := &bytes.Buffer{}
		if err := templates.ExecuteTemplate(b, "index.html", nil); err != nil {
			panic(err)
			//writeError(w, r, err)
			return
		}
		b.WriteTo(w)
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
