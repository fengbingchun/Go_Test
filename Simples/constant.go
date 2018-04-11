// 常量相关测试代码
package main

import "fmt"
import "unsafe"

// 常量还可以用作枚举
const (
	o = "abc"
	p = len(o)
	q = unsafe.Sizeof(o)
)

// itoa: iota可以被用作枚举值,第一个iota等于0，每当iota在新的一行被使用时，它的值都会自动加1
const (
	d = iota
	e
	f
)

func main() {
	const a string = "abc" // 显示类型定义
	const b = "def"        // 隐式类型定义
	fmt.Println(a, b)

	const x, y, z = 1, 0.2, "blog"
	fmt.Println(x, y, z)

	fmt.Println(o, p, q)

	fmt.Println(d, e, f)
}
