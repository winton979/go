package main

import (
	"fmt"
)

func main() {
	// * 与 &
	var a int = 1
	b := &a
	fmt.Println(b)

	fmt.Println(*b)

	// ++ --只能作为语句，不能作为运算符
	c := 1
	c++
	// c = c--

}
