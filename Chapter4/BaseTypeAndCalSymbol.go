package chapter4

import "fmt"

//表达式是一种特定的类型的值
//它可以由其他的值以及运算符组合而成
//每个类型定义了可以和自己结合的运算符的集合
//如果使用了不在这个集合中的运算符，则编译错误

var test_same_name = 1

func TestSameName()  {
	fmt.Println("\n===TestSameName===")
	test_same_name:="abc"
	{
		fmt.Println(test_same_name)
	}
	fmt.Println(test_same_name)
	//暂时不清楚如何访问被屏蔽的变量，也可能没有类似C++中的::全局作用域的语法
}

//一元运算符只可以用于一个值的操作（作为后缀）而二元运算符则可以和两个值或者操作数结合
//注意：Go是强类型语言，因此不会进行隐式转换，在任何不同类型之间的转换都必须显式说明
//因此只有两个类型相同的值才能和二元运算符结合
//Go不存在像C++那样的运算符重载（Java也不支持）表达式的解析顺序是从左到右
//可以通过括号来提高部分的优先级

//布尔类型bool
//var b bool = true // var b = true
//布尔型的值只可以是常量true和false
//两个类型相同的值可以使用==或!=运算符进行比较并获得一个布尔型的值
//
//Go对于值之间的比较有非常严格的
