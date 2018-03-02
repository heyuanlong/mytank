package rest

import (
	"net/http"
	"regexp"
	"strconv"
	"time"
)


type UserController struct {
	BaseController
}

func (this *UserController) Init(context *Context)  {
	this.BaseController.Init(context)
}

func (this *UserController) RegisterRouters()map[string]func(w http.ResponseWriter,r *http.Request)  {
	routeMap := make(map[string]func(w http.ResponseWriter, r *http.Request))
	routeMap["/api/user/login"] = this.Wrap()
}

func (this *UserController) Login(writer http.ResponseWriter, request *http.Request) *WebResult  {
	email := request.FormValue("email")
	password := request.FormValue("password")
	if "" == email || "" == password {
		return this.Error("请输入邮箱和密码")
	}
	user := this.userDao.FindByEmail(email)
	if user == nil {
		return this.Error("邮箱或密码错误")
	} else {
		if !MatchBcrypt(password,user.Password){
			return this.Error("邮箱或密码错误")
		}
	}
	expiration := time.Now()
	expiration = expiration.AddDate(0,0,7)

	/*
	//持久化用户的session.
	session := &Session{
		UserUuid:   user.Uuid,
		Ip:         GetIpAddress(request),
		ExpireTime: expiration,
	}
	session.ModifyTime = time.Now()
	session.CreateTime = time.Now()
	session = this.sessionDao.Create(session)
	*/
	//设置用户的cookie
	cookie := http.Cookie{
		Name:		COOKIE_AUTH_KEY,
		Path: 		"/",
		Value:
	}

}