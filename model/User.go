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
