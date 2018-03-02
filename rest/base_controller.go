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

//返回错误的结果。
func (this *BaseController) Error(err interface{}) *WebResult {
	var webResult *WebResult = nil
	if value, ok := err.(string); ok {
		webResult = &WebResult{Code: RESULT_CODE_UTIL_EXCEPTION, Msg: value}
	} else if value, ok := err.(int); ok {
		webResult = ConstWebResult(value)
	} else if value, ok := err.(*WebResult); ok {
		webResult = value
	} else if value, ok := err.(error); ok {
		webResult = &WebResult{Code: RESULT_CODE_UTIL_EXCEPTION, Msg: value.Error()}
	} else {
		webResult = &WebResult{Code: RESULT_CODE_UTIL_EXCEPTION, Msg: "服务器未知错误"}
	}
	return webResult
}