package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User3 struct {
	gorm.Model
	Name   string
	Age    int64
	Active bool
}

func main5() {
	//链接数据库
	db, err := gorm.Open("mysql", "root:jmqs555@(127.0.0.1:3306)"+
		"/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//2、模型与数据库中表对应
	db.AutoMigrate(&User3{})

	//u1 := User3{Name: "qimi", Age: 18, Active: true}
	//db.Create(&u1)
	//u2 := User3{Name: "jinzhu", Age: 20, Active: false}
	//db.Create(&u2)

	//4、查询
	//var user User2
	user := new(User3)
	db.First(&user)
	fmt.Printf("u: %+v\n", user)
	//更新
	user.Name = "wuwu"
	user.Age = 99
	db.Save(&user) //默认修改所有字段
	db.Model(&user).Update("name", "HHH")

	m1 := map[string]interface{}{
		"name":   "UUU",
		"age":    999,
		"active": false,
	}
	db.Model(&user).Select("name").Updates(m1)                       //只更新name字段
	rowsNum := db.Model(&user).Omit("name").Updates(m1).RowsAffected //排除name字段，更新其他
	fmt.Println(rowsNum)

	//user表中用户年龄+2
	db.Debug().Model(&User3{}).Update("age", gorm.Expr("age*?+?", 2, 2))
	//db.Debug().Model(&user).UpdateColumn("name", "555")
}
