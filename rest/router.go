package rest

import (
	"fmt"
	"github.com/json-iterator/go"
	"io"
	"net/http"
	"os"
	"strings"
)

type Router struct {
	context *Context
	routeMap map[string]func(w http.ResponseWriter,r *http.Request)
}

func NewRouter(context *Context) *Router  {
	router := &Router{
		context :		context,
		routeMap: 		make(map[string]func(w http.ResponseWriter,r *http.Request)),
	}
	for _,controller := range context.ControllerMap{
		routes := controller.RegisterRoutes()
		for k,v := range routes{
			router.routeMap[k] = v
		}
	}
	return router
}

func (this *Router) GlobalPanicHandle( writer http.ResponseWriter, request *http.Request )  {
	if err:= recover();err != nil{
		LogError(fmt.Sprintf("全局异常: %v", err))

		var webResult *WebResult = nil
		if value , ok := err.(string);ok{
			webResult = &WebResult{Code: RESULT_CODE_UTIL_EXCEPTION, Msg: value}
		}else if value, ok := err.(int); ok {
			webResult = ConstWebResult(value)
		}else if value, ok := err.(*WebResult); ok {
			webResult = value
		} else if value, ok := err.(WebResult); ok {
			webResult = &value
		}else if value, ok := err.(error); ok {
			webResult = &WebResult{Code: RESULT_CODE_UTIL_EXCEPTION, Msg: value.Error()}
		} else {
			webResult = &WebResult{Code: RESULT_CODE_UTIL_EXCEPTION, Msg: "服务器未知错误"}
		}

		writer.Header().Set("Content-Type","application/json;charset=UTF-8")

		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		b,_ := json.Marshal(webResult)
		if webResult.Code == RESULT_CODE_OK {
			writer.WriteHeader(http.StatusOK)
		} else {
			writer.WriteHeader(http.StatusBadRequest)
		}
		fmt.Fprintf(writer, string(b))
	}
}

func (this *Router ) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer this.GlobalPanicHandle(writer,request)

	path := request.URL.Path
	if strings.HasPrefix(path,"/api"){
		if handler,ok := this.routeMap[path];ok{
			handler(writer,request)
		}else{
			canHandle := false
			for _,controller := range this.context.ControllerMap{
				if handler ,exist := controller.HandleRoutes(writer,request);exist{
					canHandle =  true
					handler(writer,request)
					break
				}
			}
			if !canHandle {
				panic(fmt.Sprintf("没有找到能够处理%s的方法\n", path))
			}
		}
	}else{
		//当作静态资源处理。默认从当前文件下面的static文件夹中取东西。
		dir := GetHtmlPath()
		requestURI := request.RequestURI
		if requestURI == "" || requestURI =="/"{
			requestURI = "/index.html"
		}
		filePath := dir + requestURI
		LogDebug(filePath)

		exists ,_ := PathExists(filePath)
		if !exists{
			panic(fmt.Sprintf("404 not found:%s", requestURI))
		}
		writer.Header().Set("Content-Type",GetMimeType( GetExtension(filePath) ))
		diskFile ,err := os.Open(filePath)
		if err != nil {
			panic("cannot get file.")
		}
		defer diskFile.Close()
		_,err = io.Copy(writer,diskFile)
		if err != nil {
			panic("cannot get file.")
		}
	}
	
}


