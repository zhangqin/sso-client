package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type User struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
}

func (this *MainController) Index() {
	this.Data["ticket"] = this.GetSession("ticket")
	this.Data["username"] = this.GetSession("username")
	this.Data["phone"] = this.GetSession("phone")

	this.TplNames = "index.tpl"
}

func (this *MainController) GetTicket() {
	ticket := this.GetString("ticket")
	if ticket != "" {
		fmt.Println(ticket)
		this.SetSession("ticket", ticket)
		u := this.GetUserInfo(ticket)
		this.SetSession("username", u.Username)
		this.SetSession("phone", u.Phone)
	}
	this.Ctx.WriteString("")
}

func (this *MainController) GetTest() {
	return
}

func (this *MainController) Logout() {
	this.DelSession("ticket")
	this.DelSession("username")
	this.DelSession("phone")
	this.Redirect("/", 302)
}

func (this *MainController) GetUserInfo(ticket string) (u *User) {
	var url = "http://sso.xxx.com:8080/sso/userinfo?ticket=" + ticket
	resp, _ := http.Get(url)
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	u = &User{}
	json.Unmarshal(data, &u)
	fmt.Println(u)
	return u
}
