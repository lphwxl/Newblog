package controllers
import (
	"github.com/astaxie/beego"
	"shop/Tools"
	"shop/models"
)


type LoginController struct {
	beego.Controller  // 匿名结构体 -- 实现继承
}

func (that *LoginController) Index() {
	userName := that.Ctx.GetCookie("userName")
	if userName != "" {
		that.Data["userName"] = userName
	}
	uchecked := that.Ctx.GetCookie("uchecked")
	if userName != "" {
		that.Data["uchecked"] = uchecked
	}
	that.TplName = "login.html"
}

// 登陆
func (that *LoginController)ActionLogin(){
	userName := that.GetString("userName")
	password := that.GetString("password")

	if userName == "" || password == ""{
		that.Ctx.WriteString(string(Tools.ResJson(404,"用户名或密码不能为空！",[]string{})))
		return
	}

	// 判断是否设置cookie
	if that.GetString("remember") != ""{
		that.Ctx.SetCookie("userName",userName,140000)
		that.Ctx.SetCookie("uchecked","checked",140000)
	}else {
		that.Ctx.SetCookie("userName",userName,-1)
		that.Ctx.SetCookie("uchecked","checked",-1)
	}
	var user models.Users
	if user.GetUser(userName) != nil{
		that.Ctx.WriteString(string(Tools.ResJson(404,"用户名或密码错误！",[]string{})))
		return
	}

	if user.Password != password {
		that.Ctx.WriteString(string(Tools.ResJson(404,"用户名或密码错误！",[]string{})))
		return
	}

	that.SetSession("userName",user.Name)
	that.Ctx.WriteString(string(Tools.ResJson(200,"登陆成功",[]string{"/article/list"})))
}

