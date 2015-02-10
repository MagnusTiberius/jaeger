package users


import (
	"fmt"
	//"net/http"
	//"html/template"
	//"bytes"
	s "database/sql"
	//mysql "api/mysqlmaster"
	dbi "api/db"
	//"github.com/gorilla/mux"
)

func SignUp(u *User) {

	var sql string

	sql = fmt.Sprintf("insert into users (Email, Password, UserName) values ('%v', '%v', '%v')", u.Email, u.Password, u.UserName)

	dbi.Exec(sql)

}

func EmailExists(u *User) bool {
	var sql string

	sql = fmt.Sprintf("select count(*) from users where Email = '%v'", u.Email)

	var rows s.Result

	rows = dbi.Query(sql)
	_ = rows

	//var ctr int

	for rows.Next() {
		
	}


	return true
}

