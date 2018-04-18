// 切片(slice)相关测试代码
package main

import "fmt"

func main() {
	// 定义切片
	//var slice1 []int
	//slice2 := make([]int, 5)
	//slice3 := make([]int, 5, 10)

	// 切片初始化
	//slice4 := []int{1, 2, 3}
	//var arr = [5]float32{0., 1., 2., 3., 4.}
	//slice5 := arr[:]

	// len()和cap()函数
	var numbers = make([]int, 3, 5)
	printSlice(numbers)

	// 空(nil)切片
	var number2 []int
	printSlice(number2)

	// 切片截取
	numbers3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers3)

	fmt.Println("numbers3 ==", numbers3)

	// 打印子切片从索引1(包含) 到索引4(不包含)
	fmt.Println("numbers3[1:4] ==", numbers3[1:4])
	// 默认下限为
	fmt.Println("numbers3[:3] ==", numbers3[:3])
	// 默认上限为 len(s)
	fmt.Println("numbers3[4:] ==", numbers3[4:])

	// 增加切片的容量
	var numbers4 []int
	printSlice(numbers4)

	// 允许追加空切片
	numbers4 = append(numbers4, 0)
	printSlice(numbers4)

	// 向切片添加一个元素
	numbers4 = append(numbers4, 1)
	printSlice(numbers4)

	// 同时添加多个元素
	numbers4 = append(numbers4, 2, 3, 4)
	printSlice(numbers4)

	// 创建切片 numbers5是之前切片的两倍容量
	numbers5 := make([]int, len(numbers4), (cap(numbers4))*2)
	printSlice(numbers5)

	// 拷贝numbers4的内容到numbers5
	copy(numbers5, numbers4)
	printSlice(numbers5)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
