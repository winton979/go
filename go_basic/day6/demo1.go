package main

// func
import "fmt"

func main() {
	fmt.Println(getSmall(1, 3))
	fmt.Println(getMapInfo(1, 2))
	fmt.Println("---测试值传递--")
	x := 1
	y := 2
	fmt.Println(x, y)
	fmt.Println("after ")
	test1(x, y)
	fmt.Println(x, y)

	fmt.Println("---引用值传递--")
	x1 := 1
	y1 := 2
	fmt.Println(x1, y1)
	fmt.Println("after ")
	test2(&x1, &y1)
	fmt.Println(x1, y1)

	// 函数作为值
	fmt.Println("___函数作为值___")
	testFunc := func(x int) int {
		return x
	}
	fmt.Println(testFunc(1))

}

func getSmall(i int, i2 int) int {
	if i < i2 {
		return i
	} else if i == i2 {
		return 0
	} else {
		return i2
	}
}

// 多返回值
func getMapInfo(x, y int) (int, map[int]int) {
	m1 := map[int]int{x: y}
	return len(m1), m1
}

// 默认值传递
func test1(x, y int) (int, int) {
	x = 0
	y = 0
	return x, y
}

// 测试引用传递
func test2(x *int, y *int) {
	i := x // 保存地址???
	*x = *i
	*y = *i
}
