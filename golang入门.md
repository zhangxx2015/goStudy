# golang学习

# 历史背景
> Go 语言主要是为了解决 Google 内部在自己开发过程中面临的一些挑战所诞生的：
1. 多核硬件架构
2. 超大规模分布式计算集群
3. Web 模式导致的前所未有的开发规模和更新速度

# 特性

## 1. 开源 (google)

## 2. google背书

## 3. 核心开发团队强大

### 三位创始人
#### Rob Pike
>     Unix 的早期开发者
>     UTF - 8创始人

#### Ken Thompson
>     Unix 的创始人
>     C 语言创始人
>     1983 年获图灵奖

#### Robert Griesemer
>     Google V8 JS Engine
>     Hot Spot 开发

## 2.编译型语言
1. 性能高
2. 易于部署

## 3. 21世纪的C语言
1. 简单易学习
2. 原生支持并发
3. 开发效率搞
4. 执行性能好
5. 云原生支持好

## 4. golang语言特性
1. 自动立即回收。
2. 更丰富的内置类型。
3. 函数多返回值。
4. 错误处理。
5. 匿名函数和闭包。
6. 类型和接口。
7. 并发编程。
8. 反射。
9. 语言交互性。

## 5. 语言优势
1. 高效
* Go 语言是一种编译的强类型语言。
* Go 在支持了垃圾回收的同时，为了提供更高效的内存访问，Go也提供了通过指针可以直接进行内存访问。
2. 生产力
* Go 不仅语法简洁，还有特别的接口类型
* Go 还有些编程约束，直接就为开发者做出了更好的选择，譬如在程序的扩展上：一般语言都会支持 复合 和 继承，很多面向对象编程的书籍中都会谈到 复合大于继承，及相关原理，那么在 Go 语言中，只支持 复合。
3. 云计算语言
* 越来越多的应用都采用了 Go 语言进行开发，例如 docker 、kubernetes 等。由于云端大量使用了 kubernetes 和 doker ，所以 Go 语言也被称为 云计算语言
4. 区块链语言
* 区块链是最近继 AI 以来最热门的话题了，非常热门的 以太网 和 HYPERLEDGER 等都是可以用 Go 语言来进行开发的，所以 Go 语言也被称为 区块链开发语言 。


# 1. go的安装

### 1.1 下载地址


> golang官方网站[需要翻墙]
> https://go.dev/
>
> 下载
>
> https://go.dev/dl/

> 国内镜像下载
>
> 七牛云 - Goproxy.cn
> https://goproxy.cn/

### 1.2 根据不同的构架和平台下载相应的安装包

![下载包](.\goStudy\下载包.png)

### 1.3 安装

#### 1.3.1 windows平台安装

##### 1.3.1.1 双击安装包

![win_install](.\goStudy\win\install.png)

##### 1.3.1.2 点击 Next

![win_setup1](.\goStudy\win\setup1.png)

##### 1.3.1.3 点击 Next

![win_setup2](.\goStudy\win\setup2.png)

##### 1.3.1.4 设置安装路径, 建议使用 c:\go

![win_setup3](.\goStudy\win\setup3.png)

##### 1.3.1.5 开始进行安装, 点击 Install

![win_setup4](.\goStudy\win\setup4.png)

##### 1.3.1.6 等待安装完成

![win_setup5](.\goStudy\win\setup5.png)

##### 1.3.1.7 点击 是

![win_setup6](.\goStudy\win\setup6.png)

##### 1.3.1.8 等待安装完成

![win_setup7](.\goStudy\win\setup7.png)

##### 1.3.1.9 完成安装, 点击 Finish

![win_setup8](.\goStudy\win\setup8.png)

#### 1.3.2 linux平台安装

##### 1.3.2.1 SSH登录linux主机

![登录linux主机](.\goStudy\linux\0001_ssh.png)

##### 1.3.2.2 执行安装命令

```bash
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.1.linux-amd64.tar.gz
```

![登录linux主机](.\goStudy\linux\0002_setup.png)

##### 1.3.2.3 设置环境变量

```bash
export PATH=$PATH:/usr/local/go/bin
```

![设置环境变量](.\goStudy\linux\0003_env.png)

##### 1.3.2.4 验证安装版本

```bash
go version
```

![设置环境变量](.\goStudy\linux\0004_ver.png)




### 1.4 配置GOPATH

> `GOPATH`是一个环境变量, 用来表明你写的`go`项目的存放路径
>
> `GOPATH`路径最好只设置一个, 所有的项目代码都放到`GOPATH`的`src`目录下。

#### 1.4.1 windows平台配置

##### 1.4.1.1 查看是否配置 GOPATH

```powershell
echo %GOPATH%
```

![win_gopath1](.\goStudy\win\gopath1.png)

##### 1.4.1.2 添加 GOPATH

##### 1.4.1.2.1 我的电脑->属性

![win_gopath2](.\goStudy\win\gopath2.png)

##### 1.4.1.2.2 高级系统设置

![win_gopath3](.\goStudy\win\gopath3.png)

##### 1.4.1.2.3 环境变量

![win_gopath4](.\goStudy\win\gopath4.png)

##### 1.4.1.2.4 新建

![win_gopath5](.\goStudy\win\gopath5.png)

##### 1.4.1.2.5 新建系统变量

> 变量名 GOPATH
> 变量值 d:\ws\ ( 设置为go项目地址 )

![新建系统变量](.\goStudy\win\gopath6.png)

##### 1.4.1.2.6 选择 path 变量, 编辑

![选择 path 变量, 编辑](.\goStudy\win\gopath7.png)

##### 1.4.1.2.7 添加go的安装目录和GOPATH目录

![添加go的安装目录和GOPATH目录](.\goStudy\win\gopath8.png)

##### 1.4.1.2.8 在GOPATH目录下新建三个文件夹
> bin 存放编译后生成的可执行文件
> pkg 存放编译后生成的归档文件
> src 存放源代码

![在GOPATH目录下新建三个文件夹](.\goStudy\win\gopath9.png)



#### 1.4.2 linux平台配置

##### 1.4.2.1 查看是否配置 GOPATH

```shell
echo $GOPATH
```

![查看是否配置 GOPATH](.\goStudy\linux\0005_show_env.png)


##### 1.4.2.2 创建项目目录

> /home目录下, 建立一个名为gopath的目录, 作为存放项目的目录
> gopath下再建立3个目录pkg、bin、src

```shell
mkdir -p /home/gopath /home/gopath/src /home/gopath/pkg /home/gopath/bin
```

![创建项目目录](.\goStudy\linux\0006_mkdir.png)

##### 1.4.2.3 添加环境变量

```shell
echo "export GOROOT=/usr/local/go">>/etc/profile && echo "export GOPATH=/home/gopath">>/etc/profile && echo "export PATH=$PATH:$GOROOT/bin:$GOPATH/bin">>/etc/profile
```

![添加环境变量](.\goStudy\linux\0007_set_env.png)

##### 1.4.2.4 使环境变量立即生效

```shell
source /etc/profile
```

![使环境变量立即生效](.\goStudy\linux\0008_apply.png)

##### 1.4.2.5 验证环境变量

```shell
go env|grep GOROOT && go env|grep GOPATH
```

![验证环境变量](.\goStudy\linux\0009_show_env.png)

## 2. 第一个golang程序
### 2.1 进入 GOPATH 的 src 目录

![](.\goStudy\hello\0001.png)

### 2.2 创建项目文件夹 hello
![](.\goStudy\hello\0002.png)

### 2.3 进入项目文件夹
![](.\goStudy\hello\0003.png)

### 2.4 创建 main.go 文件

![](.\goStudy\hello\0004.png)

### 2.5 添加 代码

```go
package main

import (
    "fmt"
)

func main(){
    fmt.Println("hello from golang")
}
```

![](.\goStudy\hello\0005.png)


### 2.6 初始化项目

```shell
go mod init
```

![](.\goStudy\hello\0006.png)


### 2.7 编译代码

> 在 hello 目录下执行
> go编译器会去 GOPATH的src目录下查找你要编译的hello项目
> 编译得到的可执行文件会保存在执行编译命令的当前目录下, 如果是windows平台会在当前目录下找到hello.exe可执行文件。
```shell
go build
```

![](.\goStudy\hello\0007.png)


### 2.8 执行编译后的二进制文件

```shell
ls
./hello
```

![](.\goStudy\hello\0008.png)

### 2.9 不生成二进制文件, 直接运行

```shell
go run main.go
```

![](.\goStudy\hello\0009.png)


### 2.10 代码解析

```go
package main // 包名称 主程序必须为 main

import ( // 导入包
    "fmt"
)

func main(){// 主程序
    fmt.Println("hello from golang")
}
```

## 3. 开发工具
### VSCode
> VScode 安装教程参见：https://www.runoob.com/w3cnote/vscode-tutorial.html

> 然后我们打开 VSCode 的扩展（Ctrl+Shift+P）：
![](.\goStudy\vscode\0001.jpg)

> 搜索 go:
![](.\goStudy\vscode\0002.jpg)

> 点击安装，安装完成后我们就可以使用代码提示、测试、调试等功能了。
![](.\goStudy\vscode\0003.gif)

### GoLand
> GoLand 是 Jetbrains 家族的 Go 语言 IDE，有 30 天的免费试用期。
> 安装也很简单访问 Goland 的下载页面，根据你当期的系统环境三大平台（Mac、Linux、Windows）下载对应的软件。

![](.\goStudy\goland\0001.jpg)

### LiteIDE
> LiteIDE 是一款开源、跨平台的轻量级 Go 语言集成开发环境（IDE）。
> 支持的 操作系统 
> Windows x86 (32-bit or 64-bit)
> Linux x86 (32-bit or 64-bit)
>
> 下载地址 ：http://sourceforge.net/projects/liteide/files/
> 源码地址 ：https://github.com/visualfc/liteide

![](.\goStudy\liteide\0001.png)




## 4. 极速语法入门

### 基础语法
```go
package main // 包名称, 入口函数包名必须为 main

import "fmt" // 导入库

// 定义一个结构体类型
type Person struct {
	name string // 名称, 类型
	age  int
}

func main() {
	defer func(){// 延迟调用, 将在函数执行完成后执行, 一般用于关闭打开的句柄
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
```

### 进阶语法
> 内容覆盖：变量、判断、循环、函数、数组、指针、结构体、类型转换、接口类、并行

```go
package main // 包名称 主程序必须为 main

import ( // 导入包
	"fmt"
	"time"
)

var ( //第四种, 这种因式分解关键字的写法一般用于声明全局变量
	k int
	l bool
)

// 第五种,常量
const q1, q2, q3 = 1, true, "str"

func main() {
	fmt.Println("Hello")         // 使用fmt包的Println方法,输出内容
	var str string = "xx" + "yy" // 定义变量
	fmt.Println(str)

	// 一, 变量
	// 第一种，指定变量类型，如果没有初始化，则变量默认为零值。
	var a, b int = 1, 2
	var c int
	fmt.Println(a+b, c)

	// 第二种，根据值自行判定变量类型。(值推导)
	var d = true
	fmt.Println(d)

	// 第三种, 直接声明+赋值
	f := "test"
	f2 := 23.456
	fmt.Println(f, f2)

	// 第四种 全局变量
	fmt.Println(k, l)

	// 第五种 常量
	fmt.Println(q1, q2, q3)
	fmt.Println(len(q3)) // 求字符串长度
	fmt.Println("--------------------------------")
	fmt.Println()
	fmt.Println()

	// 二, 判断
	if a < 20 {
		fmt.Println("a<20 if ")
	} else {
		fmt.Println("a<20 else")
	}

	var grade = ""
	score := 90
	switch score { // 条件分支 int类型
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	fmt.Println("分数:", score, "级别:", grade)

	switch { // 条件分支 判断
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Println("--------------------------------")
	fmt.Println()
	fmt.Println()

	// 三, 循环
	for i := 0; i < 5; i++ {
		fmt.Println("循环输出:", i)
	}

	sum := 7
	for sum < 10 {
		fmt.Println("循环输出:sum=", sum)
		sum += 2
	}
	for true { // 死循环, 也可以写为 for {}
		sum += 2
		if sum > 20 {
			fmt.Println("循环输出:sum=", sum)
			break
		}
	}

	// 数组
	stirngList := []string{"Tom", "Lily", "Wang", "Kim"}
	for idx, item := range stirngList { // 遍历
		fmt.Println("索引:", idx, "值:", item)
	}

	// 关键字 make 用来为 slice，map 或 chan 类型（注意：也只能用在这三种类型上）分配内存和初始化一个对象
	// 注意: make 返回 类型的本身 而不是 指针，而返回值也依赖于具体传入的类型，
	// 因为这三种类型（slice，map 和 chan）本身就是引用类型，所以就没有必要返回他们的指针了
	map1 := make(map[int]float32) // 字典
	map1[1] = 1.2
	map1[2] = 1.5
	for key, value := range map1 {
		fmt.Println("key:", key, "val:", value)
	}
	fmt.Println("--------------------------------")
	fmt.Println()
	fmt.Println()

	// 四, 函数
	x1 := "ss"
	y1 := "xx"
	fmt.Println("交换前:", "x1=", x1, "y1=", y1)
	x1, y1 = swap(x1, y1) // 调用函数, 返回多个结果
	fmt.Println("交换后:", "x1=", x1, "y1=", y1)
	fmt.Println("--------------------------------")
	fmt.Println()
	fmt.Println()

	// 五, 数组
	arrayTest := []int{1, 2}                                       // 定义数组,并初始化
	fmt.Println("数组长度:", len(arrayTest))                           // 计算数组长度
	arrayTest = append(arrayTest, 3)                               // 追加
	fmt.Println("数组长度:", len(arrayTest))                           // 计算数组长度
	fmt.Println("数组内容:", arrayTest[0], arrayTest[1], arrayTest[2]) // 输出数组的每个元素的值

	matrix := [][]int{}       // 二维数组
	row1 := []int{1, 2, 3}    // 定义数组row1并初始化
	row2 := []int{4, 5, 6, 7} // 定义数组row2并初始化
	matrix = append(matrix, row1)
	matrix = append(matrix, row2)
	fmt.Println("matrix数组的第0个元素:", matrix[0])
	fmt.Println("matrix数组的第1个元素:", matrix[1])
	fmt.Println("matrix数组的内容:", matrix)
	fmt.Println("--------------------------------")
	fmt.Println()
	fmt.Println()

	// 六, 指针
	pp := 10
	fmt.Println("变量pp的值是:", pp)
	fmt.Println("变量pp的地址是:", &pp) // 取地址
	var ip *int                   // 指针类型
	ip = &pp                      // 赋值
	fmt.Println("指针ip的值是:", ip)
	fmt.Println("指针ip指向的变量值是:", *ip)
	fmt.Println("--------------------------------")
	fmt.Println()
	fmt.Println()

	// 七, 结构体
	type Books struct { // 定义结构体, 局部
		title   string
		book_id int
		price   float32
	}

	book1 := Books{title: "Go语言设计", book_id: 001, price: 123.50} // 构造结构体,写法1
	book2 := Books{"C++从入门到放弃", 002, 10.50}                      // 构造结构体,写法2
	book3 := Books{title: "精通Python语言", price: 23.50}            // // 构造结构体, 忽略部分字段
	fmt.Println(book1)
	fmt.Println(book2)
	fmt.Println(book3)

	shop := []Books{book1, book2, book3} // 结构体数组
	fmt.Printf(shop[0].title)            // 访问结构体数组的第 0 个元素的 title 属性
	fmt.Println("--------------------------------")
	fmt.Println()
	fmt.Println()

	// 八, 类型转换
	xxx := 1
	yyy := 2
	fmt.Println("xxx:", xxx, "yyy:", yyy)
	fmt.Println("整形除法:", xxx/yyy)                      // 整形除法
	fmt.Println("类型转换后除法:", float32(xxx)/float32(yyy)) // 浮点除法
	fmt.Println("--------------------------------")
	fmt.Println()
	fmt.Println()

	//九, 函数接口和多态
	var myPhone Phone         // 定义 phone 类型的变量
	// # new 关键字一般用于实例化对象, 只能传递一个参数, 为指定的类型分配内存，并初始化为零值, 返回指针
	myPhone = new(ApplePhone) // 实例化为 ApplePhone
	myPhone.call()            // 调用 call 方法

	myPhone2 := new(MiPhone) // 定义变量, 实例化为 MiPhone
	myPhone2.call()          // 调用 call 方法
	fmt.Println("--------------------------------")
	fmt.Println()
	fmt.Println()

	// 十, 并行
	// Go 语言支持并发，我们只需要通过 go 关键字来开启 goroutine 即可。
	// 需要 import "time"
	say("hello") // 顺序执行
	say("world") // 顺序执行
	fmt.Println()

	// 使用并行,如果主函数结束了,并行也会被杀死,所以需要sleep等待才能看出效果
	go say("hi")  // 协程并行执行函数
	say("golang") // 主程序顺序执行

}

//九, 函数接口和多态
type Phone interface { // 定义接口抽象接口
	// 接口类 phone
	call() // 函数名 call, 无参数, 无返回类型
}
type ApplePhone struct { // 定义 ApplePhone 结构
}

func (i ApplePhone) call() { // 实现 call 方法
	fmt.Println("I'm ApplePhone, calling you")
}

type MiPhone struct { // 定义 MiPhone 结构
}

func (i MiPhone) call() { // 实现 call 方法
	fmt.Println("I'm MiPhone, calling you")
}

func swap(x string, y string) (string, string) { // 定义函数, 多返回值
	return y, x
}

func say(str string) {
	for i := 0; i < 5; i++ { // 循环 5 次
		time.Sleep(100 * time.Millisecond) // 延时 100 毫秒
		fmt.Println(str)                   // 输出传入的参数
	}
}
```

### 异常处理
```go
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
```


### 面向对象编程
#### 1. OOP基础

```go
package main

import "fmt"

// 定义一个名为Profile 的结构体
type Profile struct { // 在 Go 语言中没有没有 class 类的概念，只有 struct 结构体的概念，因此也没有继承
	name   string // 定义结构体的属性
	age    int
	gender string
	mother *Profile // 指针
	father *Profile // 指针
}

// 定义一个与 Profile 的绑定的方法, 将 fmt_profile 方法与 Profile 的实例绑定。
// 我们把 Profile 称为方法的接收者, 而 person 表示实例本身，它相当于 Python 中的 self，在方法内可以使用 person.属性名 的方法来访问实例属性
func (person Profile) FmtProfile() { // 以值做为方法接收者
	fmt.Printf("名字：%s\n", person.name)
	fmt.Printf("年龄：%d\n", person.age)
	fmt.Printf("性别：%s\n", person.gender)
}

// 重点在于这个星号: *
func (person *Profile) increase_age() { // 以指针做为方法接收者
	person.age += 1
}

func main() {
	// 实例化
	myself := Profile{name: "小明", age: 24, gender: "male"}
	// 调用函数
	myself.FmtProfile()
	fmt.Printf("当前年龄：%d\n", myself.age)
	myself.increase_age()
	fmt.Printf("明年年龄：%d\n", myself.age)
}
```


#### 2. 实现 "继承" 

> Go 语言本身并不支持继承。但我们可以使用组合的方法，实现类似继承的效果。

```go
package main

import "fmt"

type company struct {
    companyName string
    companyAddr string
}

type staff struct {
    name string
    age int
    gender string
    position string
    company
}

func main()  {
    myCom := company{
        companyName: "Tencent",
        companyAddr: "深圳市南山区",
    }
    staffInfo := staff{
        name:     "小明",
        age:      28,
        gender:   "男",
        position: "云计算开发工程师",
        company: myCom,
    }

    fmt.Printf("%s 在 %s 工作\n", staffInfo.name, staffInfo.companyName)
    fmt.Printf("%s 在 %s 工作\n", staffInfo.name, staffInfo.company.companyName)
}
```

#### 3. 接口和多态

```go
package main

import (
	"fmt"
)

type Phone interface { // 定义接口抽象接口
	// 接口类 phone
	call() // 函数名 call, 无参数, 无返回类型
}
type ApplePhone struct { // 定义 ApplePhone 结构
}

func (i ApplePhone) call() { // 实现 call 方法
	fmt.Println("I'm ApplePhone, calling you")
}

type MiPhone struct { // 定义 MiPhone 结构
}

func (i MiPhone) call() { // 实现 call 方法
	fmt.Println("I'm MiPhone, calling you")
}

func main() {
	var myPhone Phone         // 定义 phone 类型的变量
	myPhone = new(ApplePhone) // 实例化为 ApplePhone
	myPhone.call()            // 调用 call 方法

	myPhone2 := new(MiPhone) // 定义变量, 实例化为 MiPhone
	myPhone2.call()          // 调用 call 方法
}

```
#### 4. 对象的访问权限

> 在 Go 语言中，函数名的首字母大小写非常重要，它被来实现控制对方法的访问权限。
>
> 当方法的首字母为大写时，这个方法对于所有包都是Public，其他包可以随意调用
> 当方法的首字母为小写时，这个方法是Private，其他包是无法访问的。





### 协程
#### 1. 初步使用
```go
package main

import (
	"fmt"
	"time"
)

func mytest() {
	fmt.Println("hello, go")
}

func main() {
	go mytest()
	fmt.Println("hello, world")
	// 初学 Go 协程可以使用 time.Sleep 来使 main 阻塞，
	// 使其他协程能够有机会运行完全，但你要注意的是，这并不是推荐的方式
	// 并发程序需要结合信道（channel）来实现控制
	time.Sleep(time.Second) // 添加延时函数, 阻塞进程
}
```

#### 2. 多个协程
```go
package main

import (
	"fmt"
	"time"
)

func mygo(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("In goroutine %s\n", name)
		// 为了避免第一个协程执行过快，观察不到并发的效果，加个休眠
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	go mygo("协程1号") // 第一个协程
	go mygo("协程2号") // 第二个协程
	time.Sleep(time.Second)
}
```

### 并行信道

> 信道，就是一个管道，连接多个goroutine程序 ，它是一种队列式的数据结构，遵循先入先出的规则。

#### 示例1
```go
package main

import "fmt"

func main() {
	var result int // 声明一个整型变量result
	ch := make(chan int) // 创建一个整型通道ch

	go func() { // 在一个新的Go协程中执行匿名函数
		result = <-ch // 从通道ch中接收数据，并将其赋值给result变量
	}()

	ch <- 1 // 向通道ch发送数据1
	fmt.Println(result) // 打印result的值
}
```

#### 示例2
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// var result int
	ch := make(chan int) // 创建一个整型通道ch

	go func() { // 在一个新的Go协程中执行匿名函数
		fmt.Println(<-ch) // 从通道ch中接收数据并打印
	}()

	ch <- 1 // 向通道ch发送数据1
	time.Sleep(10e9) // 睡眠10秒，以确保协程有足够的时间执行
}
```

#### 示例3
```go
// 发送者堵塞
package main

import "fmt"

func main() {
	var result int // 声明一个整型变量result
	ch := make(chan int) // 创建一个整型通道ch

	ch <- 1 // 向通道ch发送数据1，发送者会被阻塞，因为通道没有接收者

	go func() { // 在一个新的Go协程中执行匿名函数
		result = <-ch // 从通道ch中接收数据，并将其赋值给result变量
	}()

	fmt.Println(result) // 打印result的值，由于通道接收操作在新的协程中执行，此时result还未被赋值，因此输出结果为0
}
/* error:
fatal error: all goroutines are asleep - deadlock!
 */
/*
对于同一个通道，发送操作（协程或者函数中的），在接收者准备好之前是阻塞的。
 */
```

#### 示例4
```go
// 接收者堵塞
package main

import "fmt"

func main() {
	var result int       // 声明一个整型变量result
	ch := make(chan int) // 创建一个整型通道ch

	result = <-ch // 从通道ch接收数据，但此时通道ch中没有数据可接收，因此接收操作会被阻塞

	go func() { // 在一个新的Go协程中执行匿名函数
		ch <- 1 // 向通道ch发送数据1
	}()

	fmt.Println(result) // 由于接收操作被阻塞，此处不会执行，程序会发生死锁错误
}

/* error:
fatal error: all goroutines are asleep - deadlock!

接收操作是阻塞的（协程或函数中的），直到发送者可用：如果通道中没有数据，接收者就阻塞了。
*/

```

#### 示例5
```go
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

```

#### 示例6
```go
// 用channel保证所有goroutine能被执行完毕 
package main 
import ( 
	"fmt" 
) 
func main() { 
	loops := 999999 // 循环次数
	waitc := make(chan int) // 创建一个整型通道waitc

	for i := 0; i < loops; i++ { // 循环创建goroutine
		go func(i int) { // 在一个新的Go协程中执行匿名函数
			fmt.Printf("This is %d\n", i) // 打印当前循环的索引i
			if i == loops-1 { // 如果当前循环的索引i等于循环次数减1
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
```



#### 示例7
```go
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

```

#### 示例8
```go
// 通过range进行消费
package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		time.Sleep(1 * time.Hour) // 在一个新的Go协程中睡眠1小时，模拟长时间操作
	}()

	c := make(chan int) // 创建一个整型通道c

	go func() {
		for i := 0; i < 10; i = i + 1 { // 循环10次
			c <- i // 向通道c发送数据i
		}
		close(c) // 关闭通道c
	}()

	for i := range c { // 从通道c中循环接收数据，直到通道被关闭
		fmt.Println(i) // 打印接收到的数据
	}

	fmt.Println("Finished") // 打印"Finished"表示程序执行完毕
}

/*
0
1
2
3
4
5
6
7
8
9
Finished
*/

```





### select语句

> Go中的select和channel配合使用，通过select可以监听多个channel的I/O读写事件，当 IO操作发生时，触发相应的动作

#### 示例1
```go
package main

import (
	"fmt"

	"github.com/kirinlabs/HttpRequest"
)
// 同时监听不同的channel，做同一件工作，可以最快的返回结果。
func main() {
	ch1 := make(chan int) // 创建一个整型通道ch1
	ch2 := make(chan int) // 创建一个整型通道ch2
	ch3 := make(chan int) // 创建一个整型通道ch3

	go Getdata("https://www.baidu.com", ch1) // 在一个新的Go协程中执行Getdata函数，将结果发送到通道ch1
	go Getdata("https://www.baidu.com", ch2) // 在一个新的Go协程中执行Getdata函数，将结果发送到通道ch2
	go Getdata("https://www.baidu.com", ch3) // 在一个新的Go协程中执行Getdata函数，将结果发送到通道ch3

	select {
	case v := <-ch1: // 从通道ch1接收数据，并将其赋值给变量v
		fmt.Println(v) // 打印变量v的值
	case v := <-ch2: // 从通道ch2接收数据，并将其赋值给变量v
		fmt.Println(v) // 打印变量v的值
	case v := <-ch3: // 从通道ch3接收数据，并将其赋值给变量v
		fmt.Println(v) // 打印变量v的值
	}
}

func Getdata(url string, ch chan int) {
	req, err := HttpRequest.Get(url) // 发送HTTP GET请求
	if err != nil {
		// 处理错误
	} else {
		ch <- req.StatusCode() // 将HTTP响应状态码发送到通道ch
	}
}

```

#### 示例2
```go
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

```

#### 示例3
```go
// 通过select来检测channel的关闭事件
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()         // 记录当前时间
	c := make(chan interface{}) // 创建一个通道c

	go func() { // 在一个新的Go协程中执行匿名函数
		time.Sleep(2 * time.Second) // 睡眠2秒钟
		close(c)                    // 关闭通道c
	}()

	fmt.Println("Blocking on read...") // 打印信息，表示正在阻塞等待读取通道

	select {
	case <-c: // 从通道c中接收数据
		fmt.Printf("Unblocked %v later.\n", time.Since(start)) // 打印从开始阻塞到解除阻塞所经过的时间
	}
}

/*
Blocking on read...
Unblocked 2.013281s later.
*/

```

#### 示例4
```go
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

```

#### 示例5
```go
// 空 select
package main
func main() { 
    select {}
}
/*
除非有 case 执行，select 语句就会一直阻塞着。在这里，select 语句没有任何 case，因此它会一直阻塞，导致死锁。该程序会触发 panic，输出如下：

fatal error: all goroutines are asleep - deadlock!
*/
```

#### 示例6
```go

```

#### 示例7
```go
// 超时机制
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)    // 创建一个整型通道ch
	quit := make(chan bool) // 创建一个布尔型通道quit

	// 在一个新的Go协程中执行匿名函数
	go func() {
		for {
			select {
			case num := <-ch: // 从通道ch中接收数据，并将其赋值给变量num
				fmt.Println("num = ", num) // 打印变量num的值
			case <-time.After(3 * time.Second): // 经过3秒后，从time.After通道接收到数据
				fmt.Println("超时") // 打印超时信息
				quit <- true      // 向通道quit发送true值
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i                 // 向通道ch发送数据i
		time.Sleep(time.Second) // 睡眠1秒钟
	}

	<-quit              // 从通道quit接收数据，阻塞直到接收到数据
	fmt.Println("程序结束") // 打印程序结束信息
}

/*
num =  0
num =  1
num =  2
num =  3
num =  4
超时
程序结束
*/

```


### 协程、通道、Timer和Ticker的示例
```go

package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个通道，用于接收定时器触发的事件
	eventCh := make(chan string)

	// 创建一个定时器，在5秒钟后向通道发送事件
	timer := time.NewTimer(5 * time.Second)
	go func() {
		<-timer.C // 当定时器到期时，从定时器的通道中接收到数据
		eventCh <- "Timer expired"
	}()

	// 创建一个Ticker，每隔1秒向通道发送事件
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			<-ticker.C // 当Ticker触发时，从Ticker的通道中接收到数据
			eventCh <- "Ticker ticked"
		}
	}()

	// 在主协程中，等待接收事件并处理
	for {
		select {
		case event := <-eventCh:
			fmt.Println(event)
		}
	}
}
/*
在这个示例中，我们创建了一个通道 `eventCh` 来接收定时器和Ticker触发的事件。首先，我们创建一个定时器，在5秒钟后触发，并在协程中等待定时器到期，然后向通道发送事件。接下来，我们创建一个Ticker，每隔1秒触发一次，并在协程中循环等待Ticker触发，然后向通道发送事件。最后，在主协程中，我们使用 `select` 语句来等待并处理接收到的事件。
*/
```

### 综合示例
```go
package main

import "fmt"
//-  sum  函数接收一个整数切片  s  和一个整数通道  c ，用于计算切片  s  中所有元素的总和，并将结果发送到通道  c  中。
func sum(s []int, c chan int) {
        sum := 0
        for _, v := range s {
                sum += v
        }
        c <- sum // 把 sum 发送到通道 c
}


func main() {
		//- 在  main  函数中，我们定义了一个整数切片  s ，其中包含了数字 7、2、8、-9、4 和 0。
        s := []int{7, 2, 8, -9, 4, 0}
        // 然后，我们创建了一个整数通道  c 。
		c := make(chan int)
		// 使用  go  关键字并发地调用  sum  函数，将切片的前半部分和后半部分的元素分别计算总和，并将结果发送到通道  c  中。 
        go sum(s[:len(s)/2], c)
        go sum(s[len(s)/2:], c)
		// 最后，通过从通道  c  中接收两个结果，我们将它们分别赋值给变量  x  和  y 。然后，打印出  x 、 y  和它们的和  x+y 。
        x, y := <-c, <-c // 从通道 c 中接收
		// 输出结果
        fmt.Println(x, y, x+y)
}
```





## 5. 示例程序

### 控制台应用程序
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // 创建一个读取器来读取用户输入

	fmt.Println("请输入您的名字：")
	name, _ := reader.ReadString('\n') // 读取用户输入的名字，直到遇到换行符

	fmt.Println("请输入您的年龄：")
	var age int
	fmt.Scanln(&age) // 读取用户输入的年龄

	fmt.Println("您好，", name)
	fmt.Println("您的年龄是：", age)

	fmt.Println("欢迎来到控制台示例程序！")
	fmt.Println("请输入您的选择（1、2、3）：")

	for {
		fmt.Println("1. 执行操作1")
		fmt.Println("2. 执行操作2")
		fmt.Println("3. 退出程序")

		var choice int
		fmt.Scanln(&choice) // 读取用户输入的选择

		switch choice {
		case 1:
			fmt.Println("执行操作1")
		case 2:
			fmt.Println("执行操作2")
		case 3:
			fmt.Println("程序已退出")
			return
		default:
			fmt.Println("无效的选择")
		}
	}
}
```


### web应用程序
```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 定义路由处理函数
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	// 启动Web服务器，监听端口8080
	fmt.Println("Server is running on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

// 处理主页请求
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!")
}

// 处理关于页面请求
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the About Page!")
}
```

> 编译程序, 运行
> 打开浏览器, 范围 http://localhost:8080/ 即可查看到golang程序输出的欢迎信息

