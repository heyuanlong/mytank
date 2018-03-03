package rest

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
)

//上下文，管理数据库连接，管理所有路由请求，管理所有的单例component.
type Context struct {
	DB 		*gorm.DB
	Router		*Router
	BeanMap 	map[string]IBean
	ControllerMap map[string]IController
}

func (this *Context) OpenDb()  {
	var err error = nil
	this.DB  ,err = gorm.Open("mysql",CONFIG.MysqlUrl)
	if err != nil{
		panic("failed to connect mysql database")
	}
	this.DB.LogMode(false)
}
func (this *Context) CloseDb() {
	if this.DB != nil {
		this.DB.Close()
	}
}
func NewContext() *Context  {
	context := &Context{}
	context.OpenDb()
	context.BeanMap = make(map[string]IBean)
	context.ControllerMap = make(map[string]IController)

	context.registerBeans()
	context.initBeans()

	context.Router = NewRouter(context)
	return context
}

func (this *Context) registerBean(bean interface{}) {
	typeOf := reflect.TypeOf(bean)
	typeName := typeOf.String()

	if element,ok := bean.(IBean);ok{
		//LogDebug("-------------" + reflect.TypeOf(element).String() )
		if _,ok:= this.BeanMap[typeName];ok{
			LogError(fmt.Sprintf("【%s】已经被注册了，跳过。", typeName))
		}else{
			this.BeanMap[typeName] = element
			LogDebug("BeanMap:" + typeName)
			if controller, ok1 := bean.(IController); ok1 {
				LogDebug("ControllerMap:" + typeName)
				this.ControllerMap[typeName] = controller
			}
		}
	}else{
		err := fmt.Sprintf("注册的【%s】不是Bean类型。", typeName)
		panic(err)
	}
}

//注册各个Beans
func (this *Context) registerBeans() {
	this.registerBean( new(UserDao) )
	this.registerBean( new(SessionDao))

	this.registerBean( new(UserController))
}

func (this *Context) GetBean(bean interface{}) IBean {
	typeOf := reflect.TypeOf(bean)
	typeName := typeOf.String()
	if val,ok := this.BeanMap[typeName];ok{
		return val
	}else {
		err := fmt.Sprintf("【%s】没有注册。", typeName)
		panic(err)
	}
}

//初始化每个Bean
func (this *Context) initBeans() {
	for _, bean := range this.BeanMap {
		bean.Init(this)
	}
}
//销毁的方法
func (this *Context) Destroy() {
	this.CloseDb()

}


