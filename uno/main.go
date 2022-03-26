package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i]);
	}

	for i, arg := range os.Args {
		fmt.Println(i, arg);
	}
}