package main

import (
	"fmt"
	"runtime"
)

func init() {
	fmt.Printf("Map:%v\n", m)
	info = fmt.Sprint("OS:", runtime.GOOS, ",Arch:", runtime.GOARCH)

}

// 非全局变量，map类型，已经初始化完成
var m = map[int]string{1: "A", 2: "B", 3: "C"}

// 非全局变量，String 类型，未初始化
var info string

func main() {
	fmt.Println(info)
}
