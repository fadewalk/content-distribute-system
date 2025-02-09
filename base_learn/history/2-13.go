package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func init() {
	fmt.Println("init")
}

func main() {
	a := 1
	b := 2
	num := Add(a, b)
	fmt.Println(num)
}

func Add(a, b int) int {
	num := a + b
	return num
}

func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer func() {
		fmt.Println("exec file close")
		_ = file.Close()
	}()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	fmt.Println("finish read")

	return string(data), nil
}
