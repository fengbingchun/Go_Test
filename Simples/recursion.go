// 递归相关测试代码
package main

import "fmt"

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	// 阶乘
	var i = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))

	// 斐波那契数列
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
	fmt.Println()
}
