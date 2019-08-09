package main

//注意同一个目录下不能有不同名的包
import (
	"fmt"
	ch4 "./Chapter4"
	helloworld	"./HelloWorld"
)

func main() {
	//C4
	ch4.TestBase()

	//
	helloworld.Hello()

	fmt.Println("xxx")
}
