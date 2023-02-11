package main

import (
	"fmt"
	"runtime"
)

type data struct {
	i, j int
}

func main() {
	var N = 4000000
	var structure []data
	for i := 0; i < N; i++ {
		value := int(i)
		structure = append(structure, data{value, value})
	}
	runtime.GC()
	res := structure[0]
	fmt.Println(res)
}
