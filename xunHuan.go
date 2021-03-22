package main

import (
	"fmt"
)

type abc struct {
	A int
	B int
}

func solution(day int) int {
	if day == 10 {
		return 1
	}

	return (solution(day + 1) + 1) * 2
}

func solution2(n int) int {
	tmp := 1
	for i:=1; i < n; i++ {
		tmp = (tmp+1) * 2
	}
	return tmp
}

func main() {

	//res := solution2(3)
	//fmt.Println("res=", res)

	//fmt.Println(solution(8))

	//var count int = 0
	//for {
	//	rand.Seed(time.Now().UnixNano())
	//	n := rand.Intn(100) + 1 //生成[0, 100)
	//	fmt.Println("n=", n)
	//	count++
	//	if n == 99 {
	//		break
	//	}
	//}
	//fmt.Printf("生成99用了%d次\n", count)

	//break细节
	//break语句出现在多层嵌套的语句块中时，可以通过标签指明要终止的是哪一层语句块
	//break默认会跳出最近的for循环
	var a = 0
	var n int = 30
	fmt.Println("ok1")
	if n > 20 {
		goto for1
	}
for1:
	for {
		a++
		fmt.Println("a=", a)
		for {
			break for1
		}
	}

}
