package main

import (
	"fmt"
)

func main() {
	// 短变量 var
	//intList := [5]int{1, 2, 3, 4, 5}
	//updateArray(intList)

	//var 关键字初始化
	//var slice []int
	//slice = append(slice, 1)
	//slice = append(slice, 2)
	//slice = append(slice, 3)
	//slice = append(slice, 4)
	//slice = append(slice, 5)
	//slice = append(slice, 6)
	//fmt.Println(slice)

	//使用make关键字初始化    指针，len  cap    cap >= len
	//slice := make([]int, 5) // 长度为5，容量为5
	//slice[0] = 1
	//fmt.Println(slice, len(slice), cap(slice))

	// 使用make关键字初始化 同时设置cap长度
	//slice := make([]int, 6, 12) // 长度为5，容量为10
	//fmt.Println(slice, len(slice), cap(slice))

	// 切片化：语法u[low: high]创建对已存在数组进行操作的切片  指针  长度   容量
	// 1, 2, 3, 4, 5, 6

	//slice := []int{1, 2, 3, 4, 5, 6}
	//fmt.Println(slice, len(slice), cap(slice))
	//newSlice := slice[1:3]
	//fmt.Println(newSlice, len(newSlice), cap(newSlice))
	//newSlice[0] = 20
	//fmt.Println(slice, newSlice)

	// slice 扩容
	//var slice []int
	//slice = append(slice, 1) // 1
	//slice = append(slice, 2) // 2   20
	//slice = append(slice, 3) // 4
	//slice = append(slice, 4) // 4
	//fmt.Println(slice, len(slice), cap(slice))
	//
	//newSlice := slice[1:3] // 2 3  cap 边界一旦被触发
	//fmt.Println(newSlice, len(newSlice), cap(newSlice))
	//newSlice[0] = 20
	//fmt.Println(newSlice, slice)
	//
	//slice = append(slice, 5)
	//newSlice[0] = 7
	//fmt.Println(newSlice, slice)
}

func updateArray(arr [5]int) {
	arr[0] = 11
	fmt.Println(arr)
}
