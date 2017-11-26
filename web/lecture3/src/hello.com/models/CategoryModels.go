package models

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "time"
    "fmt"
    "strconv"
)

/*
添加分类操作
 */
func AddCategory(title string) error {

    //检查同名分类是否存在
    if CheckCategoryExists(title) == true {
        beego.Informational("标题为:", title, "的分类已存在!")
        return nil
    }

    t := time.Now()
    beego.Warning("time.Time:",t)
    //分类对象指针
    category := new(Category)
    category.Title = title
    category.Created = t
    category.TopicTime = t

    beego.Error(category)

    fmt.Println("category:",category)

    //添加分类
    orm := orm.NewOrm()
    insert, err := orm.Insert(category)

    beego.Warning(insert)

    if err != nil {
        beego.Warning("插入失败!",err)
        return err
    }

    return nil
}

/**
通过名称查看是否有同名分类
@author jianwei
 */
func CheckCategoryExists(title string) bool {
    //创建一个 orm 对象
    orm := orm.NewOrm()

    //实例化 Category 结构体指针
    //最终也会把数据通过指针传递的特性传赋予给此值
    category := new(Category)

    //此处是传入的是需要操作的表，返回 QuerySeter
    qs := orm.QueryTable(category)

    //过滤条件
    qs = qs.Filter("title", title)

    //查询一条
    //最终也会把数据通过指针传递的特性传赋予给此值
    err := qs.One(category)

    if err != nil {
        beego.Warning("one category err:", err)
        return false
    }

    return true
}

/**
获取所有的分类列表
@author jianwei
 */
func AllCategoryList() (*[]Category,error) {
    orm := orm.NewOrm()

    category := new([]Category)

    //qs := orm.QueryTable(category)
    qs := orm.QueryTable("category")
    _,err := qs.All(category)

    if err != nil{
        return nil,err
    }

    return category,nil
}


/*
删除分类操作
@author jianwei
 */
func DelCategory(id string) error {

    //cid,e := strconv.Atoi(id)
    cid,e := strconv.ParseInt(id,10,32)
    if e != nil{
        panic(e)
    }

    orm := orm.NewOrm()

    //category := &Category{}

    qs := orm.QueryTable("category")
    qs = qs.Filter("id",cid)
    del,err := qs.Delete()

    beego.Info("delete:",del)
    if err != nil{
        return err
    }

    return nil
}