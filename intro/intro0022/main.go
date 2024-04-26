// 通过range进行消费
package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		time.Sleep(1 * time.Hour) // 在一个新的Go协程中睡眠1小时，模拟长时间操作
	}()

	c := make(chan int) // 创建一个整型通道c

	go func() {
		for i := 0; i < 10; i = i + 1 { // 循环10次
			c <- i // 向通道c发送数据i
		}
		close(c) // 关闭通道c
	}()

	for i := range c { // 从通道c中循环接收数据，直到通道被关闭
		fmt.Println(i) // 打印接收到的数据
	}

	fmt.Println("Finished") // 打印"Finished"表示程序执行完毕
}

/*
0
1
2
3
4
5
6
7
8
9
Finished
*/
