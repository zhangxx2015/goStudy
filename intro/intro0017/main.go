package main

import (
	"fmt"

	"github.com/kirinlabs/HttpRequest"
)

// 同时监控不同的channel，配上default，select也不会阻塞。
func main() {
	ch1 := make(chan int) // 创建一个整型通道ch1
	ch2 := make(chan int) // 创建一个整型通道ch2
	ch3 := make(chan int) // 创建一个整型通道ch3

	go func() {
		for {
			Getdata("https://www.baidu.com", ch1) // 在一个新的Go协程中执行Getdata函数，将结果发送到通道ch1
			Getdata("https://cn.bing.com", ch2)   // 在一个新的Go协程中执行Getdata函数，将结果发送到通道ch2
			Getdata("https://cn.bing.com", ch3)   // 在一个新的Go协程中执行Getdata函数，将结果发送到通道ch3
		}
	}()

	go func() {
		for {
			select {
			case v := <-ch1: // 从通道ch1接收数据，并将其赋值给变量v
				fmt.Println("信道1的结果：", v) // 打印变量v的值
			case v := <-ch2: // 从通道ch2接收数据，并将其赋值给变量v
				fmt.Println("信道2的结果：", v) // 打印变量v的值
			case v := <-ch3: // 从通道ch3接收数据，并将其赋值给变量v
				fmt.Println("信道3的结果：", v) // 打印变量v的值
			default:
				continue
			}
		}
	}()

	select {}
}

func Getdata(url string, ch chan int) {
	req, err := HttpRequest.Get(url) // 发送HTTP GET请求
	if err != nil {
		// 处理错误
	} else {
		ch <- req.StatusCode() // 将HTTP响应状态码发送到通道ch
	}
}
