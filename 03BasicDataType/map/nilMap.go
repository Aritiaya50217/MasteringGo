package main

import "fmt"

func main() {
	aMap := map[string]int{}
	aMap["test"] = 1
	fmt.Println(aMap)

	// nil maps
	aMap = nil
	fmt.Println(aMap)

}
