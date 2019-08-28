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
//switch var1{
//	case val1:
//	 	...
//	case val2:
//		...
//}
//变量var1可以是任何类型，而val1则可以是同类型的任意值，类型不被局限于常量或整数
//但必须是相同的类型，或者最终结果为相同的类型的表达式，注意：前花括号必须与switch同一行

//可以同时测试多个可能符合条件的值，使用逗号分割它们，例如： case val1,val2,val3
//每一个case分支都是唯一的，从上到下逐一测试，直到匹配为止，（）
//Go语言使用快速的查找算法来测试switch条件与case分支的匹配情况
//直到算法匹配到某个case或进入default条件为止
//
//注意：一旦成功地匹配到某个分支，在执行完相应代码后就会退出整个switch代码块，也就是
//不需要特别使用break语句来表示结束
//
//如果在执行完每个分支的代码后还希望继续执行后续分支的代码，可以使用fallthrough关键字来达到目的
//
func TestSwitch(){
	i:=3
	switch i {
	case 1,2:
		fmt.Println("1,2")
	case 3:
		fmt.Println("3")
		fallthrough
	case 4:{
		fmt.Println("4")
		//fallthrough error 报错：out of place
		//break ok 正确
	}
	fallthrough
	case 5:
		r:=5
		fmt.Println("5",r)
		fallthrough
	default:
		fmt.Println("default")
	}
	//fallthrough只对相邻的下一个case或default有效
	//如果希望后面的case完全fallthrough，则需要每个case后面加上fallthrough关键字
	//Go中的Switch中的case可以使用{}组织
	//注意：在{}中不能使用fallthrough关键字
	//但可以使用break

	//关于变量1的声明，在C++中switch的case内不能直接声明定义一个变量，需要使用{}进行作用域的限定才可
	//如果仅仅进行变量的声明但却没有进行初始化则不会报错，C++
	// 以下为C++代码
	//		//std::string name; //错误 控制流绕过一个隐式初始化值
	//		//int cc = 1; //错误 控制流绕过一个显示初始化值
	//		{
	//			std::string name = "ddddd"; //定义在一个块中
	//		}
	//		int c; //正确 因为c没有被初始化
	//其中c没有被初始化的原因是C++在非全局变量中声明内置类型时不会自动初始化
	//而std::string会隐式调用默认构造函数进行初始化
	//因为控制流可能会跳过此case，如果没有作用域限定，则此声明与初始化在switch
	//而case的限定才合情合理，因此必须使用{}进行变量的声明与初始

	//不同于C++，可能由于Go有自己的垃圾回收处理机制，case中完全不需要{}，但也可以使用
	//可选的default分支可以出现在任何顺序，但最好将它放在最后

	//==========
	//switch的第二种的形式是不提供任何被判断的值，实际上是默认判断是否为true
	//然后在每个case分支中进行测试不同的条件，当任一分支的测试结果为true时，该分支的代码会被
	//执行，形似链式的if-else语句，在测试条件非常多的情况下，提供了可读性的书写方式
	//
	j := 5
	switch {
	case j<5:
		fmt.Println("<5")
	case j==5:
		fmt.Println(" == 5")
	case i>5,j>5://注意只关心表达式的最后结果,同样可以使用,进行多个表达式测试
	default:
		fmt.Println("Default")
	}

	//=========
	//switch的第三种形态是包含一个初始化语句
	//可以非常优雅地进行条件判断
	//switch  result:= calc() {
	//case result < 0:
	//	...
	//case result > 0:
	//	...
	//default:
	//	...
	//}
	//当然也可以使用平行初始化
	//switch a, b := x[i], y[j] {
	//	case a < b: t = -1
	//	case a == b: t = 0
	//	case a > b: t = 1
	//}

	//此外switch语句还可以初始用于type-switch来判断某个interface变量中实际存储的变量类型
	//
}

//For结构：
//如果想要重复执行某些语句，Go语言中只有for结构可以使用
//注意：Go中是没有while语句的
func TestFor(){
	//======
	//基于计数器的迭代，基本的形式
	//for 初始化语句; 条件语句; 修饰语句 {}
	for i := 0; i <5; i++{
		fmt.Print("口")
	}
	fmt.Println()
	//注意：这三部分组成的循环的头部，之间使用分号;相隔，但并不需要括号将它们括起来
	//例如： for(i = 0; i<10; i++){} //错误
	//同样，左花括号需要和for语句在同一行，计数器的生命周期在遇到右花括号就结
	//而且可以在循环中同时使用多个计数器：
	for i,j := 0,10 ; i<j ; i, j = i+1, j-1{
		fmt.Println("O")
	}//主要得益于Go语言具有的平等赋值的特性
	//for可以相互嵌套

	//使用Goto重写循环
	//注意： 标签后面的语句就是循环体！
	//注意这与Java似乎有点不同
	i,j:=0,0
	loop:
	j++
	i++
	//注意：使用'-'会默认输出整型(%v)，使用"-"
	fmt.Print(j,"-")
	if i<15 {
		goto loop
	}
	//如果三部分其中想要省略，同C++一样直接使用;;空语句即可


	//========
	//基于条件判断的迭代
	//没有头部的条件判断迭代（类似于while循环）
	//基本形式: for 条件语句 {}
	 cc := 5
	 for cc > 0 {
	 	//for ;cc > 0; { //实际上也可以使用第一形态但只使用条件判断语句，其他使用;相隔的空语句
			 //do some things
			 cc--
			 //}
	 }

	 //===========
	 //无限循环
	 //形式 for{}
	 //条件语句是可以被省略的，如i := 0;; i++ 或 for ;; {} ;;会在使用gofmt时被
	 //移除，这些循环的本质就是无限循环 最后一个形式也可以被改写为 for true {}
	 //但一般情况下直接写 for {}
	 //注意：跳出循环的方式是return 或 break
	 //无限循环的经典应用是服务器，用于不断等待和接受新的请求

	 //===========
	 //for-range结构
	 //这是Go特有的一种迭代结构（C++11针对迭代器对象有类似语法： for(auto c:map),但不能获得索引）
	 //语法类似于其他语言的foreach，一般形式为
	 //for ix, val = range coll {} //ix 为索引值， val为数组中值的拷贝
	 //需要注意的是，val始终为集合中对应索引的值拷贝，因此一般只具有只读性质，
	 //对它所做的任何修改都不会影响到集合中的原有的值
	 //注意：排除原来数组中的值为指针,一个字符串是Unicode编码的字符串数组，因此可以使用range
	 //

	 var ss string
	 ss = "aabbccddee"

	 for pos, char := range ss {
	 	//经测试不能使用pos, &char = range ss
	 	//除非数组本身是一个指针的索引
	 	fmt.Printf("%02d位置的字符是%c",pos,char)
	 	fmt.Println()
	 }

	 //针对一个rune数组
	 sss := "这是一个rune数组"
	 fmt.Println("========")
	 for pos, val := range sss{
	 	//注意：range是会自动解析一个unicode编码的字符串的，从而进行正确的迭代
	 	//它能够自动根据 UTF-8 规则识别 Unicode 编码的字符。
	 	fmt.Printf("%02d位置的字符是%c",pos,val)
	 	fmt.Println()
	 }
	fmt.Println()
	fmt.Println("index int(rune) rune    char bytes")
	for index, rune := range sss {
		fmt.Printf("%-2d      %d      %U '%c' % X  % X\n", index, rune, rune, rune, []byte(string(rune)), rune)
	}
	//注意这里
	//[]byte(string(rune)) 而不能直接使用 rune
	//rune是一个码点（使用int32的别名），可能由于多个代码单元组成，因此本质是一个代码单元的数组表达转换来的int32
	//首先rune本质是一个整型类型，先通过string强制转换后成为一个string代码单元的数组
	//string可以使用码点来生成一个字符串，包括unicode
	//[]byte()最后转换string类型，其实不用强制转换结果也一样
	//


}

//break 与 continue

func TestBTC(){
	i := 1
	for{
		i+=1;
		if i==10 {
			break
		}
	}
	//使用Break可以退出循环，一个break的作用范围是该语句出现后的最内部的结构
	//可以被利用于任何形式的for循环（计数器，条件判断等）
	//但在switch或select语句中，break语句的作用结果是跳过整个代码块执行后续的代码
	//注意：break只会退出最内层的循环

	//关键字continue忽略剩余的循环体而直接进入下一次的循环的过程
	//但不是无条件执行下一次循环，执行之前依旧需要满足循环的判断条件

}

//标签与goto

func TestGoto(){
	//for switch select语句都可以配合标签label形式的标识符使用
	//即某一行第一个以冒号结尾的单词（gofmt会将后续代码移至下一行）
	//注意：Label的名称是大小写敏感的，为了提升可读性，一般使用全部大写字母
	i := 64
	LABEL2:
		i++
		fmt.Printf("%c",i)
	if i < 99 {
		goto LABEL2
	}

	LABEL1:
	for i := 0; i <= 5; i++ {
		LABEL3:
		for j := 0; j <= 5; j++ {
			if j==3 {
				continue LABEL3
			}
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
	//for i:=1;i<2;i++{
	//	continue LABEL1
	//}


	//注意：这里是continue Label 而不是goto Label
	//如果是continue Label而不会进行i:=0的重声明与赋值
	//但是continue Label只能在Label后面紧跟着for
	//而且Label内只有一个for并且这个for内部有continue Label...
	//但可以多个循环使用多个LABEL并同时相互嵌套
	//
	//外部的循环不能使用continue其它LABEL范围的
	//
	//结果是i==5&&j==4是不会打印的
	//但正常来说冒号后面的就是循环体
	//对于有初始化的for语句来说，实际情况是紧跟在标签后面
	//上面这个例子，如果使用goto替换continue则会成为无限循环

	//如果一定需要使用goto，应当使用正序的标签
	//即标签位于goto语句之后
	//注意：标签和goto语句之间不能出现定义新变量的语句

}
