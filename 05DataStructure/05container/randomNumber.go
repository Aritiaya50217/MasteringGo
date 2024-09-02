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
	max := 100
	total := 100
	seed := time.Now().Unix()
	arguments := os.Args

	switch len(arguments) {
	case 2:
		fmt.Println("Usage: ./randomNumbers MIN MAX TOTAL SEED")
		min, _ = strconv.Atoi(arguments[1])
		max = min + 100
	case 3:
		fmt.Println("Usage: ./randomNumbers MIN MAX TOTAL SEED")
		min, _ = strconv.Atoi(arguments[1])
		max, _ = strconv.Atoi(arguments[2])
	case 4:
		fmt.Println("Usage: ./randomNumbers MIN MAX TOTAL SEED")
		min, _ = strconv.Atoi(arguments[1])
		max, _ = strconv.Atoi(arguments[2])
		total, _ = strconv.Atoi(arguments[3])
	case 5:
		min, _ = strconv.Atoi(arguments[1])
		max, _ = strconv.Atoi(arguments[2])
		total, _ = strconv.Atoi(arguments[3])
		seed, _ = strconv.ParseInt(arguments[4], 10, 64)
	default:
		fmt.Println("Using default values!")
	}
	rand.Seed(seed)
	for i := 0; i < total; i++ {
		myrand := random(min, max)
		fmt.Print(myrand)
		fmt.Print(" ")
	}
	fmt.Println()
}
