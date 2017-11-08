package main

// map
import (
	f "fmt"
)

func main() {
	// 如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
	var m map[int]string
	m = make(map[int]string)
	m[1] = "1"
	m[3] = "3"
	m[3] = "2"
	f.Println(m)

	// 循环遍历
	for k := range m {
		f.Println(k, m[k])
	}
	// 查看值是否存在
	key, ok := m[1]
	f.Println(ok, key)
	// 与if使用
	if _, ok := m[1]; ok {
		f.Println("存在")
	}

	f.Println("--delete---")
	// delete
	map1 := map[int]string{1: "a", 2: "b"}
	for k := range map1 {
		f.Println(k, map1[k])
	}
	f.Println("after delete")
	delete(map1, 1)
	for k := range map1 {
		f.Println(k, map1[k])
	}
}
