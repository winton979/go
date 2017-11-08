package main

/*
切片
*/
import (
	"fmt"
)

func main() {
	// 一般初始化
	a := []int{19: 1, 7: 4}
	var b []int = []int{}
	var b1 []int // 缺省状态下，为nil
	// 通过切片初始化
	s := b[:]
	fmt.Println(a, b, s)
	fmt.Println(&b, &s, &b == &s /*b == s*/, b1 == nil)

	// 从x下标创建新切片
	c := a[4:]
	c1 := a[4:10]
	fmt.Println(c)
	fmt.Println(c1)
	// 从x下标到第一个
	c2 := a[:9]
	fmt.Println(c2)

	fmt.Println("----len cap----")
	// len和cap
	fmt.Println(len(a), cap(a), len(b), cap(b))
	d := make([]int, 10, 20)
	fmt.Println(d, len(d), cap(d))

	// copy append
	fmt.Println("----copy append----")
	var f []int
	fmt.Println(f == nil, f)
	f = append(f, 0)
	f = append(f, 2, 3, 4)
	f1 := make([]int, len(f), cap(f)*2)

	fmt.Println(f)
	fmt.Println(f1, len(f1), cap(f1))

}
