package main

import (
	"fmt"
	"os"
	"strings"
	"io/ioutil"
)

func main() {
	if len(os.Args[1:]) < 1 {
		fmt.Println("I just need file path(s) to work...")
		os.Exit(1)
	}

	linesCounts := make(map [string] int)

	for _, path := range os.Args[1:] {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Printf("Error reading file %s %v", path, err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			linesCounts[line]++
		}
	}

	for line, count := range linesCounts {
		if count > 1 {
			fmt.Printf("Line '%s', count %d\n", line, count)
		}
	}
}