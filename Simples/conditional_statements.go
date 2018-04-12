// 条件语句相关测试代码
package main

import "fmt"

func main() {
	// if
	a := 1.2
	if a > 0 {
		fmt.Printf("a > 0, a = %f\n", a)
	}

	// if else
	b := 2.5
	if b > 10 {
		fmt.Printf("b > 10\n")
	} else {
		fmt.Printf("b <= 10\n")
	}
	fmt.Printf("b = %f\n", b)

	// if 嵌套
	if a > 0 {
		if b > 0 {
			fmt.Printf("a > 0, b > 0\n")
		}
	}

	// switch
	grade := "B"
	marks := 90
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60:
		grade = "C"
	default:
		grade = "D"
	}
	fmt.Printf("grade = %s\n", grade)

	// type switch
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型: %T\n", i)
	case int:
		fmt.Printf("x 是 int 型\n")
	case float64:
		fmt.Printf("x 是 float64 型\n")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型\n")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型\n")
	default:
		fmt.Printf("未知型\n")
	}

	// select
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}

}
