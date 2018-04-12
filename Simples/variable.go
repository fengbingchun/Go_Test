// 变量相关测试代码
package main

import "fmt"

// 全局变量声明
var c float32
var x = 11

func main() {
	// 局部变量声明
	// 指定变量类型，声明后若不赋值，使用默认值
	var available bool
	available = true
	fmt.Println(available)

	// 根据值自动判定变量类型
	var a = 1
	var b = 2.5
	fmt.Println(a, b)

	c = float32(a) + float32(b)
	fmt.Printf("c = %f\n", c)

	// 省略var, 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误,这种形式只能被用在函数体中
	valid := false
	_, d := 5, 6 // 空白标识符
	fmt.Println(valid, d)

	// 多变量声明
	var enabled, disabled = true, "111"
	fmt.Println(enabled, disabled)

	// 获取变量a的内存地址
	fmt.Println(&a)

	// Go语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑
	var x = -11
	fmt.Printf("x = %d\n", x)
}
