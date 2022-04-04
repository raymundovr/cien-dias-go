package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
)

func blue(img image.Image) image.Image {
	fmt.Println("Applying blue filter!")
	filteredImg := image.NewRGBA(img.Bounds())

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			ru8 := uint8(r / 0x101)
			gu8 := uint8(g / 0x101)
			bu8 := uint8(b / 0x101)
			au8 := uint8(a / 0x101)
			filteredImg.Set(x, y, color.RGBA{bu8, gu8, ru8, au8})
		}
	}

    return filteredImg
}