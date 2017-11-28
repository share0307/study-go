package models

import (
    "time"
    "github.com/astaxie/beego/orm"
    _ "github.com/mattn/go-sqlite3"
    "os"
    "path"
)

const(
    CONNECTION_INFO = "data/beeblog.db"
    DB_NAME = "default"
    SQLLITE3_DRIVER = "sqlite3"

)

type Category struct {
    Id int64
    Title string
    Created time.Time `orm:"index""`
    Views int64 `orm:"index"`
    TopicTime   time.Time `orm:"time"`
    TopicCount int64
    TopicLastUserId int64
}

type Topic struct {
    Id int64
    Uid int64
    CategoryId int64 `orm:index`
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

type Comment struct {
	Id int64
	Tid int64
	Name string
	Content string `orm:"size(1000)"`
	Created time.Time `orm:"index"`
}

/**
注册数据库时会被调用
 */
func RegisterDB()  {
    //获取文件说在的目录
    dbDir := path.Dir(CONNECTION_INFO)
    //判断目录是否存在，不存在则创建
    if _,err := os.Stat(dbDir);err != nil{
        //创建目录
        os.MkdirAll(dbDir,os.ModePerm);
        //创建文件
        os.Create(CONNECTION_INFO)
    }

    // 需要在init中注册定义的model
    orm.RegisterModel(new(Category),new(Topic),new(Comment))

    // 参数1   driverName
    // 参数2   数据库类型
    // 这个用来设置 driverName 对应的数据库类型
    // mysql / sqlite3 / postgres 这三种是默认已经注册过的，所以可以无需设置
    orm.RegisterDriver(SQLLITE3_DRIVER,orm.DRSqlite)

    // 参数1        数据库的别名，用来在 ORM 中切换数据库使用
    // 参数2        driverName
    // 参数3        对应的链接字符串
    orm.RegisterDataBase(DB_NAME,SQLLITE3_DRIVER,CONNECTION_INFO)

    //根据数据库的别名，设置数据库的最大空闲连接
    orm.SetMaxIdleConns(DB_NAME,30)
    //根据数据库的别名，设置数据库的最大数据库连接 (go >= 1.2)
    orm.SetMaxOpenConns(DB_NAME,120)

    //设置时区
    //orm.DefaultTimeLoc = time.UTC
}
