package main

import (
	"fmt"
	"os"
	"strconv"
)

func biggest(numbers []int) int {
	var max int

	for _, n := range numbers {
		if n > max {
			max = n
		}
	}
	
	return max
}

func main() {
	args := os.Args[1:]
	var numbers []int 
	
	for _, a := range args {
		val, err := strconv.Atoi(a);
		if err == nil {
			numbers = append(numbers, val)
		}
	}

	if (len(numbers) < 2) {
		fmt.Println("Need at least two numbers");
		os.Exit(1);
	}

	max := biggest(numbers)
	fmt.Println("El más grande es", max);
}