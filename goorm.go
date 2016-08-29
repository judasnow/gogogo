package main

// gorm 的测试

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Profile struct {
    gorm.Model
    Name string
}

type User struct {
    gorm.Model
    Profile      Profile `gorm:"ForeignKey:ProfileRefer"` // use ProfileRefer as foreign key
    ProfileRefer int
}

type User1 struct {
    gorm.Model

    Username string `gorm:"size:255"`
    Password string `gorm:"size:128"`
    Email string `gorm:"size:255"`
    Posts []Post `gorm:"many2many:user_posts"`
}

// 全局关闭表名的复数化
// db.SingularTable(true)

// 修改默认的表名
func (u User) TableName() string {
    return "t_users"
}

type Post struct {
    gorm.Model

    UserID uint
    Content string `gorm:"type:varchar(140);"`
}

type Relationship struct {
    gorm.Model

    UserID uint
    PostID uint
}

func main() {
    db, _ := gorm.Open("mysql",
        "vagrant:vagrant@(vagrant:3306)" +
            "/gorm?charset=utf8&parseTime=True&loc=Local")

    db.CreateTable(&Profile{})
    db.CreateTable(&User{})

    defer db.Close()
}
