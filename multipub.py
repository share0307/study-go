#/usr/bin/env python3.6

# import gittle
import os
#import git
import getpass;
import time;

class MultiPub(object):
    #存储目录
    __path = '';

    #构造方法
    def __init__(self,path):
        #检查目录是否存在
        self.checkDirExists(path);
        #切换到该目录
        os.chdir(path)
        #设置目录
        self.__setDir(path)

    #设置目录
    def __setDir(self,path):
        self.__path = path

    #判断目录是否存在
    def checkDirExists(self,path):
        if not os.path.exists(path):
            raise "目录不存在!"

    #到各个目录拉取代码从 git 拉取代码
    def gitMultiPub(self):
        #循环目录下的目录，然后
        list_dir = os.listdir(self.__path)
        for i in list_dir:
            #先获取完整路径
            full_path = os.path.abspath(i)
            if os.path.isdir(full_path):
                #切换目录
                os.chdir(full_path)
                #取得当前目录
                # print("当前目录：",full_path);
                self.gitPub(full_path)
                #切换回原目录
                os.chdir(self.__path)


    #拉取代码
    def gitPub(self,path):
        #直接拉取一下代码
        first_pull = os.system('git pull');
        print(first_pull)

        #拉取代码
        out = os.popen('git branch -a');
        branch_str = out.read()
        print(branch_str)
        branch_list = branch_str.splitlines()
        # branch_list = 'a|a|n'.splitlines()
        for i in branch_list:
            #取得当前目录
            print("当前目录：",path);

            branch = i;
            idx = branch.rfind('/')
            if(idx > 0 ):
                branch = branch[idx + 1:]
            branch = branch.strip("*| ")
            chk_command = 'git checkout '+branch;
            print(chk_command)
            os.system(chk_command)
            #拉取代码
            second_pull = 'git pull'
            os.system(second_pull)
            #切换回 master 分支
            os.system('git checkout master')
            print("branch:",branch,"拉取完毕!")



if __name__ == '__main__':
    print("开始执行程序!")
    print("当前执行用户为：",getpass.getuser())

    time.sleep(1);

    pub = MultiPub("/usr/local/var/www/sscf")
    pub.gitMultiPub();
