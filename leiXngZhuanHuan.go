package main

import (
	"fmt"
	"strconv"
)

func main() {

	//var i int32 = 1000
	//var n1 float32 = float32(i)
	//var n2 int8 = int8(i)
	//var n3 int64 = int64(i)
	//// i=1000 n1=1000 n2=-24 n3=1000
	//fmt.Printf("i=%v n1=%v n2=%+v n3=%+v\n", i, n1, n2, n3)
	//// i type is int32
	//fmt.Printf("i type is %T\n", i)

	//基本数据类型和 string 的转换
	//1. 基本类型转 string 类型 fmt.Sprintf("%参数", 表达式)
	//var i int32 = 1000
	//var n1 float32 = float32(i)
	//str := fmt.Sprintf("%d", i)
	//// str type string str "1000"
	//fmt.Printf("str type %T str %q\n", str, str)
	//str2 := fmt.Sprintf("%f", n1)
	//// str type string str "1000.000000"
	//fmt.Printf("str type %T str %q\n", str2, str2)
	//isTrue := false
	//str3 := fmt.Sprintf("%t", isTrue)
	//// str type string str "false"
	//fmt.Printf("str type %T str %q\n", str3, str3)
	//myChar := 'b'
	//str4 := fmt.Sprintf("%c", myChar)
	//// str type string str "b"
	//fmt.Printf("str type %T str %q\n", str4, str4)

	//使用strconv包的函数
	//num3 := 99
	//num4 := 23.456
	//b2 := true
	//str5 := strconv.FormatInt(int64(num3), 10)
	//fmt.Printf("str type %T str %q\n", str5, str5)
	////f格式 10 表示小数保留10位，64表示这个小数是float64
	//str6 := strconv.FormatFloat(num4, 'f', 10, 64)
	//fmt.Printf("str type %T str %q\n", str6, str6)
	//str7 := strconv.FormatBool(b2)
	//fmt.Printf("str type %T str %q\n", str7, str7)
	//var num5 int64 = 4567
	//str8 := strconv.Itoa(int(num5))
	//fmt.Printf("str type %T str %q\n", str8, str8)

	defaultVal()

	//stringToOther()

}

//基本数据类型的默认值
// 整型 0
// 浮点型 0
// 字符型 “”
// 布尔类型 false

func defaultVal() {
	var a int
	var b float32
	var c string
	var isTrue bool
	fmt.Printf("a=%d, b=%v, c=%v, isTrue=%v", a, b, c, isTrue)
}

func stringToOther() {
	//2. string 类型转基本类型
	var str string = "true"
	var b bool

	//ParseBool会返回两个值(value bool, err error)
	//我们只想获取value 不想获取err所以用_忽略
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v\n", b, b)

	var str2 string = "123456789"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2, 10, 64)
	n2 = int(n1)
	fmt.Printf("n1 type %T n1=%v\n", n1, n1)
	fmt.Printf("n2 type %T n2=%v\n", n2, n2)

	var str3 string = "123456.2345"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type %T f1=%v\n", f1, f1)

	//string 转基本数据类型的注意事项
	//在将 String 类型转成 基本数据类型时，要确保 String 类型能够转成有效的数据，
	//比如 我们可以 把 "123" , 转成一个整数，但是不能把 "hello" 转成一个整数，如果这样做，Golang 直接将其转成 0 ，
	//其它类型也是一样的道理. float => 0 bool => false
}
