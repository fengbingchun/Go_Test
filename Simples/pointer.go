// 指针相关测试代码
package main

import "fmt"

const MAX int = 3

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func main() {
	// 指针常规操作
	var a = 20
	var ip *int

	ip = &a

	fmt.Printf("a 变量的地址是: %x\n", &a)
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	var ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr)
	if ptr == nil {
		fmt.Printf("ptr为空指针\n")
	}

	// 指针数组
	b := []int{10, 100, 200}
	var ptr2 [MAX]*int

	for i := 0; i < MAX; i++ {
		ptr2[i] = &b[i]
	}

	for i := 0; i < MAX; i++ {
		fmt.Printf("b[%d] = %d\n", i, *ptr2[i])
	}

	// 指向指针的指针
	var pptr **int
	pptr = &ip
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)

	// 向函数传递指针参数
	x, y := -11, 22
	fmt.Printf("x = %d, y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("x = %d, y = %d\n", x, y)
}
