// 发送者堵塞
package main

import "fmt"

func main() {
	var result int       // 声明一个整型变量result
	ch := make(chan int) // 创建一个整型通道ch

	ch <- 1 // 向通道ch发送数据1，发送者会被阻塞，因为通道没有接收者

	go func() { // 在一个新的Go协程中执行匿名函数
		result = <-ch // 从通道ch中接收数据，并将其赋值给result变量
	}()

	fmt.Println(result) // 打印result的值，由于通道接收操作在新的协程中执行，此时result还未被赋值，因此输出结果为0
}
