package main

// 给包取别名,取别名后不可以调用原名
import (
	f "fmt"
)

func main() {
	f.Print("hello world")

	x := 1
	f.Println(*&x)
}
