package main

import "fmt"

func rerThree(x int) (int, int, int) {
	return 2 * x, x * x, -x
}

func main() {
	fmt.Println(rerThree(10))
	n1, n2, n3 := rerThree(20)
	fmt.Println(n1, n2, n3)

	n1, n2 = n2, n1
	fmt.Println(n1, n2, n3)

	x1, x2, x3 := n1*2, n1*n1, -n1
	fmt.Println(x1, x2, x3)
}
