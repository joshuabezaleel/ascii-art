package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"math"
	"os"
)

// Pixel struct
type Pixel struct {
	R, G, B int
}

func main() {
	//1. Loading the image and printing the width and height
	image.RegisterFormat("jpeg", "jpg", jpeg.Decode, jpeg.DecodeConfig)
	file, err := os.Open("ascii-pineapple_30.jpg")

	if err != nil {
		log.Fatalf("failed to open image: %v", err)
		os.Exit(1)
	} else {
		fmt.Println("Successfully loaded image!")
	}

	defer file.Close()

	img, _, err := image.Decode(file)

	if err != nil {
		log.Fatalf("%v", err)
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	fmt.Println("width x height= %f x %f", width, height)

	//2. Load image to 2-dimensional array of pixel data
	var pixels [][]Pixel

	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel8bit(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	//3. Converting RGB tuples to single brightness number
	var avgResult [][]float64
	for y := 0; y < height; y++ {
		var row []float64
		for x := 0; x < width; x++ {
			//Average of the R, G, and B value
			row = append(row, float64(pixels[y][x].R+pixels[y][x].G+pixels[y][x].B)/3)
		}
		avgResult = append(avgResult, row)
	}

	//4. Convert the single brightness number to ASCII characters
	chars := "`^\",:;Il!i~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"
	arrayOfChars := []rune(chars)

	//use the percentage of 255 times 20 to get the number
	remainder := float64(0)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			avgResult[y][x] = avgResult[y][x] / 255 * 64
			remainder = avgResult[y][x] - float64(int32(avgResult[y][x]))
			if remainder <= float64(0.5) {
				avgResult[y][x] = math.Floor(avgResult[y][x])
			} else {
				avgResult[y][x] = math.Ceil(avgResult[y][x])
			}
		}
	}

	var asciiChars [][]string
	for y := 0; y < height; y++ {
		var row []string
		for x := 0; x < width; x++ {
			row = append(row, string(arrayOfChars[int(avgResult[y][x])]))
		}
		asciiChars = append(asciiChars, row)
	}

	// 5. Printing the asciiChars
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			fmt.Print(asciiChars[y][x])
			fmt.Print(asciiChars[y][x])
			fmt.Print(asciiChars[y][x])
		}
		fmt.Println()
	}
}

func rgbaToPixel8bit(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257)}
}
