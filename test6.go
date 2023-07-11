package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User4 struct {
	gorm.Model
	Name   string
	Age    int64
	Active bool
}

func main() {
	//链接数据库
	db, err := gorm.Open("mysql", "root:jmqs555@(127.0.0.1:3306)"+
		"/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//2、模型与数据库中表对应
	db.AutoMigrate(&User4{})

	//u1 := User4{Name: "qimi", Age: 18, Active: true}
	//db.Create(&u1)
	//u2 := User4{Name: "jinzhu", Age: 20, Active: false}
	//db.Create(&u2)

	//删除
	//var u User4
	//u.ID = 1
	//u.Name = "qimi"
	//db.Delete(&u)

	//db.Where("name=?", "jinzhu").Delete(User4{})
	//db.Delete(User4{},"age=?",18)

	//var u1 []User4
	//db.Debug().Unscoped().Where("name=?", "jinzhu").Find(&u1)
	//fmt.Println(u1)

	//物理删除
	db.Debug().Unscoped().Where("name=?", "jinzhu").Delete(&User4{})
}
