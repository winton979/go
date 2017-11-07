package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a float32 = 100.1
	fmt.Println(a)

	b := int(a)
	fmt.Println(b)

	// 作业
	var c int = 65
	d := string(c)
	fmt.Println(d) // A

	// 如果想要string
	f := strconv.Itoa(c) // int->string
	fmt.Println(f)

	h, _ := strconv.Atoi(d)
	fmt.Println(h)

}
