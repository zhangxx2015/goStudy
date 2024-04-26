// 通过channel通知，从而退出死循环
package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan interface{}) // 创建一个通道done

	go func() { // 在一个新的Go协程中执行匿名函数
		time.Sleep(2 * time.Second) // 睡眠2秒钟
		close(done)                 // 关闭通道done
	}()

	workCounter := 0 // 工作计数器

loop:
	for {
		select {
		case <-done: // 从通道done中接收数据
			break loop // 跳出循环
		default:
		}

		// 模拟工作
		workCounter++
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("在通知退出循环时，执行了%d次.\n", workCounter) // 打印工作计数器的值
}

/*
在通知退出循环时，执行了2次.
*/
