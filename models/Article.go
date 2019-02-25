package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type Article struct {
	Id int
	Title string
	Content string
	ImgPath string
	Acount int
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`

	ArticleType *ArticleType `orm:"rel(fk);on_delete(do_nothing)"`  // 对应的外键是 ： article_type_id (Tags_id) --> 一对多
	// on_delete 设置是否级联删除 -- 一般不需要级联删除
	User []*Users `orm:"reverse(many);rel_table(user_article)"` // 多对多 对应的关联外键是：user_id (tags_id)
}


//保存
func (this *Article) Save() error{
	orm := orm.NewOrm()
	_,err := orm.Insert(this)
	return  err
}

func (this *Article)GetOne() *Article  {
	o := orm.NewOrm()
	var article Article
	err := o.QueryTable("article").Filter("id",this.Id).RelatedSel("ArticleType").One(&article)

	if err != nil {
		return nil
	}
	return &article
}


func (this *Article)ReadUser(userName string) bool  {
	o := orm.NewOrm()

	if o.Read(this) != nil{
		beego.Info("写入失败！")
		return false
	}

	m2m := o.QueryM2M(this,"User")

	user := Users{Name:userName}
	if o.Read(&user,"name") != nil{
		beego.Info("写入失败！")
		return false
	}
   	_,err := m2m.Add(user)
	if err != nil {
		return false
	}
	return true
}
