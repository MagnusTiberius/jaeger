package inv

import (
	"net/http"
	"api/context"
	"appengine/datastore"
	"appengine"	
)

type CarouselEntity struct {
	KeyId  		string
	ImgUrl 		string
	Caption  	string
	Heading 	string
	Content	 	string
}

func NewCarouselEntity(r *http.Request, parentkey *datastore.Key) *CarouselEntity {
	v := new(CarouselEntity)
	v.KeyId = context.RandSeq(32)
	return v
}


func (v *CarouselEntity) Set(imgUrl string, caption string, heading string, content string) {
	v.ImgUrl = imgUrl
	v.Caption = caption
	v.Heading = heading
	v.Content = content
}


func (v *CarouselEntity) SetValue(e *CarouselEntity) {
	v.ImgUrl = v.ImgUrl
	v.Caption = v.Caption
	v.Heading = v.Heading
	v.Content = v.Content
}

func (v *CarouselEntity) Save(r *http.Request, parentkey *datastore.Key) {
	apengcontext := appengine.NewContext(r)
	key := datastore.NewKey(apengcontext, "Carousel", v.KeyId, 0, parentkey)
	entity := new(CarouselEntity)
	entity.ImgUrl = v.ImgUrl
	entity.Caption = v.Caption
	entity.Heading = v.Heading
	entity.Content = v.Content
	_, err := datastore.Put(apengcontext, key, entity)
	if err != nil {
		panic(err)
	}
}


func  GetAll(r *http.Request, parentkey *datastore.Key) []CarouselEntity {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("Carousel").Ancestor(parentkey)
	var arr []CarouselEntity
	_,err := q.GetAll(c, &arr)
	if err != nil {
		panic(err)
	}
	return arr
}

func  Get(r *http.Request, skey string) []CarouselEntity {
	c := appengine.NewContext(r)
	_=c
	q := datastore.NewQuery("Carousel").Filter("KeyId =", skey)
	var arr []CarouselEntity
	_,err := q.GetAll(c, &arr)
	if err != nil {
		panic(err)
	}
	return arr
}


func  DeleteByParentKey(r *http.Request, parentkey *datastore.Key) bool {
	c := appengine.NewContext(r)
	keys, err := datastore.NewQuery("Carousel").
				KeysOnly().
				Ancestor(parentkey).
				GetAll(c, nil)
	if err != nil {
		panic(err)
	}
	datastore.DeleteMulti(c, keys)
	return true
}

func  DeleteByKey(r *http.Request, skey string) bool {
	c := appengine.NewContext(r)
	keys, err := datastore.NewQuery("Carousel").
					KeysOnly().
					Filter("KeyId =", skey).
					GetAll(c, nil)
	if err != nil {
		panic(err)
	}
	datastore.DeleteMulti(c, keys)
	return true
}

