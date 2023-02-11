package main

import "fmt"

const (
	c1 = "c1"
	c2 = "c2"
	c3 = "c3"
)

func main() {
	fmt.Println(c1)

	s1 := "My String"
	var s2 = "My String"
	var s3 string = "My String"

	fmt.Printf("%v\n%v\n%v",s1, s2, s3)

	const s1 = "Const1"
	const s2 = "Const2"

	fmt.Println(s1,s2) // error
}
