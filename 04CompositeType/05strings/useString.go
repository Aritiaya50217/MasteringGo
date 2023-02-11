package main

import (
	"fmt"
	s "strings"
)

func main() {
	upper := s.ToUpper("Hello there !")
	fmt.Println(upper)

	lower := s.ToLower("HELLO")
	fmt.Println(lower)

	title := s.Title("this will be a title")
	fmt.Println(title)

	equal := s.EqualFold("Mihalis", "MIHAlis")
	fmt.Println(equal)

	fmt.Printf("Prefix: %v\n", s.HasPrefix("Mihalis", "Mi"))
	fmt.Printf("Suffix: %v\n", s.HasSuffix("Mihalis", "is"))
	fmt.Printf("Index: %v\n", s.Index("Mihalis", "ha"))
	fmt.Printf("Count: %v\n", s.Count("Mihalis", "i"))

}
