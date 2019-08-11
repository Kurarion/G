package chapter4

import "fmt"

//声明变量的一般形式是使用var 关键字
//var identifier type
//不同于const不能省略type
//同时不同于常量，可以先只声明而不初始化

//与许多语言不同，它在声明变量时将变量的类型放在变量的名称之后，
//这是为避免C那样含糊不清的声明方式，int* a,b 在这个 声明中
//只有a是指针，而b是一个int
//而在Go中可以都声明为指针类型
//var a,b *int
//也可使用因式分解的写法
//注意：一个包中不能有重名的全局变量，即使是一个包中的不同文件也不允许出现同名
var(
	aa int
	b bool
)
func notForError(){
	fmt.Print(aa,b)
}
//注意：不同于C++，与Java相同，当一个变量被声明后，系统自动赋予它该类型的零值
//int = 0 float = 0.0 bool = false
//string 为 空字符串
//指针为 nil
//记住，所有的内存在Go中都是经过初始化的
//变量的命名最好骆驼命名法

//一个变量在程序中都有一定性的作用范围
