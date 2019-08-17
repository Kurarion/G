package chapter4

import (
	"fmt";
	"strconv"
	"strings"
	"unicode/utf8"
)
//注意 单个字符仍与C++一样使用单引号
//字符串是UTF-8字符的一个序列
//当字符是ASCII码时占用一个字节，其他字符根据需要占用2-4个字节
//UTF-8是被广泛使用的编码格式，包括xml JSON
//与C++ JAVA Python不同，Java始终使用2个字节
//GO不仅减少内存与硬盘的空间占用，并且不需要对UTF-8进行编码解码
//
//字符串是一种值类型，且值不可变，即创建某个文本后无法再次修改这个文本的内容
//这与Java有点类似，字符串是字节的数组
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
	s := "01大1234aabbccddeeff"
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

	//判断子字符串或字符在父字符串中出现的位置（索引）
	//Index返回字符串str在字符串s中的索引(str的第一个字符的索引)，-1表示字符串s中不包含字符串str
	fmt.Println("包含的位置",strings.Index(s,str)) //5
	fmt.Println("此位置的字符：", string(s[5]))

	//LastIndex返回str在s中最后出现位置的索引，-1不包含

	//如果要查询非ASCII编码的字符在父字符串中的位置，建议使用以下函数
	//注意是单个的rune字符
	//注意第二个参数是一个Rune类型
	fmt.Println("包含的位置",strings.IndexRune(s,rune('大'))) //2
	fmt.Println("此位置的字符：", string(s[2]))

	//字符串的替换
	//Replace用于将字符串str中的前n个字符串old替换为字符串new，并返回一个新的字符串
	//如果n=-1则替换所有字符串old为字符串new
	//注意：Go中string类型与Java一样是不可更改的，是一种值类型
	//因此这个函数并非字面上直接修改原字符串的方法
	fmt.Println("字符替换：",strings.Replace(str,"1234","oo",2))
	fmt.Println("原字符串：",str) //不变

	//统计字符串出现的次数
	//Count用于计算字符串str在字符串s中出现的非重叠次数
	//注意是非重叠的
	fmt.Println("统计次数：",strings.Count(str,"a"))

	//重复字符串
	//Repeat用于重复count次字符串s并返回一个新的字符串
	fmt.Println("重复字符串：",strings.Repeat(str,3))

	//修改字符串的大小写
	strings.ToLower(str)
	strings.ToUpper(str)

	//修剪字符串
	//可以使用trimSpace(s)来剔除字符串开关与结尾的空白符号
	//如果要剔除指定字符，则可以使用strings.Trim(s, "cut")来将开头与结尾的cunt去除
	//该函数的第二个参数可以包含任何字符，如果只想剔除开头或结尾的字符串
	//则可以使用TrimLeft或或者TrimRight来实现
	//对于字符串中间的也要去除，可以使用Repalce来替换成空字符

	//分割字符串
	//strings.Fields(s)将会利用一个或多个空白符号来作为动态长度的分隔符将字符串
	//分割成若干小块，并返回一个slice如果字符只包含空白符号，则返回一个长度为0的slice
	//
	//strings.Split(s, sep)用于自定义分割符号来对指定的字符串进行分割，同样返回slice
	//
	//因为这个2函数都返回slice，所以习惯使用for-range循环进行处理
	for c := range str{
		fmt.Print(c,' ')
	}
	fmt.Println()
	//注意：因为string实则是一个值类型的数组，因此使用Print里如果不用格式化则会打印出值

	//拼接slice到字符串
	//Join用于将元素类型为stirng的slice使用分割符号业拼接组成一个字符串
	var test_join []string
	test_join = []string{"aa","bb","cc"}
	fmt.Println("测试Join:",strings.Join(test_join," - "))

	//从字符串中读取内容
	//strings.NewReader(str)用于生成一个Reader并读取字符串中的内容
	//然后返回指向该Reader的指针，从其他类型读取内容的函数还有
	//Read()从[]byte中读取内容
	//ReadByte()和ReadRune()从字符串中读取下一个byte或rune
	r_str:="我的天！"
	reader_r := strings.NewReader(r_str)
	for curr_c,i:=rune(0),0;i<5;i++{
		curr_c,_,_ = reader_r.ReadRune()
		fmt.Print(curr_c," ")
	}

	//字符串与其他类型的转换
	//与字符相关的类型转换都是通过strconv包实现的
	//该包包含了一些变量用于获取程序运行操作系统下平台下int类型所占的位数
	//如strconv.IntSize

	//任何类型T转换为字符串总是成功的
	//针对从数字类型转换到字符串，Go提供了以下函数
	//strconv.Itoa(i int) string 返回数字i所表示的字符串类型的十进制数
	//strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) string 将 64 位浮点型的数字转换为字符串，其中 fmt 表示格式（其值可以是 'b'、'e'、'f' 或 'g'），prec 表示精度，bitSize 则使用 32 表示 float32，用 64 表示 float64。

	//将字符串转换为其它类型tp并不总是可能的，可能会在运行时抛出错误

	//针对从字符串类型转换为数字类型
	//strconv.Atoi(s string) (i int, err error) 将字符串转换为 int 型。
	//strconv.ParseFloat(s string, bitSize int) (f float64, err error) 将字符串转换为 float64 型。
	//
	//利用多返回值的特性，这些函数会返回2个值
	fmt.Println()
	fmt.Println("计算机的IntSize",strconv.IntSize)
}