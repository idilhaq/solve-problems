package main

import (
	"fmt"

	commonpath "github.com/solve-problems/common-path"
	"github.com/solve-problems/palindrome"
)

func main() {
	input := "avdca"
	fmt.Println(palindrome.Palindrome(input))

	paths := []string{
		"/a/folder1/data/file.txt",
		"/b/folder2/../folder1/data/file.txt",
		"/a/folder3/folder1/folder1/../data/file.txt",
	}
	fmt.Println(commonpath.CommonPath(paths))
}
