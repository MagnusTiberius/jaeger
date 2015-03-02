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
	"appengine/blobstore"
	wsusers "api/ws/wsusers"
	"encoding/json"
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
		"vehicleadmin.html",
		"vehicleview.html",
		"vehicles.html",
		"upload.html",
		"userlist.html",
		"useradminvehicles.html",
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
	rtr.HandleFunc("/user/list", handleUserList).Methods("GET","POST")
	rtr.HandleFunc("/user/{name:[a-z]+}/myaccount", handleProfile).Methods("GET","POST")
	rtr.HandleFunc("/user/{name:[a-z]+}/signout", handleSignout).Methods("GET","POST")
	rtr.HandleFunc("/user/{name:[a-z]+}/vehicles", handleVehicles).Methods("GET","POST")
	rtr.HandleFunc("/user/{name:[a-z]+}/vehicles/admin", handleVehiclesUserAdmin).Methods("GET","POST")
	rtr.HandleFunc("/entity/{name:[a-zA-Z0-9]+}/create", handleInvItm).Methods("GET","POST")
	rtr.HandleFunc("/vehicle/{name:[a-zA-Z0-9]+}/create", handleVehicleCreate).Methods("GET","POST")
	rtr.HandleFunc("/vehicle/{name:[a-zA-Z0-9]+}/edit", handleVehicleEdit).Methods("GET","POST")
	rtr.HandleFunc("/vehicle/{name:[a-zA-Z0-9]+}/{vehicle:[a-zA-Z0-9]+}/admin", handleVehicleAdmin).Methods("GET","POST")
	rtr.HandleFunc("/vehicle/{name:[a-zA-Z0-9]+}/view", handleVehicleView).Methods("GET","POST")
	rtr.HandleFunc("/", handleHome).Methods("GET","POST")
	rtr.HandleFunc("/upload", handleUpload).Methods("GET","POST")
	rtr.HandleFunc("/upload/url", getUploadUrlSession).Methods("GET")
	rtr.HandleFunc("/upload/angular", handleUploadAngular).Methods("GET","POST")
	rtr.HandleFunc("/uploadcomplete", handleUploadComplete).Methods("GET","POST")
	rtr.HandleFunc("/uploadcomplete/angular", handleUploadCompleteAngular).Methods("GET","POST")
	rtr.HandleFunc("/blob/", handleBlob).Methods("GET","POST")

	rtr.HandleFunc("/ws/user/list", handleWsUserList).Methods("GET","POST")
	rtr.HandleFunc("/ws/user/{name:[a-zA-Z0-9]+}/vehicles/getall", handleWsVehiclesUserAdminGetall).Methods("GET","POST")
	rtr.HandleFunc("/ws/vehicle/{vehicle:[a-zA-Z0-9]+}/carousel/getall", handleWsCarouselGetall).Methods("GET","POST")

	http.Handle("/", rtr)

}


func handleWsUserList(w http.ResponseWriter, r *http.Request) {
	list := wsusers.HandleWsUserList(w,r)
	js, err := json.Marshal(list)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func handleUserList(w http.ResponseWriter, r *http.Request) {
	h := main.GetSessionWebPage(w,r,appcontext)

	b := &bytes.Buffer{}
	if err := templates.ExecuteTemplate(b, "userlist.html", h); err != nil {
		http.Redirect(w,r,"/error",301)
		//writeError(w, r, err)
		return
	}
	b.WriteTo(w)
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

    uploadURL, err := blobstore.UploadURL(c, "/uploadcomplete", nil)
    if err != nil {
            panic(err)
            return
    }

	b := &bytes.Buffer{}
	h := main.GetSessionWebPage(w,r, appcontext)
	h.UploadURL = uploadURL
	h.RequestURI = uploadURL.RequestURI()
	if err := templates.ExecuteTemplate(b, "upload.html", h); err != nil {
		//writeError(w, r, err)
		return
	}
	b.WriteTo(w)
}

func handleUploadAngular(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

    uploadURL, err := blobstore.UploadURL(c, "/uploadcomplete", nil)
    if err != nil {
            panic(err)
            return
    }

	b := &bytes.Buffer{}
	h := main.GetSessionWebPage(w,r, appcontext)
	h.UploadURL = uploadURL
	h.RequestURI = uploadURL.RequestURI()
	if err := templates.ExecuteTemplate(b, "upload.html", h); err != nil {
		//writeError(w, r, err)
		return
	}
	b.WriteTo(w)
}


func handleUploadComplete(w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)

    blobs, _, err := blobstore.ParseUpload(r)
    if err != nil {
            panic(err)
            return
    }
    file := blobs["file"]
    if len(file) == 0 {
    	panic("file invalid")
    }
    http.Redirect(w, r, "/blob/?blobKey="+string(file[0].BlobKey), http.StatusFound)
}

func handleUploadCompleteAngular(w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)

    blobs, _, err := blobstore.ParseUpload(r)
    _ = blobs
    if err != nil {
            panic(err)
            return
    }
    file := blobs["file"]
    if len(file) == 0 {
    	panic("file invalid")
    }
    //http.Redirect(w, r, "/blob/?blobKey="+string(file[0].BlobKey), http.StatusFound)
    key := fmt.Sprintf("{\"blobKey\":\"%s\"}",string(file[0].BlobKey))
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(key))
}

func handleBlob(w http.ResponseWriter, r *http.Request) {
	blobstore.Send(w, appengine.BlobKey(r.FormValue("blobKey")))
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


func handleWsCarouselGetall(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	params := mux.Vars(r)
	vehicleKey := params["vehicle"]
	_ = vehicleKey
	session, _ := appcontext.Store.Get(r, "jaegersignup")
	keyIdString := session.Values["KeyIdString"].(string) 

	q := datastore.NewQuery("Vehicle").
                Filter("KeyName =", vehicleKey)
    var vehicles []inv.VehicleEntity
    keysV, err := q.GetAll(c, &vehicles)
    _ = keysV
    if err != nil {
            panic(err)
            return
    }

    if len(vehicles) == 0 {
    	panic("Vehicle not found")
    	return
    }

    //vehicle := vehicles[0]

	ancestorKey := datastore.NewKey(c, "Vehicle", keyIdString, 0, nil)	
	q = datastore.NewQuery("Carousel").Ancestor(ancestorKey)
	var carouselItems []inv.CarouselEntity
	keysCar, err := q.GetAll(c, &carouselItems)
	_ = keysCar
	js, err := json.Marshal(carouselItems)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)	
}

func handleWsVehiclesUserAdminGetall(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	params := mux.Vars(r)
	username := params["name"]
	_ = username
	session, _ := appcontext.Store.Get(r, "jaegersignup")
	keyIdString := session.Values["KeyIdString"].(string) 

	q := datastore.NewQuery("User").
                Filter("UserName =", username)
    var users []users.User
    _, err := q.GetAll(c, &users)
    if err != nil {
            panic(err)
            return
    }
	ancestorKey := datastore.NewKey(c, "User", keyIdString, 0, nil)	
	q = datastore.NewQuery("Vehicle").Ancestor(ancestorKey)
	var vehicles []inv.VehicleEntity
	_, err = q.GetAll(c, &vehicles)
	js, err := json.Marshal(vehicles)
	if err != nil {
		panic(err)
	}
	if (vehicles == nil) {
		v := new(inv.VehicleEntity)
		v.ManufacturerCode = "undefined"
		v.ModelCode = "undefined"
		v.TrimCode = "undefined"
		v.KeyName = context.RandSeq(32)
		v.KeyId = v.KeyName
		v2 := new(inv.VehicleEntity)
		v2.ManufacturerCode = "undefined"
		v2.ModelCode = "undefined"
		v2.TrimCode = "undefined"
		v2.KeyName = context.RandSeq(32)
		v2.KeyId = v2.KeyName
		vehicles = []inv.VehicleEntity{*v, *v2}
		js, err = json.Marshal(vehicles)
		if err != nil {
			panic(err)
		}
	} 
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)	
}


func handleVehiclesUserAdmin(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)
	//session, _ := appcontext.Store.Get(r, "jaegersignup")
	//appcontext := context.GetContext()

	//panic(userkey)
	session, _ := appcontext.Store.Get(r, "jaegersignup")
	email := session.Values["Email"].(string)
	k := datastore.NewKey(c, "User", email, 0, nil)
	//panic(k)
	_=k

    //panic(appcontext)

	b := &bytes.Buffer{}
	h := main.GetSessionWebPage(w,r,appcontext)

	//panic(h)
	if err := templates.ExecuteTemplate(b, "useradminvehicles.html", h); err != nil {
		http.Redirect(w,r,"/error",301)
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
		http.Redirect(w,r,"/error",301)
		//writeError(w, r, err)
		return
	}
	b.WriteTo(w)

}



func handleVehicleView(w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)
	params := mux.Vars(r)
	name := params["name"]
	_ = name
	b := &bytes.Buffer{}
	h := main.GetSessionWebPage(w,r,appcontext)
	vehicles := inv.GetVehiclesByKey(r,appcontext,name)
	if len(vehicles) > 0 {
		h.Vehicle = vehicles[0]
		if err := templates.ExecuteTemplate(b, "vehicleview.html", h); err != nil {
			http.Redirect(w,r,"/error",301)
			//writeError(w, r, err)
			panic(err)
			return
		}
		b.WriteTo(w)
	} else {
		http.Redirect(w,r,"/error",301)
	}

}

func getUploadUrlSession(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
    uploadURL, err := blobstore.UploadURL(c, "/uploadcomplete/angular", nil)
    if err != nil {
            panic(err)
            return
    }
	sessn := uploadURL.RequestURI()
    jsn := fmt.Sprintf("{\"uploadurl\":\"%s\"}",sessn)
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(jsn))

}

func handleVehicleAdmin(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	_ = c
	params := mux.Vars(r)
	name := params["name"]
	_ = name
	vehicleKey := params["vehicle"]
	_ = vehicleKey

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
		medit := fmt.Sprintf("/vehicle/%v/admin",name)
		http.Redirect(w,r,medit,301)
		return
	}

	b := &bytes.Buffer{}
	h := main.GetSessionWebPage(w,r, appcontext)

	if (len(h.Email) == 0) {
		panic("Invalid session, no email found in session.")
		return
	}

	q := datastore.NewQuery("VehicleEntity").
                Filter("KeyName =", vehicleKey)
    var vehicles []inv.VehicleEntity
    _, err := q.GetAll(c, &vehicles)
    if err != nil {
            panic(err)
            return
    }

	session, _ := appcontext.Store.Get(r, "jaegersignup")
	keyIdString := session.Values["KeyIdString"].(string)
	userKey := datastore.NewKey(c, "User", keyIdString, 0, nil)	
	apengcontext := appengine.NewContext(r)

    if len(vehicles) == 0 {
		key := datastore.NewKey(apengcontext, "Vehicle", vehicleKey, 0, userKey)
		entity := new(inv.VehicleEntity)
		entity.ManufacturerCode = "undefined"
		entity.ModelCode = "undefined"
		entity.TrimCode = "undefined"
		entity.KeyName = vehicleKey
		entity.KeyId = vehicleKey
		keyVehicle, err := datastore.Put(apengcontext, key, entity)
		_ = keyVehicle
		q := datastore.NewQuery("VehicleEntity").
        		Filter("KeyName =", vehicleKey)
	    var vehicles []inv.VehicleEntity
	    _, err = q.GetAll(c, &vehicles)		

	    if err != nil {
	            panic(err)
	            return
	    }

    }

    vk := datastore.NewKey(apengcontext, "Vehicle", vehicleKey, 0, userKey)
    q = datastore.NewQuery("CarouselEntity").Ancestor(vk)
    carouselEntityList := []inv.CarouselEntity{}
    _, err = q.GetAll(c, &carouselEntityList)		
    if err != nil {
            panic(err)
            return
    }

    if len(carouselEntityList) == 0 {
    	itm := new(inv.CarouselEntity)
    	itm.KeyId = context.RandSeq(32)
    	key := datastore.NewKey(apengcontext, "CarouselEntity", itm.KeyId, 0, vk)
    	_, err := datastore.Put(apengcontext, key, itm)
	    if err != nil {
	            panic(err)
	            return
	    }
    }

    h.VehicleKey = vehicleKey
	if err := templates.ExecuteTemplate(b, "vehicleadmin.html", h); err != nil {
		//writeError(w, r, err)
		panic(err)
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
		u.KeyIdString = context.RandSeq(32)
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
