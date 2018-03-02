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

	db := this.context.DB.Create(user)
	this.PanicError(db.Error)

	return user
}

func (this *UserDao) FindByEmail(email string) *User{
	var user *User = &User{}
	db := this.context.DB.Where(&User{Email:email}).first(user)
	if db.Error != nil {
		return nil
	}
	return user
}