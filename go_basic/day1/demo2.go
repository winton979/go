package main

import (
	"fmt"
	"math"
)

type (
	byte uint8
	rune int32
	文本   string // 不规范
)

func main() {
	var a int
	var b float32

	var c bool
	var d string
	fmt.Println(a, b, c, d)

	// 数组
	var f []int
	var f1 [1]int
	fmt.Println(f, f1)

	fmt.Println(math.MaxFloat64)

	var g 文本
	g = "文本类型"
	fmt.Println(g)

	var h int
	h = 1.1
	fmt.Print(h)
}
