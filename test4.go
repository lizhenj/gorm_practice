package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User2 struct {
	gorm.Model
	Name string
	Age  int64
}

func main4() {
	//链接数据库
	db, err := gorm.Open("mysql", "root:jmqs555@(127.0.0.1:3306)"+
		"/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//2、模型与数据库中表对应
	db.AutoMigrate(&User2{})

	//3、创建
	//u1 := User2{Name: "qimi", Age: 18}
	//db.Create(&u1)
	//u2 := User2{Name: "jinzhu", Age: 20}
	//db.Create(&u2)

	//4、查询
	//var user User2
	user := new(User2)
	db.First(&user)
	fmt.Printf("u: %+v\n", user)

	var users []User2
	db.Find(&users)
	fmt.Printf("us: %+v\n", users)

	user1 := new(User2)
	db.Where("name = ?", "jinzhu").First(&user1)
	fmt.Printf("u1: %+v\n", user1)

	user2 := new(User2)
	db.First(&user2, "name = ?", "jinzhu")
	fmt.Printf("u2: %+v\n", user2)

	user3 := new(User2)
	db.Debug().FirstOrInit(&user3, User2{Name: "lzj"})
}
