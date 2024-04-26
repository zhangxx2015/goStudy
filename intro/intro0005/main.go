package main // 包名称 主程序必须为 main

import ( // 导入包
	"fmt"
)

func init() { // 初始化过程, 在调用主函数前由系统调用, 类似于 php 的魔法函数
	fmt.Println("hello from init")
}

func main() { // 主函数
	fmt.Println("hello from main")

}
