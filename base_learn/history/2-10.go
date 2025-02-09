package main

import (
	"fmt"
)

func main() {
	//for i := 0; i < 10; i++ {
	//}

	//i := 0
	//for i < 10 {
	//	fmt.Println(i)
	//	i++
	//}

	//i := 0
	//for {
	//	i++
	//	fmt.Println(i)
	//	if i == 10 {
	//		break
	//	}
	//}

	//names := [3]string{"lucky", "nike", "allen"}
	//for index, name := range names {
	//	fmt.Println(index, name)
	//}
	//
	//fmt.Println("=================")
	//
	//for i := 0; i < len(names); i++ {
	//	fmt.Println(i, names[i])
	//}

	name := "allen"
	for _, v := range name {
		fmt.Println(string(v))
	}

	ages := map[string]int{
		"allen": 19,
		"nike":  29,
	}
	for k, v := range ages {
		fmt.Println(k, v)
	}

}
