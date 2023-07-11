package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

//定义模型
type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64 //零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        //设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` //设置会员号（member number）唯一且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  //设置num为自增类型
	Address      string  `gorm:"index:addr"`      //给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               //忽略本字段
}

type Admin struct {
	Id   int
	Naem string
}

func (Admin) TableName() string {
	return "dada"
}

func main2() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "sos_" + defaultTableName
	}

	//链接数据库
	db, err := gorm.Open("mysql", "root:jmqs555@(127.0.0.1:3306)"+
		"/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Admin{})

	//使用User结构体创建xiaoqi的表
	db.Table("xiaoqi").CreateTable(&Admin{})
}
