// 循环语句相关测试代码
package main

import "fmt"

func main() {
	// for
	a, b := 0, 15
	fmt.Printf("a的值为: ")
	for a = 0; a < 10; a++ {
		fmt.Printf("  %d  ", a)
	}
	fmt.Printf("\n")

	fmt.Printf("a的值为: ")
	for a < b {
		a++
		fmt.Printf("  %d  ", a)
	}
	fmt.Printf("\n")

	// for嵌套、break
	var i, j int
	fmt.Printf("2到100之间的素数包括：")
	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break
			}
		}
		if j > (i / j) {
			fmt.Printf("  %d  ", i)
		}
	}
	fmt.Printf("\n")

	// continue
	a = 0
	fmt.Printf("a的奇数：")
	for a < 20 {
		a++
		if a%2 == 0 {
			continue
		}
		fmt.Printf("  %d  ", a)
	}
	fmt.Printf("\n")

	// goto
	a = 0
	fmt.Printf("打印a的值：")
Loop:
	for a < 10 {
		if a == 5 {
			a++
			goto Loop
		}
		fmt.Printf("  %d  ", a)
		a++
	}
	fmt.Printf("\n")
}
