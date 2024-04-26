package main

import "fmt"

func main() {
	var result int       // 声明一个整型变量result
	ch := make(chan int) // 创建一个整型通道ch

	go func() { // 在一个新的Go协程中执行匿名函数
		result = <-ch // 从通道ch中接收数据，并将其赋值给result变量
	}()

	ch <- 1             // 向通道ch发送数据1
	fmt.Println(result) // 打印result的值
}
