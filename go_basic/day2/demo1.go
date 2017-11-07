package main

// 常量的初始化和枚举
// iota是每遇到const 从0开始
import (
	"fmt"
)

const a int = 1
const b = 'A'

const (
	c = a
	d = 2
	e = 3 + a
)

const (
	h, g = 1, 2
)

const (
	m1 = 'A'
	m2
	m3 = iota
	m4
	m5 = iota
	m6 = iota
)

const (
	// 常量请用大写
	// 如果b无类型，那么给b上一行的类型
	i = 1 // iota 0,和顺序有关，和出现次数无关
	j
	l = "RRRR"
	k = len(l)

	_K1 = 3
)

func main() {
	fmt.Println(a, b)
	fmt.Println(i, j, k)

	fmt.Println(m1, m2, m3, m4, m5, m6)
}
