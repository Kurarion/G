package Chapter6

//每个程序都包含很多的函数
//Go是编译型语言，所以函数编写的顺序是无关紧要的
//DRY原则，即不要重复类型行为的代码

//return语句可以用于结束for或结束一个协程goroutine
//Go里面有三种类型的函数
//1 普通的带有名字的函数
//2 匿名函数或lambda函数
//3 方法 Methods

//除了main() init()函数外，其它所有类型的函数都可以有参数与返回值
//函数参数，返回值以及它们的类型被称为函数签名

//函数可以将其他函数作为它的参数，只要函数的返回值个数，类型，顺序与调用函数的形参一致即可

//函数重载(function overloading)
//指的是可以编写多个同名函数，只要它们拥有不同的形参与/或者不同的返回值
//注意：在Go语言中函数重载是不被允许的
//Java 支持函数的重载但不支持运算符的重载
//C++ 以上都支持

//GO不支持的原因是函数重载需要进行多余的类型匹配影响性能
//没有重载意味着只是一个简单的函数调度

//注意：函数也可以以申明的方式被使用，作为一个函数的类型
//type binOp func(int, int) int
//这里不需要函数体{}
//函数是一等值(first-class value)它们可以赋值给变量
//就像add := binOp 一样
//这个变量知道自己指向的函数的签名，所以给它赋一个具有不同签名的
//函数值是不可能的

//函数值(functions value)之间可以相互比较，如果它们引用的是相同的函数或都是Nil
//的话，则认为它们是相同的函数，函数不能在其它函数里面声明（即不能嵌套）
//但可以通过使用匿名函数来破除这个限制
//
//目前Go没有泛型（generic）的概念，也就是说它不支持那种支持多种类型的函数
//不过在大部分情况下可以通过接口(interface)特别是空接口与类型选择
//switch与/或者通过使用反射(reflection)来实现相似的功能
//但性能会变低，所以在非常注意性能的场合下，最好是为每一个类型单独创建一个函数

//函数的参数与返回值