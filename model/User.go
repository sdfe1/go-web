package model

import (
	"Project1/message"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	//已经默认添加了一个自增的id主键字段
	Username string `gorm:"type: varchar(30);not null" json:"username"`
	Password string `gorm:"type: varchar(30);not null" json:"password"`
}

func CreateUser(user *User) int {
	err := db.Create(&user).Error
	fmt.Printf("正在创建用户: %+v\n", user)
	if err != nil {
		log.Printf("用户创建失败: %v", err)
		return message.ERROR
	}
	return message.SUCCESS
}

// 检验登录
func CheckLogin(user *User) int {
	var dbUser User
	//1.检验用户名是否存在
	result := db.Where("username = ?", user.Username).First(&dbUser)
	if result == nil {
		return message.ERROR_USER_NO_RIGHT
	}
	//2.检验对应的密码是否正确
	if user.Password != dbUser.Password {
		return message.ERROR_PASSWORD_WRONG
	}
	return message.SUCCESS
}
