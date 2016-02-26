package routers

import (
	"github.com/astaxie/beego"
	"github.com/zhangqin/sso-client/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/get_ticket", &controllers.MainController{}, "*:GetTicket")
	beego.Router("/logout", &controllers.MainController{}, "*:Logout")
	beego.Router("/get_test", &controllers.MainController{}, "*:GetTest")
}
