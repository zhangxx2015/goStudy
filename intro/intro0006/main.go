package main // 包名称, 入口函数包名必须为 main

import "fmt" // 导入库

// 定义一个结构体类型
type Person struct {
	name string // 名称, 类型
	age  int
}

func main() {
	defer func() { // 延迟调用, 将在函数执行完成后执行, 一般用于关闭打开的句柄
		fmt.Println("all done")
	}()
	// 变量声明和赋值
	var num int = 10       // 声明变量写法1
	str := "Hello, World!" // 声明变量写法2 ( 类型由初始化值自动推导 )

	// 条件语句
	if num > 0 {
		fmt.Println("num is positive")
	} else if num < 0 {
		fmt.Println("num is negative")
	} else {
		fmt.Println("num is zero")
	}

	// 循环语句
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 数组和切片
	nums := [3]int{1, 2, 3}
	slice := []int{4, 5, 6}

	// 结构体
	person := Person{name: "Alice", age: 25}

	// 函数
	result := add(3, 4)

	// 匿名函数, 直接调用
	func(msg string) {
		fmt.Println("msg is ", msg)
	}("HAHA")

	// 打印输出
	fmt.Println(num)
	fmt.Println(str)
	fmt.Println(nums)
	fmt.Println(slice)
	fmt.Println(person)
	fmt.Println(result)
}

// 函数定义
func add(a, b int) int { // func 关键字, 函数名称 ( 函数入参 类型, ... ) 返回值 ...
	return a + b // 函数体
}
