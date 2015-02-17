// Copyright 2011 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package cellularhybrid849

import (
	"fmt"
	"net/http"
	"html/template"
	"bytes"
	//"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "api/db"
	users "api/users"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"api/context"
	"api/inv"
 	"api/main"
	"appengine/datastore"
	"appengine"	 	
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
		"invitm.html",
		"navbar.html",
		"maindochdr.html",
		"maindocftr.html",
		"vehiclecreate.html",
		"vehicleedit.html",
		"vehicles.html",
	))
)

var store *sessions.CookieStore //= sessions.NewCookieStore([]byte("something-very-secret"))

var appcontext *context.Context
var userkey *datastore.Key

func init() {
	appcontext = context.NewContext()

	rtr := mux.NewRouter()
	//http.HandleFunc("/", handleHome)
	http.HandleFunc("/signup", handleSignUp)
	http.HandleFunc("/signin", handleSignIn)
	http.HandleFunc("/signup/good", handleSignUpGood)
	http.HandleFunc("/about", handleAbout)
	http.HandleFunc("/contact", handleContact)
	rtr.HandleFunc("/error", handleError).Methods("GET","POST")
	rtr.HandleFunc("/user/{name:[a-z]+}/myaccount", handleProfile).Methods("GET","POST")
	rtr.HandleFunc("/user/{name:[a-z]+}/signout", handleSignout).Methods("GET","POST")
	rtr.HandleFunc("/user/{name:[a-z]+}/vehicles", handleVehicles).Methods("GET","POST")
	rtr.HandleFunc("/entity/{name:[a-z]+}/create", handleInvItm).Methods("GET","POST")
	rtr.HandleFunc("/vehicle/{name:[a-z]+}/create", handleVehicleCreate).Methods("GET","POST")
	rtr.HandleFunc("/vehicle/{name:[a-z]+}/edit", handleVehicleEdit).Methods("GET","POST")
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
	h := main.GetSessionWebPage(w,r, appcontext)
	if err := templates.ExecuteTemplate(b, "myaccount.html", h); err != nil {
		//writeError(w, r, err)
		return
	}
	b.WriteTo(w)
}



func handleVehicles(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	//session, _ := appcontext.Store.Get(r, "jaegersignup")
	//appcontext := context.GetContext()

	//panic(userkey)
	session, _ := appcontext.Store.Get(r, "jaegersignup")
	email := session.Values["Email"].(string)
	k := datastore.NewKey(c, "User", email, 0, nil)
	//panic(k)
	_=k
    q := datastore.NewQuery("Vehicle").Ancestor(k)

    //panic(appcontext)

    var vehicles []inv.VehicleEntity
    _, err := q.GetAll(c, &vehicles)
    //panic(vehicles)
    if err != nil {
    	panic(err)
    }

	b := &bytes.Buffer{}
	h := main.GetSessionWebPage(w,r,appcontext)
	h.Vehicles = vehicles
	//panic(h)
	if err := templates.ExecuteTemplate(b, "vehicles.html", h); err != nil {
		//writeError(w, r, err)
		return
	}
	b.WriteTo(w)

}

func handleVehicleEdit(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	_ = name

	if r.Method == "POST" {
		/*
		b := &bytes.Buffer{}
		h := GetSessionWebPage(w,r,appcontext)
		if err := templates.ExecuteTemplate(b, "vehicleedit.html", h); err != nil {
			//writeError(w, r, err)
			return
		}
		b.WriteTo(w)
		*/
		medit := fmt.Sprintf("/vehicle/%v/edit",name)
		http.Redirect(w,r,medit,301)
		return
	}

	b := &bytes.Buffer{}
	h := main.GetSessionWebPage(w,r,appcontext)
	if err := templates.ExecuteTemplate(b, "vehicleedit.html", h); err != nil {
		//writeError(w, r, err)
		return
	}
	b.WriteTo(w)
}

func handleVehicleCreate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	_ = name

	if r.Method == "POST" {
		/*
		b := &bytes.Buffer{}
		h := GetSessionWebPage(w,r,appcontext)
		if err := templates.ExecuteTemplate(b, "vehicleedit.html", h); err != nil {
			//writeError(w, r, err)
			return
		}
		b.WriteTo(w)
		*/

		key, ok := inv.AddVehicleEntity(r,appcontext)
		_=key
		if !ok {
			http.Redirect(w,r,"/error",301)
		}

		medit := fmt.Sprintf("/vehicle/%v/edit",name)
		http.Redirect(w,r,medit,301)
		return
	}

	b := &bytes.Buffer{}
	h := main.GetSessionWebPage(w,r,appcontext)
	if err := templates.ExecuteTemplate(b, "vehiclecreate.html", h); err != nil {
		//writeError(w, r, err)
		return
	}
	b.WriteTo(w)
}


func handleInvItm(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	_ = name
	b := &bytes.Buffer{}
	h := main.GetSessionWebPage(w,r,appcontext)
	if err := templates.ExecuteTemplate(b, "invitm.html", h); err != nil {
		//writeError(w, r, err)
		return
	}
	b.WriteTo(w)
}




func handleHome(w http.ResponseWriter, r *http.Request) {
	h := main.GetSessionWebPage(w,r,appcontext)

	vehicles := inv.GetFfeaturedVehicles(r,appcontext)
	h.Vehicles = vehicles

	b := &bytes.Buffer{}
	if err := templates.ExecuteTemplate(b, "index.html", h); err != nil {
		http.Redirect(w,r,"/error",301)
		//writeError(w, r, err)
		return
	}
	b.WriteTo(w)
}


func handleSignUpGood(w http.ResponseWriter, r *http.Request) {
	h := main.GetSessionWebPage(w,r,appcontext)
	b := &bytes.Buffer{}
	if err := templates.ExecuteTemplate(b, "template1/signupgood.html", h); err != nil {
		http.Redirect(w,r,"/error",301)
		//writeError(w, r, err)
		return
	}
	b.WriteTo(w)

}


func handleSignout(w http.ResponseWriter, r *http.Request) {
	session, _ := appcontext.Store.Get(r, "jaegersignup")
	session.Values["Email"] = ""
	session.Values["UserName"] = ""
	session.Save(r, w)
	http.Redirect(w,r,"/",301)
}

func handleSignIn(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("handleSignIn \n")

	if r.Method == "POST" {
		email := r.FormValue("email")
		pwd := r.FormValue("passwd")
		if len(email) == 0 {
			http.Redirect(w,r,"/error",301)
			//panic("email is empty")
			return
		}
		if len(pwd) == 0 {
			http.Redirect(w,r,"/error",301)
			//panic("pwd is empty")
			return
		}
		u := users.NewUser()
		u.Email = email
		u.Password = pwd

		ok, usr, uk := users.SignIn(u, w, r,appcontext)
		_ = usr
		userkey = uk
		//panic(userkey)
		if ok {

			//session, _ := appcontext.Store.Get(r, "jaegersignup")
			//session.Values["Email"] = usr.Email
			//session.Values["UserName"] = usr.UserName
			//session.Save(r, w)
			http.Redirect(w,r,"/",301)
		}
		http.Redirect(w,r,"/error",301)
	}
}

func handleSignUp(w http.ResponseWriter, r *http.Request) {
	h := main.GetSessionWebPage(w,r,appcontext)
	if r.Method == "POST" {
		
		email := r.FormValue("email")
		pwd := r.FormValue("passwd")
		uname := r.FormValue("username") 

		if len(email) == 0 {
			http.Redirect(w,r,"/error",301)
			return
			//panic("email is empty")
		}
		if len(pwd) == 0 {
			http.Redirect(w,r,"/error",301)
			return
			//panic("pwd is empty")
		}
		if len(uname) == 0 {
			http.Redirect(w,r,"/error",301)
			return
			//panic("uname is empty")
		}
		u := users.NewUser()
		u.Email = email
		u.Password = pwd
		u.UserName = uname
		u.FirstName = r.FormValue("firstname") 
		u.LastName = r.FormValue("lastname") 
		if users.SignUp(u, r) {
			b := &bytes.Buffer{}
			if err := templates.ExecuteTemplate(b, "signupgood.html", h); err != nil {
				http.Redirect(w,r,"/error",301)
				panic(err)
				//writeError(w, r, err)
				return
			}
			b.WriteTo(w)
		} else {
			b := &bytes.Buffer{}
			if err := templates.ExecuteTemplate(b, "signupinvalid.html", h); err != nil {
				http.Redirect(w,r,"/error",301)
				panic(err)
				//writeError(w, r, err)
				return
			}
			b.WriteTo(w)
		}
	}

	if r.Method == "GET" {
		b := &bytes.Buffer{}
		if err := templates.ExecuteTemplate(b, "signup.html", h); err != nil {
			http.Redirect(w,r,"/error",301)
			//writeError(w, r, err)
			return
		}
		b.WriteTo(w)
	}
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	h := main.GetSessionWebPage(w,r,appcontext)
	
	b := &bytes.Buffer{}
	if err := templates.ExecuteTemplate(b, "about.html", h); err != nil {
		http.Redirect(w,r,"/error",301)
		//writeError(w, r, err)
		return
	}
	b.WriteTo(w)
}

func handleContact(w http.ResponseWriter, r *http.Request) {
	h := main.GetSessionWebPage(w,r,appcontext)
	
	b := &bytes.Buffer{}
	if err := templates.ExecuteTemplate(b, "contact.html", h); err != nil {
		http.Redirect(w,r,"/error",301)
		//writeError(w, r, err)
		return
	}
	b.WriteTo(w)
}

func handleError(w http.ResponseWriter, r *http.Request) {
	h := main.GetSessionWebPage(w,r,appcontext)
	
	b := &bytes.Buffer{}
	if err := templates.ExecuteTemplate(b, "error.html", h); err != nil {
		//http.Redirect(w,r,"/error",301)
		//writeError(w, r, err)
		return
	}
	b.WriteTo(w)
}
