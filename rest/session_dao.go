package rest

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
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