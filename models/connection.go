package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"time"
)

func init(){
	orm.DefaultTimeLoc = time.Local
	orm.RegisterDriver("mysql", orm.DRMySQL)
	info := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		beego.AppConfig.String("UserName"),
			beego.AppConfig.String("PassWord"),
				beego.AppConfig.String("Host"),
					beego.AppConfig.String("Port"),
						beego.AppConfig.String("DB"))

	orm.RegisterDataBase("default","mysql",info)
	orm.RegisterModel(new(Users),new(Article),new(ArticleType))
	//orm.RegisterModel()
}
