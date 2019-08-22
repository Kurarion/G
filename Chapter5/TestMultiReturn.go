package chapter5

import (
	"fmt"
	"strconv"
)

//测试多返回值函数的错误
//Go语言经常使用两个返回值来表示执行是否成功
//返回nil或零值 false表示错误
//当使用bool表示时也可以使用一个error类型的变量来代替作为第二个返回值
//成功执行的话error的值为Nil，否则就会包含相应的错误信息（）
//Go语言中的错误类型为error: var err error
//这样一来就需要使用if语句来测试执行的结果
//由于其符号的原因，这样的形式又叫做comma, ok模式(pattern)

func TestErrReturn(){
	//在strconv.Atoi函数的返回值中，第二个就是一个错误类型
	var orig string = "ABC"
	news := "100"


	if an,err:=strconv.Atoi(orig);err==nil{
		fmt.Println("Result:",an)
	}else{
		fmt.Println("Failed!")
	}

	//也可使用以下写法（推荐）
	var an int
	var err error
	//或an, err := strconv.Atoi(orig)在外面
	if an,err =strconv.Atoi(news);err!=nil{
		fmt.Println("Failed!")
		//发生问题中止
		return
	}
	//没有问题继续
	fmt.Println("Result:",an)

	//当函数的调用者是main函数时，使用return会结束程序
	//但当是其它时，如希望结束程序可以使用os包的Exit函数
	//os.Exit(1)
	//此外的退出代码1可以使用外部脚本获取

	//当字符串转换为整数并确定没有问题可以使用 res,_ = 进行封装
	//实际上fmt.Println也有两个返回值
	//打印的字符数与error
	count1, err:= fmt.Println("TestGO")
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("刚才打印了：",count1,"个字符")
}
