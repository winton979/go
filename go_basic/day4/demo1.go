package main

import (
	"fmt"
)

func main() {
	var a [2]int
	b := a
	// 不能对长度不同的类型不同的进行赋值，或者比较
	fmt.Println(a)
	fmt.Println(b)

	c1 := [2]int{}
	c := [5]int{4: 1} // (从0开始)索引方式给固定位置的元素赋值
	c3 := [...]string{3: "3", 2: "2"}
	fmt.Println(c1, c, c3)

	d := [...]int{1, 2, 3}
	d1 := &d
	fmt.Println(d, d1)
	//	if d1 == d {
	//		fmt.Println("d1==d")
	//	}
	fmt.Println(*d1)
	// go数组是值拷贝

	p := new([10]int) // new数组指针
	p[1] = 10
	p1 := [10]int{}
	p1[1] = 10
	fmt.Println(p, p1, *p == p1)

	// 多维数组
	q := [3][2]int{2: {1, 2}} // 最外的数组长度必须指定
	fmt.Println(q)
}
