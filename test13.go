package main

import (
	"fmt"
	//"net/http"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func testA() {
	for i := 1; i <= 3; i++ {
		fmt.Println("test() hello, world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func testB() {
	for i := 1; i <= 3; i++ {
		fmt.Println("test() hello, world" + strconv.Itoa(i) + strconv.Itoa(time.Now().Nanosecond()))
		time.Sleep(time.Second)
	}
	// 相当于 Add(-1) 这个携程执行完了
	wg.Done()
}

func testC() {
	fmt.Println("test() hello, world" + strconv.Itoa(1))
	// 相当于 Add(-1) 这个携程执行完了
	wg.Done()
}

var wg = sync.WaitGroup{}

//var locker = sync.Mutex{}
//func add() {
//	locker.Lock()
//	defer locker.Unlock()
//	num++
//}

func add(num *int, ch chan int, done chan int) {
	for v := range ch {
		*num = *num + v
		fmt.Println("v:", v)
	}
	fmt.Println("close 2")
	done <- 33
	fmt.Println("close 3 ")
}

func main() {

	//http.Handle("/", http.FileServer(http.Dir(".")))
	//http.ListenAndServe(":8080", nil)

	// select
	// 只接收IO 操作,如果select 没有default 则会一直阻塞等待io
	//timeout := make (chan bool, 1)
	//go func() {
	//	time.Sleep(1e9) // sleep one second
	//	timeout <- true
	//}()
	//ch := make (chan int)
	//select {
	//case <- ch:
	//case <- timeout:
	//	fmt.Println("timeout!")
	////default:
	////	fmt.Println("default!")
	//}

	// 管道容量 为1的时候 推1个消费1个, 如果不消费则会阻塞
	//var ch = make(chan int)
	//go func() {
	//	for i := 0; i < 5; i++ {
	//		time.Sleep(2 * time.Second)
	//		ch <- i
	//		fmt.Println(i)
	//	}
	//
	//	fmt.Println("clone 1 ")
	//}()
	////time.Sleep(5 * time.Second)
	//for i :=0;i< 20;i++ {
	//	fmt.Println("v:", <-ch)
	//}

	// 协程流程
	/**
		注: 打印顺序并不一定是程序执行顺序,打印显示可能存在延迟
		主线程 开始执行 启动 协程 1 协程 2 打印 close 4
		协程1 执行 ch 管道中推入数据 i = 1
		同时 协程2 也在执行 for range 从 ch 中 取出 i= 1
		(管道特性 长度为1 只能 推一个 取一个 , 不取就会阻塞, 新数据无法推入 ch 中)
		协程1 ch 管道中推入 i= 2 数据
		协程2 从 ch中取出 i=2 的数据
		协程1 以此类推 直到推入 i=4的数据  (协程2 跟随 协程1 从 ch中取出 i=4 的数据)
		然后 往 管道 done 推送 11
		执行打印 close 1 协程1 执行完毕
		可见 打印 <-done 11
		此时 协程2 已经ch中取出 i=4 的数据
		★此时因为使用的 for range 会继续尝试读取 ch 中的数据,因为 协程1 已经不在写入数据
		★导致 协程2 阻塞形成死锁
		这时需要 close 关闭 ch  for range 不在从 ch 中读取数据,
		协程2 继续执行
		打印 close 5  close 2
		协程2 往 <-done 推入 33
		主线程 打印 <-done 33
		输出 num 计算结果
	*/
	var (
		num  = 0
		ch   = make(chan int)
		done = make(chan int)
	)
	// 协程1
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println( i )
			ch <- i
		}
		done <- 11
		fmt.Println( "close 1 ")
	}()
	// 协程2
	go add(&num, ch, done)
	fmt.Println( "close 4 ")
	fmt.Println("<-done", <-done)
	time.Sleep(2 * time.Second)
	close(ch)
	fmt.Println("close 5 ")
	fmt.Println("<-done", <-done)
	fmt.Printf("num = %d\n", num)

	// 管道
	//var ch = make(chan int)
	//// 数据写入管道
	//go func() {
	//	ch <- 1
	//	ch <- 2
	//	ch <- 3
	//	close(ch)
	//}()

	//从管道中获取数据
	//for i := 0; i < 1; i++ {
	//	fmt.Println(<- ch)
	//}
	//fmt.Println(<- ch)

	// 判断管道是否关闭
	//v, ok := <-ch
	//if ok {
	//	fmt.Println(v)
	//	fmt.Println("open")
	//} else {
	//	fmt.Println("close")
	//}

	//for v := range ch {
	//	fmt.Println(v)
	//}

	//n := 0
	//wg.Add(4)
	//for i := 0; i < 4; i++ {
	//	go func(v int) {
	//		v++
	//		fmt.Println(v)
	//		wg.Done()
	//	}(n)
	//}
	//wg.Wait()
	//fmt.Println("xxx",n)

	//for i := 0; i <= 100 ; i++{
	//	go fmt.Println(i)
	//}
	//time.Sleep(time.Second)

	//wg := sync.WaitGroup{}
	//wg.Add(100)
	//for i := 0; i < 100; i++ {
	//	go func(i int) {
	//		fmt.Println(i)
	//		wg.Done()
	//	}(i)
	//}
	//wg.Wait()

	//c := make(chan bool, 100)
	//for i := 0; i < 100; i++ {
	//	go func(i int) {
	//		fmt.Println(i)
	//		c <- true
	//	}(i)
	//}
	//for i := 0; i < 100; i++ {
	//	<-c
	//}

	// 最简单的 携程,主进程sleep 等待携程执行完成
	// go testA() 就启动了携程
	//go testA()
	//for i := 1; i <= 3; i++ {
	//	fmt.Println("main() hello,Golang" + strconv.Itoa(i))
	//	time.Sleep(time.Second)
	//}

	// 单携程运行
	// Add(1) 说明我有一个携程要执行
	//wg.Add(1)
	//go testB()
	// 告诉主线程,等携程执行完成再退出
	//wg.Wait()
	//fmt.Println("main() hello,Golang" + strconv.Itoa(1))

	// 并发并行 协程
	//wg.Add(3)
	//for i := 0; i < 3; i++ {
	//	go testC()
	//}
	//wg.Wait()

	// 容易发错的地方 指针 给元素赋值的时候 是拷贝还是赋值
	//pase_student()

	//testD()
}

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]student)
	//m := make(map[string]*student)
	stus := []student{
		{
			Name: "zhou",
			Age:  24,
		}, {
			Name: "li",
			Age:  23,
		}, {
			Name: "wang",
			Age:  22,
		},
	}
	// 错误写法
	// 如果 stu为指针,那实际上一致指向同一个指针,最终是一个 struct 的值拷贝
	for _, stu := range stus {
		m[stu.Name] = stu
		//m[stu.Name] = &stu
	}
	fmt.Println(m)
	for k, v := range m {
		println(k, "=>", v.Name)
	}
	// 正确
	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = stus[i]
	}
	fmt.Println(m)
	for k, v := range m {
		println(k, "=>", v.Name)
	}
}

// 协程
func testD() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
