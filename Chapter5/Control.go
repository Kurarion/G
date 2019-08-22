package chapter5

import (
	"fmt"
	"runtime"
)

//Go的控制语句有：
//if else 结构
//switch 结构
//select 结构 ， 用于channel的选择
//可以使用迭代器或循环结构来重复执行一次或多次代码（任务）
//for(range)结构
//break continue可以用于中途改变循环的状态
//也有goto和标签
//注意：GO完全省略了if switch和for 结构中条件语句两侧的括号

//if是用于测试某个条件（bool或逻辑型）的语句
//如果条件成立，则会执行if后由大括号括起来的代码块
func TestIf(){
	//如果存在第二个分支，则可以在上面代码的基础上添加else关键字以及另一块代码块
	if true{

	}else if false{

	}else{

	}
	//以上代码不会出错
	//注意代码块是不能省略大括号的
	//而且当出现else if一起时，一定要在同一行
	//注意：else 一定要在}同一行
	//!(condition) 此时的括号一般情况下是需要的 例如： !(a == b)

	//当if结构内有break, continue goto 或 return语句时Go代码常见的写法是
	//省略else部分

	//判断一个字符串是否为空
	// str == ""
	// len(str) == 0
	//判断运行Go程序的操作系统类型，可以通过常量runtime.GOOS
	if runtime.GOOS=="windows"{
		fmt.Println("Win")
	}else{
		//Unix-like
		fmt.Println("Other")
	}

	//if可以包含一个初始化语句（给一个变量赋值）
	//if initialization; condition{}
	//注意：只能有一个初始化语句，可以使用同时赋值
	if a,b:=4,6;a<5&&b>5{

	}
	//一个很好使用情景
	//if vale := process(data); value > max {}
}

//通过使用常量GOOS可以在init函数中进行不同的初始化
func init(){
	if runtime.GOOS == "windows"{
		fmt.Println("Windows: Ctrl+Z")
	}else {
		fmt.Println("Unix-like: Ctrl+D")
	}
}

//相比较C与Java等其它语言，GO语言中的switch结构使用上更加灵活
//它接受任意形式的表达式
