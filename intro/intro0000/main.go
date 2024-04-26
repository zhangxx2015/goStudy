package main // 包名称 主程序必须为 main

import ( // 导入包
	"fmt"
)

var a string = "a" // 标准写法
var (              // 批量声明
	b = "" // 类型推导
	c int
	d bool    // 默认为 false
	e float32 // 默认为 0
)

const PI = 3.1415 // 常量
const (           // 批量声明常量
	_ = iota // iota是go语言的常量计数器，只能在常量的表达式中使用。 iota在const关键字出现时将被重置为0。
	// const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func init() { // 初始化过程, 在调用主函数前由系统调用, 类似于 php 的魔法函数
	a = "第一个变量"
}

func main() { // 主函数
	var b string = "b"  // 覆盖全局b
	c := "c"            // var 简化写法 # 只能用在函数内部
	f, g := "你好", "世界"  // 使用类型推导初始化多个变量, 类似 python 的元组
	h := "hello\nworld" // 单行字符串, 支持转义字符
	i := `hello
world\n` // 多行字符字符串, 不支持转义符
	fmt.Println(PI, KB, MB, GB, TB, PB)
	fmt.Println(a, b, c, d, e, f, g)
	fmt.Println(h)
	fmt.Println(i)

}
