package controllers

import (
	"github.com/astaxie/beego"
	//"hello.com/models"
	"hello.com/models"
	"fmt"
)

type ReplyController struct {
	beego.Controller
}

/*
添加评论
@author	jianwei
 */
func (this *ReplyController)Add()  {
	tid := this.Input().Get("tid")
	nickname := this.Input().Get("nickname")
	content := this.Input().Get("content")

	err := models.AddReply(tid,nickname,content)

	if err != nil{
		this.Redirect(this.Ctx.Request.Referer(),302)
	}
	redirectUrl := fmt.Sprintf("http://%s/topic/view/%s.html",this.Ctx.Request.Host,tid);
	beego.Informational("到这里来了啊！！")
	beego.Informational(redirectUrl)

	this.Redirect(redirectUrl,302)

	return
}

/**
删除文章
@author	jianwei
 */
func (this *ReplyController)Delete()  {
	rid := this.Input().Get("rid")

	models.DeleteReply(rid)

	this.Redirect(this.Ctx.Request.Referer(),302)

	return
}