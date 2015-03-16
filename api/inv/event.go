package inv

import (
	//"fmt"
	"net/http"
	//"html/template"
	//"bytes"
	//s "database/sql"
	//mysql "api/mysqlmaster"
	//dbi "api/db"
	//"github.com/gorilla/mux"
	//"github.com/gorilla/sessions"
	//"github.com/gorilla/sessions"
	"api/context"
	//"time"
	"appengine/datastore"
	"appengine"	
)

type Event struct {
	EventName 				string
	EventDescription        string
	Key 					string
}

func NewEvent() *Event {
	var e = new(Event)
	return e
}

func AddEvent(r *http.Request, appcontext *context.Context) (*datastore.Key, bool) {
	c := appengine.NewContext(r)

	eventName := r.FormValue("EventName")
	eventDescription := r.FormValue("EventDescription")

	if len(eventName) == 0 || len(eventDescription) == 0 {
		//panic("manu model or trim cannot be empty")
		return nil, false
	}

	//appcontext := context.GetContext()
	//session, _ := appcontext.Store.Get(r, "jaegersignup")
	//panic(session.Values)
	//userKey := appcontext.UserKey

	session, _ := appcontext.Store.Get(r, "jaegersignup")
	//email := session.Values["Email"].(string)
	keyIdString := session.Values["KeyIdString"].(string)
	
	userKey := datastore.NewKey(c, "User", keyIdString, 0, nil)	

	apengcontext := appengine.NewContext(r)

	rndKey := context.RandSeq(32)

	k := datastore.NewKey(apengcontext, "Event", rndKey, 0, userKey)
	e := NewEvent()
	e.EventName = eventName
	e.EventDescription = eventDescription
	e.Key = rndKey


	keyEv, err := datastore.Put(apengcontext, k, e)
	if err != nil {
		//panic(err)
        //http.Error(w, err.Error(), http.StatusInternalServerError)
        return nil, false
    }	
	return keyEv, true
}