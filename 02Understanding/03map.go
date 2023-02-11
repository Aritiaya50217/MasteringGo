package main

import (
	"fmt"
	"runtime"
)

func main() {
	var N = 4000000
	myMap := make(map[int]*int)
	for i := 0; i < N; i++ {
		value := int(i)
		myMap[value] = &value
	}
	runtime.GC()
	res := myMap[0]
	fmt.Println(res)
}
