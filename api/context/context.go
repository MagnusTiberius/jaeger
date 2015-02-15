package context

import (
	//"fmt"
	//"net/http"
	//"html/template"
	//"bytes"
	//"database/sql"
	//_ "github.com/go-sql-driver/mysql"
	//_ "api/db"
	//users "api/users"
	//"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"appengine/datastore"
)

var Contextapp *Context

const KEYVAR string = "AFHSHFADJKHFJKADHFJKADFH"

type Context struct {
	Store		*sessions.CookieStore
	UserKey 	*datastore.Key
	VehicleKey 	*datastore.Key
}

func NewContext() *Context {
	if Contextapp == nil {
		Contextapp := new(Context)
		Contextapp.Store = sessions.NewCookieStore([]byte(KEYVAR))
		return Contextapp
	} else {
		return Contextapp
	}
}
func GetContext() *Context {
	if Contextapp == nil {
		panic("Contextapp is null")
	}
	return Contextapp
}

func (c *Context) SetUserKey(k *datastore.Key) {
	//panic(k)
	c.UserKey = k
	//panic(c.UserKey)
}