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
	"api/context"
)

func SignIn(u *User, w http.ResponseWriter, r *http.Request, appcontext *context.Context) (bool, *User, *datastore.Key) {


	c := appengine.NewContext(r)

	if EmailExists(u, r) == false {
		return false, nil, nil
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
        		return false, nil, nil
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
        	u.KeyIdInt = key.IntID()
        	u.KeyIdString = key.StringID()

        	//appcontext := context.GetContext()
        	//appcontext.SetUserKey(key)
        	//appcontext.UserKey = key
        	session, _ := appcontext.Store.Get(r, "jaegersignup")
			session.Values["Email"] = u.Email
			session.Values["UserName"] = u.UserName        	
        	session.Save(r, w)
        	//appcontext = context.GetContext()
        	//panic(key)
	        return true, u, key
	    }

	}

	return false, nil, nil


}