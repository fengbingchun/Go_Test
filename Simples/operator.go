// 运算符相关测试代码
package main

import "fmt"

func main() {
	// 算术运算符
	var a, b, c = 1, 2, 3
	fmt.Printf("a + b = %d\n", a+b)
	fmt.Printf("a - c = %d\n", a-c)
	fmt.Printf("a * b = %d\n", a*b)
	c++ // 注：++和--好像仅有前缀运算符
	fmt.Printf("c++ = %d\n", c)
	b--
	fmt.Printf("--b = %d\n", b)

	// 关系运算符
	if a > b { // 注：visual studio code会默认把if (a > b)调整为if a > b
		fmt.Printf("a > b\n")
	} else {
		fmt.Printf("a <= b\n")
	}

	// 逻辑运算符
	x, y := true, false
	if x && y {
		fmt.Printf("true\n")
	} else {
		fmt.Printf("false\n")
	}

	// 位运算符: 好像仅支持整数
	a, b = 1, 3
	c = a & b
	fmt.Printf("c = %d\n", c)
	fmt.Printf("b << 2: %d\n", b<<2)

	// 赋值运算符
	b <<= 2
	fmt.Printf("b = %d\n", b)

	// 其它运算符：&、*
	fmt.Printf("b's address: %0x\n", &b)
	var ptr *int
	ptr = &a
	fmt.Printf("pointer: %0x\n", ptr)
}
