// 数组相关测试代码
package main

import "fmt"

func getAverage(arr []int, size int) float32 {
	var avg, sum float32

	for i := 0; i < size; i++ {
		sum += float32(arr[i])
	}
	avg = sum / float32(size)

	return avg
}

func main() {
	// 一维数组
	const num = 10
	var n [num]int

	for i := 0; i < num; i++ {
		n[i] = i + 100
	}

	for j := 0; j < num; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	// 初始化数组
	const num2 = 5
	var m = [num2]float32{0., 1., 2., 3., 4.}
	for i := 0; i < num2; i++ {
		fmt.Printf("Element[%d] = %f\n", i, m[i])
	}

	// 初始化二维数组并访问
	var a = [3][4]int{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11}}
	fmt.Printf("a[2][3]: %d\n", a[2][3])

	// 向函数传递数组
	var balance = []int{1000, 2, 3, 17, 50}
	avg := getAverage(balance, 5)
	fmt.Printf("平均值：%f\n", avg)
}
