package users

import (
	//"fmt"
	"net/http"
	//"html/template"
	//"bytes"
	//s "database/sql"
	//mysql "api/mysqlmaster"
	//dbi "api/db"
	//"github.com/gorilla/mux"
	"appengine/datastore"
	"appengine"
	//"github.com/gorilla/sessions"
)

func SignIn(u *User, r *http.Request) (bool, *User) {

	c := appengine.NewContext(r)

	if EmailExists(u, r) == false {
		return false, nil
	}


	q := datastore.NewQuery("User").
			Filter("Email =", u.Email)
	var x User
	for t := q.Run(c); ; {
		key, err := t.Next(&x)
		_ = key
		_ = err
        if err == datastore.Done {
        		//panic("Done")
        		return false, nil
                break // No further entities match the query.
        }
        if err != nil {
        		panic(err)
                c.Errorf("fetching next Person: %v", err)
                break
        }		
        if x.Password == u.Password {
        	u.UserName = x.UserName
        	u.Email = x.Email
	        return true, u
	    }

	}

	return false, nil


}