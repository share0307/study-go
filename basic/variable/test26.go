package main

import (
    "fmt"
    //"sort"
)

func main() {
    m1 := map[int]string{1:"a",2:"b",3:"c",4:"d",5:"e",6:"f"};

    //m2 := map[string]int{};
    m2 := make(map[string]int);

    fmt.Println(m1);

    for k,v := range m1{
        m2[v] = k;
    }

    fmt.Println(m2);

    fmt.Println("Hello world!");
}
//
//
//list_alldir(){
//for file2 in `ls -a $1`
//do
//if [ x"$file2" != x"." -a x"$file2" != x".." ];then
//if [ -d "$1/$file2" ];then
//#echo "/Users/kwin/Desktop/www/sscf/$file2";
//cd "/Users/kwin/Desktop/www/sscf/$file2"
//pwd
//git pull
//git checkout master
//git pull
//git checkout slave
//git pull
//#list_alldir "$1/$file2"
//cd "/Users/kwin/Desktop/www/sscf/"
//fi
//fi
//done
//}