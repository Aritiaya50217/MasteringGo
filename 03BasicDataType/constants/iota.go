package main

import "fmt"

type Digit int
type Power2 int

const PI = 3.1415926
const (
	c1 = "C1"
	c2 = "C2"
	c3 = "C3"
)

func main() {
	const s1 = 123
	var v1 float32 = s1 * 12
	fmt.Println(v1)
	fmt.Println(PI)

	const (
		ZeroDigit = iota
		One
		Two
		Three
		Four
	)
	fmt.Println(ZeroDigit)
	fmt.Println(One)
	fmt.Println(Two)

	const (
		p2_0 Power2 = 1 << iota
		p2_2
	)
	fmt.Println("2^0:", p2_0)
	fmt.Println("2^2:", p2_2)
}
