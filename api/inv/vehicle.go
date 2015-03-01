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

	"appengine/datastore"
	"appengine"	
)

type VehicleEntity struct {
	ManufacturerCode  			string
	ModelCode 					string
	TrimCode 					string
	KeyName 					string
	KeyId 						string
}


func NewVehicleEntity() *VehicleEntity {
	v := new(VehicleEntity)
	return v
}

func AddVehicleEntity(r *http.Request, appcontext *context.Context) (*datastore.Key, bool) {
	c := appengine.NewContext(r)

	manu := r.FormValue("manu")
	model := r.FormValue("model")
	trim := r.FormValue("trim")

	if len(manu) == 0 || len(model) == 0 || len(trim) == 0 {
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

	key := datastore.NewKey(apengcontext, "Vehicle", rndKey, 0, userKey)
	entity := new(VehicleEntity)
	entity.ManufacturerCode = manu
	entity.ModelCode = model
	entity.TrimCode = trim
	entity.KeyName = key.String()
	entity.KeyId = rndKey
	//panic(entity)


	keyVehicle, err := datastore.Put(apengcontext, key, entity)
	if err != nil {
		//panic(err)
        //http.Error(w, err.Error(), http.StatusInternalServerError)
        return nil, false
    }	
    appcontext.VehicleKey = keyVehicle
	return keyVehicle, true
}

func GetUserVehicles() []VehicleEntity {

	var vehicles  []VehicleEntity
	vehicles = []VehicleEntity{}

	return vehicles
}

func GetFfeaturedVehicles(r *http.Request, appcontext *context.Context) []VehicleEntity {
	c := appengine.NewContext(r)

	q := datastore.NewQuery("Vehicle")

    var vehicles []VehicleEntity
    _, err := q.GetAll(c, &vehicles)
    //panic(vehicles)
    if err != nil {
    	panic(err)
    }

    return vehicles

}

func GetVehiclesByKey(r *http.Request, appcontext *context.Context, skey string) []VehicleEntity {
    var vehicles []VehicleEntity
	c := appengine.NewContext(r)
	//apengcontext := appengine.NewContext(r)
	//key := datastore.NewKey(apengcontext, "Vehicle", skey, 0, nil)

	q := datastore.NewQuery("Vehicle").
			Filter("KeyId =", skey)
    _,err := q.GetAll(c, &vehicles)
    //panic(vehicles)
    if err != nil {
    	panic(err)
    }

    return vehicles

}
