package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func rotate(image [][]int32) {
	n := len(image)
	for j := 0; j < n/2; j++ {
		for i := j; i < n-1-j; i++ {
			pixel := image[j][i]
			image[j][i] = image[n-1-i][j]
			image[n-1-i][j] = image[n-1-j][n-1-i]
			image[n-1-j][n-1-i] = image[i][n-1-j]
			image[i][n-1-j] = pixel
		}
	}
}

func createImage(n int) [][]int32 {

	var image [][]int32
	image = make([][]int32, n)
	for y := 0; y < n; y++ {
		image[y] = make([]int32, n)
		for x := 0; x < n; x++ {
			image[y][x] = int32(x + y)
		}
	}
	return image
}

func printImage(image [][]int32) {
	n := len(image)
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			fmt.Printf("%x", image[y][x])
		}
		fmt.Println("")
	}
}

func main() {
	var n int
	var err error
	if len(os.Args) > 1 {
		n, err = strconv.Atoi(os.Args[1])
	} else {
		err = errors.New("No arg")
	}
	if err == nil {
		fmt.Println("--- Original ---")
		img := createImage(n)
		printImage(img)
		rotate(img)
		fmt.Println("--- Rotated ---")
		printImage(img)

	} else {
		fmt.Println("Parse error: ", err)
	}

}
