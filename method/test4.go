package main;

import "fmt";

//类型不能不得直接使用 原类型的方法
type T struct {
    Name string
}


//实际上(t *T)也可以看作是方法的第一个参数，或许是为了区分函数或者方法的区别吧
//可以直接使用 接收者.方法(接收者类型的变量) (*T).test(&t) ，把接收者的对象作为第一个参数即可调用该方法
func (t *T)test(name string)  {
    if len(name) > 0{
        t.Name = name;
    }
    fmt.Println(t);
}

func main()  {
    var t T=T{Name:"kkk"};
    //直接使用变量调用方法
    t.test("");

    //实际上(t *T)也可以看作是方法的第一个参数，或许是为了区分函数或者方法的区别吧
    //可以直接使用 接收者.方法(接收者类型的变量) (*T).test(&t) ，把接收者的对象作为第一个参数即可调用该方法
    (*T).test(&t,"sb");


    fmt.Println("Hello World!");
}

