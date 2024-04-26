// 通道与通道的值传递
package main

import "fmt"

func main() {
	var result int      // 声明一个整型变量result
	a := make(chan int) // 创建一个整型通道a
	var b chan int      // 声明一个整型通道b
	b = a               // 将通道a赋值给通道b

	go func() { // 在一个新的Go协程中执行匿名函数
		result = <-b // 从通道b接收数据，并将其赋值给result变量
	}()

	a <- 123                // 向通道a发送数据123
	fmt.Println(result)     // 打印result的值，此时result的值为123
	fmt.Println("&a: ", &a) // 打印通道a的内存地址
	fmt.Println("&b: ", &b) // 打印通道b的内存地址
}

/* output
123
&a:  0xc000096018
&b:  0xc000096020
*/
/*
可以看出通道与通道间的传递是透过值传递。
但是通道本身也是一个‘指针’，指向一个存有通道内部值的数据结构（这里是123）。
这个数据结构并没有被值传递。
*/
