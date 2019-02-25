package routers

import (
	"shop/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.InsertFilter("/article/*",beego.BeforeExec,session)
    beego.Router("/", &controllers.LoginController{},"get:Index")
    beego.Router("/login",&controllers.LoginController{},"post:ActionLogin")
    beego.Router("/user/register",&controllers.UserController{},"get:Register")
    beego.Router("/user/saveUserInfo",&controllers.UserController{},"post:ActionRegister")
    beego.Router("/user/logout",&controllers.UserController{},"get:Logout")

	beego.Router("/article/list",&controllers.ArticleControllers{},"get:Index")
	beego.Router("/article/add",&controllers.ArticleControllers{},"get:ShowAdd;post:Add")
	beego.Router("/article/showArticleDetail",&controllers.ArticleControllers{},"get:ArticleDetail")
	beego.Router("/article/del",&controllers.ArticleControllers{},"get:DelArticle")
	beego.Router("/article/update",&controllers.ArticleControllers{},"get:ArticleShow;post:Update")

	//文章类型
	beego.Router("/article/type",&controllers.ArticleTypeControllers{},"get:ShowList;post:AddType")
	beego.Router("/article/type/del",&controllers.ArticleTypeControllers{},"post:DelType")

}


var session = func(ctx *context.Context) {
	userName := ctx.Input.Session("userName")

	if userName == nil{
		ctx.Redirect(302,"/")
	}


}