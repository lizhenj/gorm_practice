package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//1、定义模型
type Uuser struct {
	ID   int64
	Name *string `gorm:"default:'xiaoer'"`
	Age  int64
}

func main3() {
	//链接数据库
	db, err := gorm.Open("mysql", "root:jmqs555@(127.0.0.1:3306)"+
		"/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//2、模型与数据库中表对应
	db.AutoMigrate(&Uuser{})

	//3、创建
	//u := Uuser{Name: "hbb", Age: 18}
	u := Uuser{Name: new(string), Age: 65}
	fmt.Println(db.NewRecord(&u)) //判断主键是否为空，进而可得知该记录是否存在
	db.Debug().Create(&u)
	fmt.Println(db.NewRecord(&u))

}
