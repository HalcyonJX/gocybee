package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	ID      int
	Name    string
	Age     int
	Address string
}

func main() {
	dsn := "root:Zjx666946.@tcp(127.0.0.1:3306)/mysql_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("open failed,err:", err)
		return
	}
	//迁移
	db.AutoMigrate(&UserInfo{})
	//创建
	user1 := UserInfo{3, "zhang", 18, "China"}
	user2 := UserInfo{4, "li", 26, "America"}
	db.Create(&user1)
	db.Create(&user2)
	//查询
	var user UserInfo
	db.First(&user)

	//更新
	db.Model(&UserInfo{}).Where("ID", 2).Update("address", "China")
	//删除
	db.Where("ID", 1).Delete(&user)
}
