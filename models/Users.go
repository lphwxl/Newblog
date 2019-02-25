package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"errors"
	"fmt"
)

type Users struct {
	Id int
	Name string
	Password string
	Phone string
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"` //
	UpdateAt time.Time	`orm:"auto_now;type(datetime)"`

	Article []*Article `orm:"rel(m2m);rel_table(user_article)"` // 多对多 对应的关联外键是：article_id (tags_id)
}

// 查询用户
func (that *Users) GetUser(userName string) error   {
	orm := orm.NewOrm()
	that.Name = userName
	err := orm.Read(that,"name")
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("查询失败！")
	}
	return nil
}


// 写入
func(this *Users)Save()int64{
	orm := orm.NewOrm()
	id,err := orm.Insert(this)
	if err != nil {
		return 0
	}
	return id
}

