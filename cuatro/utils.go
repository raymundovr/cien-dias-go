package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)


func saveImage(fileName string, img image.Image, imgFormat string) bool {
	fmt.Println("Saving image", fileName)
	writer, err := os.Create(fileName)
	if err != nil {
		return false
	}

	if imgFormat == "png" {
		saveErr := png.Encode(writer, img)
		if saveErr != nil {
			writer.Close()
			return false;
		}
	}

	if imgFormat == "jpeg" {
		saveErr := jpeg.Encode(writer, img, &jpeg.Options{})
		if saveErr != nil {
			writer.Close()
			return false;
		}
	}
	writer.Close()
	return true
}