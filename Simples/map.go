// Map(集合)相关测试代码
package main

import "fmt"

func main() {
	// 创建map
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	// map插入key-value对，各个国家对应的首都
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	// 使用key输出map值
	for country := range countryCapitalMap {
		fmt.Printf("Capital of %s is %s\n", country, countryCapitalMap[country])
	}

	// 查看元素在集合中是否存在
	captial, ok := countryCapitalMap["United States"]
	// 如果ok是true,则存在,否则不存在
	if ok {
		fmt.Println("Capital of United States is", captial)
	} else {
		fmt.Println("Capital of United States is not present")
	}

	// 创建map
	countryCapitalMap2 := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New Delhi"}
	fmt.Println("原始 map:")
	for country := range countryCapitalMap2 {
		fmt.Println("Capital of", country, "is", countryCapitalMap2[country])
	}

	// 删除元素
	delete(countryCapitalMap2, "France")
	fmt.Println("删除元素后 map:")
	for country := range countryCapitalMap2 {
		fmt.Println("Capital of", country, "is", countryCapitalMap2[country])
	}
}
