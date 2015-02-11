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
)

func SignUp(u *User, r *http.Request) bool {

	c := appengine.NewContext(r)

	if EmailExists(u, r) {
		return false
	}

	k := datastore.NewKey(c, "User", u.Email, 0, nil)
	datastore.Put(c, k, &u)

	return true

	/*
	var sql string

	sql = fmt.Sprintf("insert into users (Email, Password, UserName) values ('%v', '%v', '%v')", u.Email, u.Password, u.UserName)

	dbi.Exec(sql)

	if EmailExists(u) {
		return true
	}

	return false
	*/
}

func EmailExists(u *User, r *http.Request) bool {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("User").
			Filter("Email =", u.Email)
	var x User
	for t := q.Run(c); ; {
		key, err := t.Next(&x)
		_ = key
		_ = err
		if err != nil {
			panic(err)
		}
		return true
	}

	return false

	/*
	var sql string
	sql = fmt.Sprintf("select count(*) from users where Email = '%v'", u.Email)

	var rows *s.Rows

	rows = dbi.Query(sql)
	_ = rows

	if rows == nil {
		return false
	}

	//var ctr int
	var count int = -1
	for rows.Next() {
		
		if err := rows.Scan(&count); err != nil {
            panic(err)
        }
	}
	var ret bool
	ret = false
	if count > 0 {
		ret = true
	}

	return ret
	*/
}

