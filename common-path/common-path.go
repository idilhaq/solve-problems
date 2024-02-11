package commonpath

import (
	"fmt"
	"strings"
)

/*
input
path = [
	"/a/folder1/data/file.txt",
	"/b/folder2/../folder1/data/file.txt",
	"/a/folder3/folder1/folder1/../data/file.txt",
]
.. = going up to prev folder, so it removes the prev folder on path
output
"/folder1/data/file.txt"
*/
func CommonPath(paths []string) string {
	master := make(map[string]bool)
	commonString := []string{}
	for i, path := range paths {
		// convert to array
		pathArr := strings.Split(path, "/")
		// simplify, remove the /../
		var newString []string
		for _, pathArrVal := range pathArr {
			if pathArrVal == ".." {
				newString = newString[1 : len(newString)-1]
			} else {
				newString = append(newString, pathArrVal)
			}
		}

		for _, data := range newString {
			if i == 0 {
				master[data] = false
				continue
			}

			if valMaster, ok := master[data]; ok {
				if valMaster {
					commonString = append(commonString, data)
					continue
				}
				master[data] = true
			}
		}
	}

	return fmt.Sprintf("/%s", strings.Join(commonString, "/"))
}
