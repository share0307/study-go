package models

import (
    "github.com/astaxie/beego/orm"
    "time"
    "errors"
    "github.com/astaxie/beego"
    //"reflect"
    //"strconv"
    "strconv"
)

/*
添加分类操作
 */
func AddTopic(title,content string) error {
    if len(title) < 1 || len(content) < 1{
        return errors.New("请输入标题或者文章！")
    }

    orm := orm.NewOrm()

    topic := &Topic{
        Title:title,
        Content:content,
        Created:time.Now(),
        Updated:time.Now(),
        ReplyTime:time.Now(),
    }

    _,err := orm.Insert(topic)

    if err != nil{
        return err
    }

    return nil
}

/*
修改分类操作
 */
func UpdateTopic(tid interface{},title,content string) error {
    if len(title) < 1 || len(content) < 1{
        return errors.New("请输入标题或者文章！")
    }
    var topicId int64
    switch v := tid.(type) {
    case int64:
        topicId = v
    case string:
        topicId,_ = strconv.ParseInt(v,10,64)
    case int32:
        topicId = int64(v)
    }

    orm := orm.NewOrm()

    topic := &Topic{
        Id:topicId,
        Title:title,
        Content:content,
        //Created:time.Now(),
        Updated:time.Now(),
        //ReplyTime:time.Now(),
    }

    _,err := orm.Update(topic,"Id","Title","Content","Updated")

    if err != nil{
        return err
    }

    return nil
}

/**
查找文章列表
@author jianwei
 */
func TopicList(title string,condition map[string]string) ([]Topic,error) {
    orm := orm.NewOrm()

    topics := new([]Topic)

    qs := orm.QueryTable("topic")
    if len(title) > 0 {
        qs = qs.Filter("title__icontains", title)
    }

    var ok bool
    if _,ok = condition["sort"];ok == true{
        beego.Error("here!")
        beego.Error(condition["sort"])
        qs = qs.OrderBy(condition["sort"])
    }

    //beego.Error("title:",title)
    l,err := qs.All(topics)

    beego.Warning("len:",l)

    if err != nil{
        beego.Error(err)
        return nil,err
    }

    return *topics,nil
}

/**
通过 id 获取 一片文章详情
@author jianwei
 */
func GetTopicDetails(id interface{}) (*Topic,error) {

    //t := reflect.TypeOf(id)
    var topicid int64
    switch v := id.(type) {
    case string:
        //beego.Warning("string:",v)
        topicid,_ = strconv.ParseInt(v,10,64)
    case int32:
        //beego.Warning("int32:",v)
        topicid = int64(v)
    case int64:
        //beego.Warning("int64:",v)
        topicid = int64(v)
    default:
        panic("数据类型错误！")
    }


    //beego.Error(reflect.ValueOf(id).Kind())
    //beego.Error("type:",t.Kind())

    beego.Informational("topicid:",topicid)

    //取得 topic 结构指针
    topic := new(Topic)

    //取得 orm 对象
    orm := orm.NewOrm()
    qs := orm.QueryTable("topic")
    qs = qs.Filter("id",topicid)
    err := qs.One(topic)

    if err != nil{
        return nil,err
    }

    //观看次数递增1
    topic.Views++
    beego.Informational("topic.Id:",topic.Id)
    beego.Informational("topicid:",topicid)
    _,err = orm.Update(topic,"views")

    if err != nil{
        return nil,err
    }

    return topic,nil
}

/**
删除文章
@author jianwei
 */
func DeleteTopic(tid interface{}) error {
    var topicId int64
    switch v := tid.(type) {
    case int64:
        topicId = v
    case string:
        topicId,_ = strconv.ParseInt(v,10,64)
    case int32:
        topicId = int64(v)
    }

    orm := orm.NewOrm()
    _,err := orm.Delete(&Topic{Id:topicId})

    beego.Error("delete error:",err)

    if err != nil{
        return err
    }

    return nil
}

