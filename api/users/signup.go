package users


import (
	"fmt"
	//"net/http"
	//"html/template"
	//"bytes"
	//s "database/sql"
	dbi "api/db"
	//"github.com/gorilla/mux"
)

func SignUp(u *User) {

	var sql string

	sql = fmt.Sprintf("insert into users (Email, Password, UserName) values ('%v', '%v', '%v')", u.Email, u.Password, u.UserName)

	dbi.Exec(sql)

}