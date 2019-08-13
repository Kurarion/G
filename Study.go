package main

//注意同一个目录下不能有不同名的包
import (
	ch4 "./Chapter4"
	"fmt"
)

func main() {
	//C4
	//ch4.TestBase()

	//
	//helloworld.Hello()

	fmt.Println("xxx")

	//TestConst
	ch4.TestConst()

	//TestVar
	//ch4.TestGoos()

	//TestPrint
	ch4.TestPint()

	//TestSameName
	ch4.TestSameName()
}
