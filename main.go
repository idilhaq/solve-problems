package main

import (
	"fmt"

	commonpath "github.com/solve-problems/common-path"
	highestdelay "github.com/solve-problems/highest-delay"
	"github.com/solve-problems/palindrome"
	swaptwo "github.com/solve-problems/swap-two"
)

func main() {
	// Palyndrome
	fmt.Println("----------Palyndrome----------")
	input := "avdca"
	fmt.Println("input:", input)
	isPaliyndrome := palindrome.Palindrome(input)
	fmt.Printf("isPalyndrome: %v", isPaliyndrome)
	fmt.Println()
	fmt.Println()

	// Common Path
	fmt.Println("----------Common Path----------")
	paths := []string{
		"/a/folder1/data/file.txt",
		"/b/folder2/../folder1/data/file.txt",
		"/a/folder3/folder1/folder1/../data/file.txt",
	}
	fmt.Println("input:", paths)
	fmt.Println("commonPath:", commonpath.CommonPath(paths))
	fmt.Println()

	// Swap Two Variables
	fmt.Println("----------Swap Two Variables----------")
	inputX := 1
	inputY := 2
	fmt.Println("input:", inputX, inputY)
	fmt.Println(swaptwo.SwapTwo(inputX, inputY))
	fmt.Println(swaptwo.SwapTwoWithGolang(inputX, inputY))
	fmt.Println(swaptwo.SwapTwoWithBitwiseXOR(inputX, inputY))
	fmt.Println()

	// Highest Delay
	fmt.Println("----------Highest Delay----------")
	inputKeyDelay := [][]int{{0, 1}, {1, 4}, {3, 5}, {7, 9}}
	fmt.Println("input:", inputKeyDelay)
	fmt.Println(highestdelay.HighestDelay(inputKeyDelay))
}
