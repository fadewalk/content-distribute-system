package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

type NetError interface {
	error
	IsLink() bool // 是否为链接错误
	IsAddr() bool
}

type MyError struct {
	Code int
	Msg  string
}

func (m MyError) Error() string {
	return "my error"
}

func (m MyError) IsLink() bool {
	return true
}

func (m MyError) IsAddr() bool {
	return true
}

type Link struct {
}

func (m Link) IsLink() bool {
	return true
}

func (m Link) IsAddr() bool {
	return false
}

func (m Link) Error() string {
	return "my error"
}

func main() {
	var x map[string]int64
	if x == nil {
		fmt.Println("yes")
	}
	x = make(map[string]int64)
	if x == nil {
		fmt.Println("1")
	}
	fmt.Println(x["3"])
}

func do() error {
	return fmt.Errorf("test error id = %d", 1)
}

func Find(id int) (int, error) {
	if id == 1 {
		return 0, &Link{}
	}

	return 1, nil
}
