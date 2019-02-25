package controllers

import (
	"github.com/astaxie/beego"
	"shop/models"
	"shop/Tools"
	"strconv"
	"github.com/astaxie/beego/orm"
)

type ArticleTypeControllers struct {
	// 匿名结构体 -- 继承
	beego.Controller
}


func (this *ArticleTypeControllers) ShowList()  {

	article := new(models.ArticleType)
	res := article.GetList()
	this.Data["types"] = res
	this.Layout = "layout.html"
	this.TplName = "addType.html"

}
func (this *ArticleTypeControllers) AddType()  {

	typeName := this.GetString("typeName")

	if typeName == "" {
		this.Data["error"] = "分类名称不能为空！"
		this.Layout = "layout.html"
		this.TplName = "addType.html"
		return
	}

	articleType := new(models.ArticleType)
	if articleType.IsExit(typeName) {
		this.Data["error"] = "分类名称已经存在！"
		this.Layout = "layout.html"
		this.TplName = "addType.html"
		return
	}


	articleType.TypeName = typeName

	err := articleType.Save()
	if err != nil {
		this.Data["error"] = "服务器异常，请重试！"
		this.Layout = "layout.html"
		this.TplName = "addType.html"
		return
	}
	this.Redirect("/article/type",302)
}

func (this *ArticleTypeControllers) DelType()  {

	id := this.GetString("id")

	if id == "" {
		this.Ctx.WriteString(string(Tools.ResJson(400,"id 不存在！",[]string{})))
		return
	}

	ids,_ := strconv.ParseInt(id,10,32)

	o := orm.NewOrm()

	o.Begin()
	_,err := o.Delete(&models.ArticleType{Id:int(ids)})
	_,err1 := o.Raw("UPDATE article SET article_type_id = ? WHERE article_type_id = ?",1,int(ids)).Exec()

	//articleType.TypeName = typeName
	if err != nil || err1 !=nil {
		o.Rollback()
		this.Ctx.WriteString(string(Tools.ResJson(400,"删除失败！",[]string{})))
		return
	}

	o.Commit()

	this.Ctx.WriteString(string(Tools.ResJson(200,"删除成功！",[]string{"/article/type"})))
}