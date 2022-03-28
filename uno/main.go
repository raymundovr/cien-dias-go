package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	var sum int

	for _, a := range args {
		val, err := strconv.Atoi(a)
		if err == nil {
			sum += val
		} else {
			fmt.Println(a, " is not a valid int")
		}
	}
	
	fmt.Println("Sum is ", sum)
}
