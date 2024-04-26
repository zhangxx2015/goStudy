// 接收者堵塞
package main

import "fmt"

func main() {
	var result int       // 声明一个整型变量result
	ch := make(chan int) // 创建一个整型通道ch

	result = <-ch // 从通道ch接收数据，但此时通道ch中没有数据可接收，因此接收操作会被阻塞

	go func() { // 在一个新的Go协程中执行匿名函数
		ch <- 1 // 向通道ch发送数据1
	}()

	fmt.Println(result) // 由于接收操作被阻塞，此处不会执行，程序会发生死锁错误
}

/* error:
fatal error: all goroutines are asleep - deadlock!

接收操作是阻塞的（协程或函数中的），直到发送者可用：如果通道中没有数据，接收者就阻塞了。
*/
