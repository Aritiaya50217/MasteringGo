package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
func main() {
	min := 0
	max := 94
	var length int64 = 8
	arguments := os.Args
	switch len(arguments) {
	case 2:
		length, _ = strconv.ParseInt(os.Args[1], 10, 64)
	default:
		fmt.Println("Using default values !")
	}
	seed := time.Now().Unix()
	rand.Seed(seed)
	startChar := "!"
	var i int64 = 1
	for {
		myRand := random(min, max)
		newChar := string(startChar[0] + byte(myRand))
		fmt.Print(newChar)
		if i == length {
			break
		}
		i++
	}
	fmt.Println()
}
