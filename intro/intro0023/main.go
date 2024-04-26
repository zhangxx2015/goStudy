package main

import "fmt"

type BaseModel struct {
	Id string
}

func (this *BaseModel) Info() {
	fmt.Println(this) // 无法获取未来派生的属性
}

type Staff struct {
	Name string
	Age  int
	BaseModel
}

func main() {
	staffInfo := Staff{
		Name: "小明",
		Age:  28,
	}
	staffInfo.Id = "123"

	fmt.Println(staffInfo.Id, staffInfo.Name, staffInfo.Age)
	staffInfo.Info()
}
