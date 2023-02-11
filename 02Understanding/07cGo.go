package main

import (
	"fmt"
)

func main() {
	fmt.Println("A Go statement !")
	C.callC()
	fmt.Println("Another Go statement !")
}
