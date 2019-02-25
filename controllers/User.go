package controllers

import (
	"github.com/astaxie/beego"
	"shop/models"
	"shop/Tools"
)


type UserController struct {
	beego.Controller
}




func(this *UserController) Register(){

	this.TplName = "register.html"
}

func (this *UserController) ActionRegister()  {
	userName := this.GetString("userName")
	password := this.GetString("password")

	if userName == "" || password == ""{
		this.Ctx.WriteString(string(Tools.ResJson(4000,"用户名或密码不能为空！",[]string{})))
		return
	}
	var user models.Users
	if user.GetUser(userName) == nil {
		this.Ctx.WriteString(string(Tools.ResJson(4000,"用户名已经存在！",[]string{})))
		return
	}

	user.Name = userName
	user.Password = password
	//user.CreatedAt = time.Time{}
	id := user.Save()

	if id != 0 {
		this.Ctx.WriteString(string(Tools.ResJson(200,"注册成功！",[]string{"/"})))
		return
	}
	this.Ctx.WriteString(string(Tools.ResJson(200,"注册失败，请重试！",[]string{})))

}


func(this *UserController)Logout(){
	this.DelSession("userName")
	this.DestroySession()

	this.Redirect("/",302)

}