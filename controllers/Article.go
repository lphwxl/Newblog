package controllers

import (
	"github.com/astaxie/beego"
	"time"
	"path"
	"shop/models"

	"github.com/astaxie/beego/orm"

	"math"
	"fmt"
)

const (
	FILETYPE  = "image/png"
	PAGESIZE = 3
	CURRENTPAGE = 1
	)

type ArticleControllers struct {
	// 匿名结构体 -- 继承
	beego.Controller
}

func (this *ArticleControllers) Index()  {
	// 总数据
	orm := orm.NewOrm()
	qs := orm.QueryTable(new(models.Article)).RelatedSel("ArticleType")
	var articles []*models.Article
	currentPage,err := this.GetInt("page")
	if err != nil{
		 currentPage = CURRENTPAGE
	}
	offset := (currentPage-1)*PAGESIZE

	typeId,err1 := this.GetInt("typeId")
	this.Data["typeId"] = typeId
	var count int64
	if err1 != nil || typeId == 0{
		count,_ = orm.QueryTable("article").Count()
		qs.Limit(PAGESIZE,offset).All(&articles)
	}else {
		count,_ = orm.QueryTable("article").Filter("ArticleType__id",typeId).Count()
		qs.Filter("ArticleType__id",typeId).Limit(PAGESIZE,offset).All(&articles)
	}
	// go 语言中，不同的数据类型之间是不能进行运算的；必须先转换，后运算

	if currentPage == 1{
		this.Data["prePage"] = currentPage
	}else {
		this.Data["prePage"] = currentPage-1
	}

	this.Data["page"] =  currentPage
	this.Data["articles"] = articles

	//总记录数

	this.Data["count"] = count
	//总页数
	total := math.Ceil(float64(count)/float64(PAGESIZE))
	this.Data["total"] = total
	//fmt.Println(total > float64(currentPage+1))
	if total >= float64(currentPage+1){
		this.Data["nextPage"] = currentPage+1
	}else {
		this.Data["nextPage"] = currentPage
	}
	// 分类类型
	var types []*models.ArticleType
	orm.QueryTable("article_type").All(&types)
	this.Data["types"] = types

	userName  := this.GetSession("userName")
	this.Data["userName"] = userName.(string)
	this.Layout = "layout.html"
	this.TplName = "index.html"

}


func (this *ArticleControllers) ShowAdd()  {

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	// 总数据
	orm := orm.NewOrm()
	// 分类类型
	var types []*models.ArticleType
	orm.QueryTable("article_type").All(&types)
	this.Data["types"] = types
	this.Layout = "layout.html"
	this.TplName = "add.html"

}


func (this *ArticleControllers) Add()  {
	content := this.GetString("content")
	title := this.GetString("articleName")
	articleType,_ := this.GetInt("articleType")
	if content == "" || title == "" {
		this.Data["errmsg"] = "文章标题/内容不能为空！"
		this.TplName = "add.html"
		return
	}

	_,head,err := this.GetFile("uploadname")
	// 1.是否上传成功
	if err != nil {
		this.Data["errmsg"] = "请选择上传文件！"
		this.TplName = "add.html"
		return
	}
	// 2. 判断文件类型
		val,ok := head.Header["Content-Type"]
		if ok != true{
			this.Data["errmsg"] = "文件类型错误！"
			this.TplName = "add.html"
			return
		}
		if FILETYPE != val[0]{
			this.Data["errmsg"] = "文件类型错误！"
			this.TplName = "add.html"
			return
		}
	// 3. 判断文件大小
	if head.Size > 4000000 {
		this.Data["errmsg"] = "文件过大！"
		this.TplName = "add.html"
		return
	}
	// 4. 上传文件
	fileName := time.Now().Format("2006-01-02-15:04:05") + path.Ext(head.Filename) // 获取文件的后缀
	err = this.SaveToFile("uploadname","./static/img/"+fileName)

	if err != nil {
		this.Data["errmsg"] = "文件上传失败！"
		this.TplName = "add.html"
		return
	}
	// 存储数据
	var article models.Article
	article.Title = title
	article.Content = content
	article.ImgPath = "/static/img/"+fileName
	article.ArticleType = &models.ArticleType{Id:articleType}
	err = article.Save()
	if err != nil{
		this.Data["errmsg"] = "添加失败！"
		this.Layout = "layout.html"
		this.TplName = "add.html"
		return
	}
	this.Redirect("/article/list",302)

}

func(this *ArticleControllers)ArticleDetail(){
	id,err := this.GetInt("articleId")
	if err != nil || id ==0 {
		this.Redirect("/article/list",302)
		return
	}

	article := models.Article{Id:id}
	article1 := article.GetOne()
	if article1 == nil {
		this.Redirect("/article/list",302)
		return
	}
	userName := this.GetSession("userName")
	userN,ok := userName.(string)
	if !ok {
		beego.Info("用户不存在！")
		this.Redirect("/article/list",302)
		return
	}
	if !article.ReadUser(userN){
		this.Redirect("/article/list",302)
		return
	}
	o :=  orm.NewOrm()
	var users []*models.Users
	o.QueryTable("users").Filter("Article__Article__id",id).Distinct().All(&users)

	o.QueryTable("article").Update(orm.Params{
		"Acount": orm.ColValue(orm.ColAdd, 1),
	})
	this.Data["users"] = users
	this.Data["article"] = article1
	this.Layout = "layout.html"
	this.TplName = "content.html"
}

func (this *ArticleControllers)DelArticle()  {
	id,_ := this.GetInt("id")

	o := orm.NewOrm()
	article := models.Article{Id:id}
	if o.Read(&article) != nil{
		this.Redirect("/article/list",302)
		return
	}

	o.Delete(&article)
	this.Redirect("/article/list",302)
	return

}

func (this *ArticleControllers) ArticleShow()  {
	id,_ := this.GetInt("id")
	o := orm.NewOrm()
	var article models.Article
	o.QueryTable("article").Filter("id",id).RelatedSel("ArticleType").One(&article)

	this.Data["article"] =  article
	this.Layout = "layout.html"
	this.TplName = "update.html"
}

func (this *ArticleControllers)Update()  {
	content := this.GetString("content")
	title := this.GetString("articleName")
	//articleType,_ := this.GetInt("articleType")
	id,_ := this.GetInt("id")

	data := orm.Params{} // map[s]
	if title != ""{
		data["title"] = title
	}
	if content != ""{
		data["content"] = content
	}
	_,head,err := this.GetFile("uploadname")
	// 1.是否上传成功
	if err == nil {
		beego.Info(12)
		// 2. 判断文件类型
		val,ok := head.Header["Content-Type"]
		if ok != true{
			this.Data["errmsg"] = "文件类型错误！"
			this.TplName = "update.html"
			return
		}
		if FILETYPE != val[0]{
			this.Data["errmsg"] = "文件类型错误！"
			this.TplName = "update.html"
			return
		}
		// 3. 判断文件大小
		if head.Size > 4000000 {
			this.Data["errmsg"] = "文件过大！"
			this.TplName = "update.html"
			return
		}
		// 4. 上传文件
		fileName := time.Now().Format("2006-01-02-15:04:05") + path.Ext(head.Filename) // 获取文件的后缀
		err = this.SaveToFile("uploadname","./static/img/"+fileName)

		if err != nil {
			this.Data["errmsg"] = "文件上传失败！"
			this.TplName = "update.html"
			return
		}
		data["img_path"] = "/static/img/"+fileName
	}

	o := orm.NewOrm()
	_,err1 := o.QueryTable("article").Filter("id",id).Update(data)

	if err1 != nil{
		this.Data["errmsg"] = "更新失败！"
		this.TplName = ".updatehtml"
		return
	}
	this.Redirect("/article/list",302)
}