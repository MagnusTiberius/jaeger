package wsusers

import (
	//"fmt"
	"net/http"
	users "api/users"
	"appengine/datastore"
	"appengine"
)

func HandleWsUserList(w http.ResponseWriter, r *http.Request) []users.User {
	c := appengine.NewContext(r)

	q := datastore.NewQuery("User")

    var users []users.User
    res, err := q.GetAll(c, &users)
    _ = res
    //panic(res)
    if err != nil {
    	panic(err)
    }
    //panic(users)

    usrs := []users.User{[]{Email:"test@email.com"}}

    return users	
}