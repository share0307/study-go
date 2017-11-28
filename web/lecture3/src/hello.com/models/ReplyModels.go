package models

import (
	"strconv"
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

func AddReply(tid,nickname,content string) error {
	//数据转换
	topic_id,_ := strconv.ParseInt(tid,10,64)

	comment := &Comment{
		Tid:topic_id,
		Name:nickname,
		Content:content,
		Created:time.Now(),
	}

	orm := orm.NewOrm()

	_,err := orm.Insert(comment)

	if err != nil{
		beego.Error("add comment error:",err)
	}

	return nil
}

func DeleteReply(rid string) error {
	reply_id,_ := strconv.ParseInt(rid,10,64)

	orm := orm.NewOrm()

	//qs := orm.QueryTable("comment")

	comment := &Comment{
		Id:reply_id,
	}

	_,err := orm.Delete(comment)

	if err != nil{
		return err
	}

	return nil
}


/**
获取所有的评论
@author	jianwei
 */
func AllComments() (*[]Comment,error) {
	orm := orm.NewOrm()

	comments := new([]Comment)

	qs := orm.QueryTable("comment")
	num,err := qs.All(comments)

	if err !=nil{
		beego.Error("find comments error:",num,err)
	}

	return comments,nil
}