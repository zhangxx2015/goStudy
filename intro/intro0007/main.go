package main

import (
	"fmt"
)

func main() {
	defer func() {
		// recover() 可以将捕获到的panic信息打印
		if err := recover(); err != nil {
			// 捕获异常，恢复程序或做收尾工作
			// revocer 调用后，抛出的 panic 将会在此处终结，不会再外抛，
			// 但是 recover，并不能任意使用，它有强制要求，必须得在 defer 下才能发挥用途。
			fmt.Println("捕获到一个异常, 错误为:", err)
		}
	}()
	panic("crash") // 主动触发一个异常
}
