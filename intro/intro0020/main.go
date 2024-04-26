// 超时机制
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)    // 创建一个整型通道ch
	quit := make(chan bool) // 创建一个布尔型通道quit

	// 在一个新的Go协程中执行匿名函数
	go func() {
		for {
			select {
			case num := <-ch: // 从通道ch中接收数据，并将其赋值给变量num
				fmt.Println("num = ", num) // 打印变量num的值
			case <-time.After(3 * time.Second): // 经过3秒后，从time.After通道接收到数据
				fmt.Println("超时") // 打印超时信息
				quit <- true      // 向通道quit发送true值
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i                 // 向通道ch发送数据i
		time.Sleep(time.Second) // 睡眠1秒钟
	}

	<-quit              // 从通道quit接收数据，阻塞直到接收到数据
	fmt.Println("程序结束") // 打印程序结束信息
}

/*
num =  0
num =  1
num =  2
num =  3
num =  4
超时
程序结束
*/
