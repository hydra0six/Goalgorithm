package main

import (
	"fmt"
)

//桶排序
func Booksort() {
	var book [1001]int
	var s, t int
	fmt.Println("开始，输入元素个数：")
	fmt.Scanf("%d", &s)
	for i := 0; i < s; i++ {
		fmt.Scanf("%d", &t)
		book[t]++
	}
	fmt.Println("顺序：")
	for i := 0; i <= 1000; i++ {
		for j := 1; j <= book[i]; j++ {
			fmt.Println(i)
		}
	}
}
