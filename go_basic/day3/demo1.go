package main

// for 与 if switch
import (
	"fmt"
)

func main() {
	// 一般
	a := 2
	if 1 < a {
		fmt.Println(true)
	}
	// 2
	if b := 2; a == 2 { // 如果初始化语句的变量名和外部变量名相同，则隐藏外部变量名之至if结束
		fmt.Println(b)
	}

	// for 1 nomal
	for {
		a++
		if a > 10 {
			fmt.Println(a)
			break
		}
	}
	fmt.Println("---for 2----")
	// for 2
	for a < 15 {
		fmt.Println(a)
		a++
		fmt.Println(a)
	}
	// for 3
	// for 初始化语句；判断语句；执行语句
	fmt.Println("for 3")
	for i := 1; a < 20; a++ {
		fmt.Println(a, i)
	}

	// switch 1
	t := 1
	switch t {
	case 0:
		fmt.Println("t=0")
	case 1:
		fmt.Println("t=1")
	default:
		fmt.Println("none")
	}
	// switch 2
	fmt.Println("switch2")
	switch {
	case t == 1:
		fmt.Println("true")
		fallthrough
	case true:
		fmt.Println("true2")
		fallthrough // 继续检查一个case
	default:
		fmt.Println("default")
	}
	// switch 3
	switch y := 1; { // 同上，作用域为switch内
	default:
		fmt.Println("switch3", y)
	}

	// goto break continue
	// goto，调整运行位置
	//LABEL:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				fmt.Println("test break")
				// break LABEL // 只用break只能跳出第二层for，无限打印。如果加了标签，就能结束最外层的标签
				goto LABEL2
			}

		}
	}
LABEL2:
	fmt.Println("haha调过来了")

}
