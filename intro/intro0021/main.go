// 同步操作
package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	time.Sleep(3 * time.Second) // 阻塞 3 秒
	// 通知任务已完成
	done <- true
}

func main() {
	fmt.Println("程序启动")
	done := make(chan bool, 1) // 创建一个容量为1的布尔型通道done
	go worker(done)            // 在一个新的Go协程中执行worker函数，传入通道done
	fmt.Println("主程序完成")
	// 等待任务完成
	<-done // 从通道done接收数据，阻塞直到接收到数据
	fmt.Println("程序退出")
}
