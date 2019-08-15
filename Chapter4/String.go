package chapter4

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

//字符串是UTF-8字符的一个序列
//当字符是ASCII码时占用一个字节，其他字符根据需要占用2-4个字节
//UTF-8是被广泛使用的编码格式，包括xml JSON
//与C++ JAVA Python不同，Java始终使用2个字节
//GO不仅减少内存与硬盘的空间占用，并且不需要对UTF-8进行编码解码
//
//字符串是一种值类型，且值不可变，即创建某个文本后无法再次修改这个文本的内容
//这与Java有点类似，字符串是字节的室长数组
//
//Go支持以下2种形式的字面值
//1。解释字符串
//该类字符串使用双引号括起来，其中的相关的转义字符将被替换
//转义字符包括
//\n：换行符
//\r：回车符
//\t：tab 键
//\u 或 \U：Unicode 字符
//\\：反斜杠自身
//2。非解释字符串
//该类字符串使用反引号括起来，支持换行，例如
//`This is a raw string \n` 中的 `\n\` 会被原样输出。
//与C/C++不同，Go中的字符串是根据长度限定的，而不是特殊字符\0 [这是指的是C风格的字符串]
//string类型的零值是长度为0的字符串，即空字符串""
//一般的比较运算符（== !=。。。）通过在内存中按字节比较来实现字符串的对比
//可以通过len()来获取字符串所占的字节长度：len(str)
//字符串的内容（纯字节）可以通过标准索引法来获取，在中括号[]内写入索引，从0开始
//字符串 str 的第 1 个字节：str[0]
//第 i 个字节：str[i - 1]
//最后 1 个字节：str[len(str)-1]
//注意：这种转换方案只对纯ASCII码有效
//获取字符中某个字节的地址的行为是非法的例如：&str[i]
//字符串拼接符+
//s:= s1+s2
//注意：由于编译器行尾自动补全分号的原因，加号必须放在第一行
//也可使用+=
//在循环中使用+号拼接字符串并不是最高效的方法
//更好的办法是你爱我函数 strings.Join()
//还有更好的方法，使用字节缓冲bytes.Buffer的拼接
//
//通过将字符串看作是字节(byte)的切片(slice)来实现对其西直门外索引法的操作

func ExConutCharaters(){
	var str string
	str =`asSASA ddd dsjkdsjsこん dk`
	fmt.Println("原数量是：",len(str))
	var str1 []byte
	str1 = []byte(str)
	n:=utf8.RuneCount(str1)
	fmt.Println("数量是：",n)
	//注意 Go是不支持隐式类型转换的，因此不能将一个字符串与整数相加
	//println多个参数是使用逗号来分割的
}

//作为一种基本数据类型，每种语言都有一些对于字符串的预定义处理函数，Go中使用strings包来完成
//对字符串的主要操作
//
func TestStrFunc(){
	str := "1234aabbccdd"
	//HasPrefix判断字符串是否以prefix开头
	res1:=strings.HasPrefix(str,"12")
	//HasSuffix判断结尾
	res2:=strings.HasSuffix(str,"cdd")
	//
	fmt.Println("前",res1," 后",res2)

	//字符串包含关系
	//Contains 判断字符串是否包含substr
	fmt.Println("包含",strings.Contains(str,"bbc"))



}