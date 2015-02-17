package inv

import (
	"net/http"
	"api/context"
	"appengine/datastore"
	"appengine"	
)

type ArticleEntity struct {
	KeyId  		string
	ImgUrl 		string
	Caption  	string
	Title 		string
	Content	 	string
}

func NewArticleEntity(r *http.Request, parentkey *datastore.Key) *ArticleEntity {
	v := new(ArticleEntity)
	v.KeyId = context.RandSeq(32)
	return v
}


func (v *ArticleEntity) Set(imgUrl string, caption string, heading string, content string) {
	v.ImgUrl = imgUrl
	v.Caption = caption
	v.Title = heading
	v.Content = content
}


func (v *ArticleEntity) SetValue(e *ArticleEntity) {
	v.ImgUrl = v.ImgUrl
	v.Caption = v.Caption
	v.Title = v.Title
	v.Content = v.Content
}

func (v *ArticleEntity) Save(r *http.Request, parentkey *datastore.Key) {
	apengcontext := appengine.NewContext(r)
	key := datastore.NewKey(apengcontext, "Article", v.KeyId, 0, parentkey)
	entity := new(ArticleEntity)
	entity.ImgUrl = v.ImgUrl
	entity.Caption = v.Caption
	entity.Title = v.Title
	entity.Content = v.Content
	_, err := datastore.Put(apengcontext, key, entity)
	if err != nil {
		panic(err)
	}
}


func  (v *ArticleEntity) GetAll(r *http.Request, parentkey *datastore.Key) []ArticleEntity {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("Article").Ancestor(parentkey)
	var arr []ArticleEntity
	_,err := q.GetAll(c, &arr)
	if err != nil {
		panic(err)
	}
	return arr
}

func  (v *ArticleEntity) Get(r *http.Request, skey string) []ArticleEntity {
	c := appengine.NewContext(r)
	_=c
	q := datastore.NewQuery("Article").Filter("KeyId =", skey)
	var arr []ArticleEntity
	_,err := q.GetAll(c, &arr)
	if err != nil {
		panic(err)
	}
	return arr
}


func (v *ArticleEntity)  DeleteByParentKey(r *http.Request, parentkey *datastore.Key) bool {
	c := appengine.NewContext(r)
	keys, err := datastore.NewQuery("Article").
				KeysOnly().
				Ancestor(parentkey).
				GetAll(c, nil)
	if err != nil {
		panic(err)
	}
	datastore.DeleteMulti(c, keys)
	return true
}

func  (v *ArticleEntity) DeleteByKey(r *http.Request, skey string) bool {
	c := appengine.NewContext(r)
	keys, err := datastore.NewQuery("Article").
					KeysOnly().
					Filter("KeyId =", skey).
					GetAll(c, nil)
	if err != nil {
		panic(err)
	}
	datastore.DeleteMulti(c, keys)
	return true
}

