// 函数相关测试代码
package main

import (
	"fmt"
	"math"
)

// 函数返回两个数的最大值，值传递
func max(num1, num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

// 函数支持返回多个值,值传递
func swap(x, y string) (string, string) {
	return y, x
}

// Go函数不支持C++中的函数重载,引用传递
func swap1(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func getSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

type Circle struct {
	radius float64
}

func (c Circle) getArea() float64 {
	// c.radius 即为Circle类型对象中的属性
	return 3.14 * c.radius * c.radius
}

func main() {
	a, b := 100, 200
	var ret = max(a, b)
	fmt.Printf("a,b最大值为: %d\n", ret)

	var x, y = "hello", "beijing"
	fmt.Printf("x: %s, y: %s\n", x, y)
	m, n := swap(x, y)
	fmt.Printf("x: %s, y: %s\n", x, y)
	fmt.Printf("m: %s, n: %s\n", m, n)

	p, q := -100, 500
	fmt.Printf("p: %d, q: %d\n", p, q)
	swap1(&p, &q)
	fmt.Printf("p: %d, q: %d\n", p, q)

	// 函数用法：函数作为值
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Printf("square root: %f\n", getSquareRoot(9))

	// 函数用法：闭包
	nextNumber := getSequence()
	// 调用nextNumber函数，i 变量自增 1 并返回
	fmt.Println("i的值: %d", nextNumber())
	fmt.Println("i的值: %d", nextNumber())
	fmt.Println("i的值: %d", nextNumber())

	// 创建新的函数nextNumber1，并查看结果
	nextNumber1 := getSequence()
	fmt.Println("i的值: %d", nextNumber1())
	fmt.Println("i的值: %d", nextNumber1())

	// 函数用法：方法
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("Area of Circle(c1) = ", c1.getArea())
}
