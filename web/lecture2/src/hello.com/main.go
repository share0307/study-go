package main

import (
	_ "hello.com/routers"
	"github.com/astaxie/beego"
	"hello.com/models"
	"github.com/astaxie/beego/orm"
)

func init() {
	//调用数据库初始化
	models.RegisterDB()
}

func main() {
	// 开启 orm 调试模式，可以打印 sql 语句
	orm.Debug = true
	//设置数据库初始化

	//用于初始化数据库
	//参数一：需要初始化的数据库别
	//参数二：是否强制建表，比如说当数据库存在，将把之前的库干掉重新建表
	//参数三：打印更多调试信息
	orm.RunSyncdb(models.DB_NAME,false,true)


	beego.Run()
}

