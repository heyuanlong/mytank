package rest

import (
	"net/http"
)
type IController interface {
	IBean
	RegisterRoutes() map[string]func(writer http.ResponseWriter,request *http.Request)
	HandleRoutes(writer http.ResponseWriter, request *http.Request)(func(writer http.ResponseWriter, request *http.Request),bool)
}
type BaseController struct {
	Bean
	userDao 		*UserDao
	sessionDao 		*SessionDao
}

func (this *BaseController) Init(context *Context)  {
	this.Bean.Init(context)

	b:= context.GetBean(this.userDao)
	if b,ok := b.(*UserDao);ok{
		this.userDao = b
	}
	b = context.GetBean(this.sessionDao)
	if b, ok := b.(*SessionDao); ok {
		this.sessionDao = b
	}
}
func (this *BaseController) RegisterRoutes() map[string]func(writer http.ResponseWriter,request *http.Request){
	return make(map[string]func(writer http.ResponseWriter, request *http.Request))
}
func (this *BaseController) HandleRoutes(writer http.ResponseWriter, request *http.Request) (func(writer http.ResponseWriter, request *http.Request), bool) {
	return nil, false
}
//需要进行登录验证的wrap包装
func (this *BaseController) Wrap(f func(writer http.ResponseWriter, request *http.Request) *WebResult, qualifiedRole string) func(w http.ResponseWriter, r *http.Request){
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}