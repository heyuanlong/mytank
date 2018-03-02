package rest

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nu7hatch/gouuid"
	"time"
)

type UserDao struct {
	BaseDao
}

func (this *UserDao) Create(user *User) *User{
	if user == nil {
		panic("参数不能为nil")
	}
	timeUUID,_  := uuid.NewV4()
	user.Uuid = string(timeUUID.String())
	user.CreateTime = time.Now()
	user.ModifyTime = time.Now()
	user.LastTime = time.Now()
	user.Sort = time.Now().UnixNano() / 1e6



	return user
}