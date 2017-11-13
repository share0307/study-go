package models

import (
    "time"
    "github.com/astaxie/beego/orm"
    "sql-"
)

type Category struct {
    Id int64
    Title string
    Created time.Time `orm:"index""`
    Views int64 `orm:"index"`
    TopicTime   time.Time `orm:"time"`
    TopCount int64
    TopicLastUserId int64
}

type Topic struct {
    Id int64
    Uid int64
    Title string
    Content string `orm:"size(5000)"`
    Attachment string
    Created time.Time `orm:"index""`
    Updated time.Time `orm:"index""`
    Views int64 `orm:"index"`
    Author string
    ReplyTime time.Time `orm:"index"`
    ReplyConut int64
    ReplyLastUserId int64
}

/**
注册数据库时会被调用
 */
func RegisterDB()  {
    
}