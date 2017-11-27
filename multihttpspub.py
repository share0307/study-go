#/usr/bin/env python3.6

import os
import time;
import json;
import requests;

class MultiClone(object):
    #token
    __token = "stzSxv1zzaSfipx4CJL4"

    #计数器
    __num = 0;

    #构造方法
    def __init__(self,path):
        #检查目录是否存在
        self.checkDirExists(path);
        #切换到该目录
        os.chdir(path)
        #设置目录
        self.__setDir(path)

    #遍历项目，clone 代码
    def cloneProjects(self):
        project_list = self.getProjectList()
        for i in project_list:
            self.initProject(i)

    #init 单个项目
    def initProject(self,info):
        print(time.strftime("%Y-%m-%d %H:%M:%S"));
        print("开始 init 第 ",info['id'],"(",self.__num,")","个项目 : ",info['ssh_url_to_repo'])
        self.__num = self.__num + 1;
        #完整的目录名称
        full_path = "%s/%s" % (self.__path,info['path'])
        print("当前目录为:",full_path)
        try:
            self.checkDirExists(full_path)
        except:
            #创建目录
            os.mkdir(full_path)

        #切换目录
        os.chdir(full_path)
        command_str = "git init"
        print(command_str)
        os.system(command_str)
        #构建关系
        command_str2 = "git remote add origin %s" % info['http_url_to_repo']
        print(command_str2)
        os.system(command_str2)
        #执行一下 pull
        command_str3 = "git pull"
        print(command_str3)
        os.system(command_str3)


    #从远端拉取git 项目列表
    def getProjectList(self):

        project_list_api = "http://git.sscf123.com/api/v3/projects";
        total_page = 10
        per_page = 100
        tmp = []
        for i in range (1,total_page):
            param = {}
            param['private_token'] = self.__token
            param['page'] = i
            param['per_page'] = per_page
            print(project_list_api,param);
            req = requests.get(project_list_api,param)
            tmp.extend(json.loads(req.text))
        # print(len(tmp))
        # exit()
        return tmp

    #判断目录是否存在
    def checkDirExists(self,path):
        if not os.path.exists(path):
            raise "目录不存在!"

    #设置目录
    def __setDir(self,path):
        print("当前目录:",path)
        self.__path = path


if __name__=='__main__':
    #path = "/usr/local/var/www/sscf";
    path = "/Users/kwin/Desktop/www/sscf-new";
    multoclone = MultiClone(path)

    multoclone.cloneProjects()