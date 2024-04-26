package main

import (
	"fmt"
)

func main() {
	loops := 999999         // 循环次数
	waitc := make(chan int) // 创建一个整型通道waitc

	for i := 0; i < loops; i++ { // 循环创建goroutine
		go func(i int) { // 在一个新的Go协程中执行匿名函数
			fmt.Printf("This is %d\n", i) // 打印当前循环的索引i
			if i == loops-1 {             // 如果当前循环的索引i等于循环次数减1
				close(waitc) // 关闭通道waitc
			}
		}(i) // 传入当前循环的索引i
	}

	<-waitc // 接收通道waitc的数据，此处会阻塞直到通道被关闭
}

/*
结果太长了,不列举了.
*/
/*
由于waitc没有被传入数据,因此在执行到<-waitc时进入接收者堵塞状态.
这时候将会有足够的时长保证所有goroutine被执行,直到最后一个goroutine宣布关闭通道waitc为止.
*/
