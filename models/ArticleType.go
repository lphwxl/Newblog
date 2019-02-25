package models

import (

	"github.com/astaxie/beego/orm"
)

type ArticleType struct {
	Id int
	TypeName string
	Articles []*Article `orm:"reverse(many)"`
}

func(this *ArticleType) GetList() []*ArticleType {
	o := orm.NewOrm()

	var collection []*ArticleType
	o.QueryTable("article_type").All(&collection)
	return collection
}


func(this *ArticleType) IsExit(typeName string) bool{
	o := orm.NewOrm()
	return o.QueryTable("article_type").Filter("type_name",typeName).Exist()
}

func(this *ArticleType) Save()error {
	o := orm.NewOrm()
	_,err := o.Insert(this)
	return  err
}

