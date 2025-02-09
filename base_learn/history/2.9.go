package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	if num := rand.Intn(10); num%2 != 0 {
		fmt.Println("num%2 yes", num)
		return
	}
}
