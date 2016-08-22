package main

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
    gorm.Model

    Username string `gorm:"size:255"`
    Password string `gorm:"size:128"`
    Email string `gorm:"size:255"`
    Posts []Post `gorm:"many2many:user_posts"`
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
    db, err := gorm.Open("mysql",
        "vagrant:vagrant@(vagrant:3306)" +
            "/goorm?charset=utf8&parseTime=True&loc=Local")

    fmt.Println(err)
    defer db.Close()
}
