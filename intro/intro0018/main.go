package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()         // 记录当前时间
	c := make(chan interface{}) // 创建一个通道c

	go func() { // 在一个新的Go协程中执行匿名函数
		time.Sleep(2 * time.Second) // 睡眠2秒钟
		close(c)                    // 关闭通道c
	}()

	fmt.Println("Blocking on read...") // 打印信息，表示正在阻塞等待读取通道

	select {
	case <-c: // 从通道c中接收数据
		fmt.Printf("Unblocked %v later.\n", time.Since(start)) // 打印从开始阻塞到解除阻塞所经过的时间
	}
}

/*
Blocking on read...
Unblocked 2.013281s later.
*/
