package main

// func
import "fmt"

func main() {
	fmt.Println(getSmall(1, 3))
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
