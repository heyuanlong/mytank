package rest

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nu7hatch/gouuid"

)

type SessionDao struct {
	BaseDao
}

//构造函数
func NewSessionDao(context *Context) *SessionDao {

	var sessionDao = &SessionDao{}
	sessionDao.Init(context)
	return sessionDao
}
func (this *SessionDao) Create(session *Session) *Session{
	timeUUID,_ := uuid.NewV4()
	session.Uuid = string(timeUUID.String())
	db:= this.context.DB.Create(session)
	this.PanicError(db.Error)

	return session
}