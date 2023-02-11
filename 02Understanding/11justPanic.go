package main

func main() {
	if len(os.Args) == 1 {
		panic.("Not enough arguments !")
	}
	fmt.Println("Thank for the argument(s) !")
}
