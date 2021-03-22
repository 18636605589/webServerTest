package main

import (
	"fmt"
)

var age = test()

var Age int = 20
//下面的语句等价与赋值语句 赋值语句不能在函数体外部,无法通过编译
//Name := "Tom"

//init函数
//每一个源文件都可以包含一个 init 函数，该函数会在 main 函数执行前，
//被 Go 运行框架调用，也 就是说 init 会在 main 函数前被调用。
func init() {
	fmt.Println("init call")
}

func main() {
	fmt.Println("so young", age)
	fmt.Println("so young", Age)
	//fmt.Println("so young", Name)
}

func test() int {
	fmt.Println("test call" +
		"---")
	return 90
}


//细节说明
//1.如果一个文件同时包含全局变量定义，init 函数和 main 函数，则执行的流程全局变量定义->init函数->main 函数
//2. init函数最主要的作用，就是完成一些初始化的工作
