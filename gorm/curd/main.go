package main

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

import "gorm.io/driver/mysql"

func ConnectDB(host, name, username, password string, port int) (db *gorm.DB, err error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", username, password, host, port, name)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}

func main() {

}

//声明模型

type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullString
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

//高级选项
//字段级权限控制

type User2 struct {
	Name string `gorm:"<-:create"`
}

func Create() {
	db, err := ConnectDB("", "", "", "", 0)
	if err != nil {
		//todo:改方法的作用
		log.Fatal(err)
	}

	//创建记录
	user := User{
		Name: "jinzhu",
		Age:  18,
	}
	result := db.Create(&user)
	log.Println(user.ID)
	log.Println(result.Error)
	log.Println(result.RowsAffected)

	//用指定的字段创建记录
	db.Select("Name", "Age", "CreatedAt").Create(&user)

	db.Omit("Name", "Age", "CreatedAt").Create(&user)

	//批量插入
	var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	db.Create(&users)

	for _, user := range users {
		log.Println(user.ID)
	}

	//分批创建
	db.CreateInBatches(users, 100)

}
