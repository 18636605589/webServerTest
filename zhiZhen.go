package main

import "fmt"

func main() {
	//var i int = 10
	//fmt.Println("i的地址=", &i)
	//
	////ptr是一个指针变量
	////ptr的类型是*int
	////ptr本身的值是&i
	////var ptr *int = &i
	//ptr := &i
	//fmt.Printf("ptr=%v ptr type is %T\n ", ptr, ptr)
	//fmt.Printf("ptr的地址是%v\n", &ptr)
	//fmt.Printf("ptr指向的值是%v\n", *ptr)

	//example()

	example2()
}


//指针使用细节
//1.值类型都有对应的指针类型，形式为 *数据类型，比如 int 的对应的指针就是 *int,
//float32 对应的指针类型就是 *float32, 依次类推。
//2.值类型包括:基本数据类型 int 系列, float 系列, bool, string 、数组和结构体 struct
func example() {
	var num int = 999
	fmt.Printf("num的地址=%v\n", &num)
	var ptr1 int = num
	fmt.Printf("ptr1=%v\n", ptr1)
	fmt.Printf("ptr1的地址=%v\n", &ptr1)
	var ptr *int = &num
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr的地址=%v\n", &ptr)
	*ptr = 1000
	fmt.Printf("num的值=%v\n", num)
}

func example2() {
	var a int = 300
	var b int = 400
	var ptr *int = &a
	*ptr = 100
	ptr = &b
	*ptr = 200
	fmt.Printf("a=%d,b=%d,*ptr=%d\n", a, b, *ptr)
}

