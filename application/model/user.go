package model

import (
	"github.com/jinzhu/gorm"
	. "urils/application/database"
)

/**
用户实体
*/

type User struct {
	gorm.Model
	Username      string `validate:"required" label:"账户名" gorm:"unique_index;size:255;comment:'账户名'" json:"username"`
	HashPassword  string `gorm:"not null;size:255;comment:'密码'" json:"-"`
	Password      string `validate:"required" label:"登陆密码" gorm:"-" json:"password,omitempty"`
	Nickname      string `label:"昵称" gorm:"size:255;comment:'昵称'" json:"nickname" sql:"index"`
	Mobile        string `gorm:"index;size:15;comment:'手机';" json:"mobile"`
	Email         string `validate:"omitempty,email" gorm:"index;size:255;comment:'邮箱';" json:"email"`
	Avatar        string `gorm:"size:255;comment:'头像'" json:"avatar"`
	Sex           bool   `gorm:"type:boolean;default:true;comment:'性别'" json:"sex"`
	Ip            string `validate:"omitempty,ipv4" gorm:"size:255;comment:'IP地址';" json:"ip"`
	Status        bool   `gorm:"type:boolean;default:true;comment:'状态'" json:"status"`
}


/**
设置表名
*/

func (User) TableName() string {
	return "user_info"
}


/**
创建用户
*/

func (user User) Insert() (id uint, err error) {
	//添加数据
	result := Orm.Create(&user)
	return user.ID, result.Error
}


/**
根据指定ID获取用户
*/

func (user *User) GetOneById(id uint){
	Orm.First(&user, id)
}


/**
根据账户信息(用户名、手机、邮箱)获取用户
*/

func (user *User) GetOneByAccount(account string) {
	Orm.First(&user, "username = ? or mobile = ? or email= ?", account, account, account)
}


/**
获取所有用户
*/

func (user User) GetAll() []User {
	var users []User
	Orm.Find(&users)
	return users
}


/**
更新密码
*/

func (user User) ChangePassword(HashPasswrd string) {
	user.HashPassword = HashPasswrd
	Orm.Save(&user)
}
