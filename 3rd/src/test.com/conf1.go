package main

import (
	"github.com/Unknwon/goconfig"
	"log"
	"fmt"
)

func main()  {
	//加载配置文件
	cfg,err := goconfig.LoadConfigFile("config/config.ini")

	if err != nil{
		log.Fatalf("无法加载配置文件：%s",err)
	}

	//fmt.Println(cfg)

	//cfg.GetValue(分区名称(没有分区则使用goconfig.DEFAULT_SECTION作为默认分区名称,使用空字符串""也是被认为是默认分区),"配置键名")
	value,err := cfg.GetValue(goconfig.DEFAULT_SECTION,"key_default")

	if err != nil{
		log.Fatalf("无法获取键值(%s)：%s","key_default",err)
	}


	log.Printf("%s > %s:%s",goconfig.DEFAULT_SECTION,"key_default",value)


	isInsert := cfg.SetValue(goconfig.DEFAULT_SECTION,"key_default","这是新的值！")

	if err != nil{
		log.Fatalf("无法设置键值(%s)：%s","key_default",err)
	}

	log.Printf("设置键值 %s 为插入操作：%v","key_default",isInsert)

	//注释
	//comment := cfg.GetSectionComments("super")

	log.Fatalf("无法设置键值(%s)：%s","key_default",err)











	fmt.Println()

}