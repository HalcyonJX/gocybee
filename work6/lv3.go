package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Username string
	Password string
}

var user User

// 连接数据库
func connectDatabase() (db *gorm.DB) {
	dsn := "root:password@tcp(127.0.0.1:3306)/mysql_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("open failed,err:", err)
		return
	}
	return db
}
func AddUser(db *gorm.DB, username, password string) {
	db.Model(&User{}).Create([]map[string]interface{}{
		{"Username": username, "Password": password},
	})
}

// 若没有这个用户返回 false，反之返回 true
func SelectUser(db *gorm.DB, username string) bool {
	db.Where("Username=?", username).Find(&user)
	if user.Password == "" {
		return false
	}
	return true
}

func SelectPasswordFromUsername(db *gorm.DB, username string) string {
	db.Where("Username=?", username).Find(&user)
	return user.Password
}
