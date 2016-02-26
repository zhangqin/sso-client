package main

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	_ "github.com/zhangqin/sso-client/routers"
)

var loginFilter = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("ticket").(string)
	fmt.Println(ctx.Request.RequestURI)
	if !ok && ctx.Request.RequestURI != "/" && !strings.HasPrefix(ctx.Request.RequestURI, "/get_ticket") {
		fmt.Println("redirect")
		ctx.Redirect(302, "/")
	}
}

func main() {
	beego.InsertFilter("/*", beego.BeforeRouter, loginFilter)
	beego.Run()
}
