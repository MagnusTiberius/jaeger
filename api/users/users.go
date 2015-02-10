package users

type User struct {
	Email  		string
	Password 	string
	UserName	string
	FirstName	string
	LastName	string
	UserID		int
}


func NewUser() *User {
	v := new(User)
	return v
}

