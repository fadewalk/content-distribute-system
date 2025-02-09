package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				const size = 64 << 10
				buf := make([]byte, size)
				buf = buf[:runtime.Stack(buf, false)]
				fmt.Println(err, string(buf))
			}
		}()
		var names = []string{"1"}
		fmt.Println(names[0])
	}()

	time.Sleep(time.Second * 10)
}

// 确保程序初始化过程的正确性
//func init() {
//	if err := loadConfig(); err != nil {
//		panic(fmt.Sprintf("Failed to load config: %v", err))
//	}
//}

// 处理不应该发生的异常情况
//func calculateArea(width, height int) int {
//	if width < 0 || height < 0 {
//		panic("Width and height must be non-negative")
//	}
//	return width * height
//}

// 确保接口或抽象类型的实现正确性
//type Animal interface {
//	Speak() string
//}
//
//type Dog struct{}
//
//func (d *Dog) Speak() string {
//	// ...
//	return ""
//}
//
//type Cat struct{}
//
//func (c *Cat) Speak() string {
//	panic("Cat's Speak method is not implemented yet")
//}
