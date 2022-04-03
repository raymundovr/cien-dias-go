package main

import (
	"fmt"
	"image"
	"os"
	_ "image/png"
	_ "image/jpeg"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Just need a single file path...")
		os.Exit(1)
	}

	filePath := args[0]
	reader, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v", filePath, err)
		os.Exit(1)
	}

	defer reader.Close()

	img, imgFormat, err := image.Decode(reader)
	
	if err != nil {
		fmt.Printf("Error reading file %s: %v", filePath, err)
		os.Exit(1)
	}

	fmt.Println("Got format", imgFormat)

	fmt.Println("Image", img.Bounds())
}