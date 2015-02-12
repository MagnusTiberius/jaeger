package inv

type InvItem struct {
	Code  			string
	Description 	string
	Category 		string
}


func NewInvItem() *InvItem {
	v := new(InvItem)
	return v
}

