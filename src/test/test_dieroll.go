package main

import (
	"fmt"
	"dieroll"
	"os"
	"strconv"
)

func main() {
	max_str := os.Args[1]
	num_str := os.Args[2]

	max, _ := strconv.Atoi(max_str)
	num, _ := strconv.Atoi(num_str)
	
	argCount := len(os.Args[1:])

	if argCount != 2 {
		fmt.Println("Incorrect number of args.  Call this program with max, num")
		os.Exit(3)
	} // if
	
	roll := dieroll.Roll(max, num)
	fmt.Println("The roll total was:  ", roll, "\n")

} // main()
