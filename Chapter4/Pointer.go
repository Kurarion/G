package chapter4

import "fmt"

//不像Java与.Net，Go为程序员提供了控制数据结构的指针的能力
//但不能进行指针运算，通过给予程序员基本内存布局，Go语言允许控制
//特定集合的数据结构，分配的数量以及内存访问模式
//
//每个内存块（或字）有一个地址，通常使用十六进制数表示
//
//Go语言的取地址符是& 放到一个变量前使用就会返回相应变量的内存地址
//
func TestPointerPrint(){
	a := 5
	fmt.Printf("Int: %d, Ptr: %v\n",a,&a)

	//这个地址可以存储在一个叫做指针的特殊数据类型中，在本例中这是一个指向int的指针
	//可以声明一个int类型的指针
	var intP *int
	intP = &a
	fmt.Println(intP)
	//一个指针变量可以指向任何一个值的内存地址，它指向的那个值的内存地址
	//在32位机上占用4个字节 64位机上占用8个字节
	//使用一个指针引用一个值为间接引用，
	//反引用，指针转移，或叫解指针同样使用*与C++几乎一样的语法

	//当一个指针被定义没有分配任何变量时，值为nil
	//一个指针变量通常编写为ptr

	//注意：在书写表达式类似 var p *type时，一定要将*与变量名用一个空格分开
	//虽然 p*type语法是正确识别的，但容易被认为是一个乘法表达式
	//
	//符号*可以放在指针类型上，C++中叫做二级指针
	//
	var test_multi_ptr **int
	test_multi_ptr = &intP
	fmt.Println(test_multi_ptr,*test_multi_ptr,**test_multi_ptr)

	s := "good bye"
	var p *string = &s
	*p = "ciao"
	fmt.Printf("Here is the pointer p: %p\n", p) // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s) // prints same string

	//注意：不能得到一个文字或常量的地址，例如
	//const i = 5
	//ptr := &i
	//ptr2 := &10
	//
	//Go语言和C++这些语言一样，但是对于经常导致C语言内存泄漏从而导致程序崩溃的指针运算
	//（即指 ptr++ 这样操作）是不被允许的，Go为了保证内存安全，更像是Java中的引用

	//注意：Go不支持指针本地的（偏移）运算
	//指针的高级应用是可以传递变量使用（参数），指针传递是廉价的，提高效率
	//
	//另一方面，由于指针导致的间接引用（一个进程执行了另一个地址），指针的过度频繁使用可以会导致性能下降
	//
	//指针也可以指向另一个指针，并且可以进行任意深度的嵌套，即多级指针（多级间接引用）
	//注意：Go语言为了方便程序员，隐藏间接引用，如：自动反向引用
	//
	//对于一个空指针的反向引用是不合法的，并且会使用程序崩溃
	//注意：相当于C++的野指针，为了防止危险操作
	//指针使用前一定要正确地初始化
	//var p *int = nil
	//*p = 0
}
