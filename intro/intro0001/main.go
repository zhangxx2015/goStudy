package main

import (
	"fmt"
)

const ( // 批量声明常量
	_ = iota // iota是go语言的常量计数器，只能在常量的表达式中使用。 iota在const关键字出现时将被重置为0。
	// const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。
	A = iota // 1
	B = iota // 2
	_ = iota // 3 golang 中 _ 表示忽略一个已经存在的变量, 这里可以用于跳过一个数
	D = iota // 4
)

func main() {
	fmt.Println(true)
	fmt.Println(false)
	fmt.Println(nil)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(D)
}
