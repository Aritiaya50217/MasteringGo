package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type myElement struct {
	Name    string
	Surname string
	Id      string
}

var DATA = make(map[string]myElement)

func Add(key string, data myElement) bool {
	if key == "" {
		return false
	} else if LOOKUP(key) == nil {
		DATA[key] = data
	}
	return false
}

func Delete(key string) bool {
	if LOOKUP(key) != nil {
		delete(DATA, key)
	}
	return false
}

func LOOKUP(key string) *myElement {
	_, ok := DATA[key]
	if ok {
		n := DATA[key]
		return &n
	} else {
		return nil
	}
}

func Change(key string, n myElement) bool {
	DATA[key] = n
	return true
}

func Print() {
	for k, d := range DATA {
		fmt.Printf("key: %s value : %v\n", k, d)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		tokens := strings.Fields(text)

		switch len(tokens) {
		case 0:
			continue
		case 1:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 2:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 3:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 4:
			tokens = append(tokens, "")
		}

		switch tokens[0] {
		case "PRINT":
			Print()
		case "STOP":
			return
		case "DELETE":
			if !Delete(tokens[1]) {
				fmt.Println("DElete operation failed!")
			}
		case "ADD":
			n := myElement{tokens[2], tokens[3], tokens[4]}
			if !Add(tokens[1], n) {
				fmt.Println("Add operation failed!")
			}
		case "LOOKUP":
			n := LOOKUP(tokens[1])
			if n != nil {
				fmt.Printf("%v\n", *n)
			}
		case "CHANGE":
			n := myElement{tokens[2], tokens[3], tokens[4]}
			if !Change(tokens[1], n) {
				fmt.Println("Update operation failed!")
			}
		default:
			fmt.Println("Unknow command - please try again!")
		}
	}
}
