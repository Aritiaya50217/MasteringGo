package main

import "fmt"

type aStructure struct {
	person string
	height int
	weight int
}

var s1 aStructure

func main() {
	p1 := aStructure{"fmt", 12, -2}
	fmt.Println(p1)
	p1 = aStructure{weight: 12, height: -2}
	fmt.Println(p1)

	type XYZ struct {
		X int
		Y int
		Z int
	}
	var s1 XYZ
	fmt.Println(s1.Y, s1.Z)
	x1 := XYZ{23, 12, 2}
	x2 := XYZ{Z: 12, Y: 13}
	fmt.Printf("x1 : %v\nx2 : %v\n", x1, x2)

	xSlice := [4]XYZ{} //  กำหนดขนาด slice
	xSlice[2] = x1
	xSlice[0] = x2
	fmt.Println(xSlice)
}
