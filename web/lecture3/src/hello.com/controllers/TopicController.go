package controllers

import (
	"github.com/astaxie/beego"
    "hello.com/models"
    "regexp"
	"strconv"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {
	c.Data["Title"] = "文章 - 我的 beego 博客"
	c.Data["IsTopic"] = true
	c.TplName = "topic/home.html"

	condition := make(map[string]string)
	condition["sort"] = "-created"
	topics,_ := models.TopicList(c.Input().Get("title"),condition)

	c.Data["Topics"] = topics
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	ck,err := c.Ctx.Request.Cookie("username")
	if err == nil{
		c.Data["Username"] = ck.Value
	}
}

/**
添加文章
@author jianwei
 */
func (this *TopicController)Add()  {
    this.Data["Title"] = "添加文章 - 我的 beego 博客"
    this.Data["IsLogin"] = checkAccount(this.Ctx)
    ck,err := this.Ctx.Request.Cookie("username")
    if err == nil{
        this.Data["Username"] = ck.Value
    }

    //获取分类列表
    this.Data["CategoryList"],_ = models.AllCategoryList()

    this.TplPrefix = "topic/"
    this.TplName = "add.html"
    return
}

/**
添加或者新增文章
@author jianwei
 */
func (this *TopicController)Post()  {
    if !checkAccount(this.Ctx){
        this.Redirect("/user/login",302)
        return
    }

    tid := this.Input().Get("tid")
    title := this.Input().Get("title")
    content := this.Input().Get("content")
    category_id,_ := strconv.ParseInt(this.Input().Get("category_id"),10,64)


    var err error
    if len(tid) >0{
        err = models.UpdateTopic(tid,category_id,title, content)
    }else {
        err = models.AddTopic(category_id, title, content)
    }

    beego.Informational(err)
    if err != nil{
        beego.Error(err)
    }

    this.Redirect("/topic",302)
    return
}


/**
文章详情
@author jianwei
 */
func (this *TopicController)View()  {

    tmp := this.Ctx.Input.Params()["0"]

    var topicid string

    //生成返回一个 Regexp 对象
    reg := regexp.MustCompile(`([\d]+)[\.html|\.html]+`)
    //查找并且获得字符串
    var tmpstr []string
    tmpstr = reg.FindStringSubmatch(tmp)
    beego.Debug("tmpstr:",tmpstr)


    for k,_ := range tmpstr{
        if k == 1 {
            topicid = tmpstr[1]
        }
    }


    //beego.Debug("tmp----------:",tmp)
    //beego.Debug("topicid----------:",topicid)

    //beego.Error("topicid:",topicid)

    topic,err := models.GetTopicDetails(topicid)

    if err != nil{
        beego.Error(err)
        this.Redirect("/blog",302)
        return
    }

    this.Data["Topic"] = topic

	this.Data["Replies"],_ = models.AllComments()
	this.Data["IsLogin"] = checkAccount(this.Ctx)

    this.TplPrefix = "topic/"
    this.TplName = "view.html"
    return
}

func (this *TopicController)Modify()  {

    tid := this.Input().Get("tid")

    topic,err := models.GetTopicDetails(tid)

    if err != nil{
        this.Redirect(this.Ctx.Request.Referer(),302)
    }

	//获取分类列表
	this.Data["CategoryList"],_ = models.AllCategoryList()

    this.TplName = "topic/modify.html"
    this.Data["Topic"] = topic
}

/**
删除文章
@author jianwei
 */
func (this *TopicController)Delete()  {
    if !checkAccount(this.Ctx){
        this.Redirect("/user/login",302)
        return
    }

    err := models.DeleteTopic(this.Input().Get("tid"))

    if err != nil{
        this.Redirect(this.Ctx.Request.Referer(),302)
    }

    this.Redirect("/topic",302)

    return
}