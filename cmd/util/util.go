package util

import "os"

func IndexOf(target string, listOfWords []string) int {
	i := 0
	for _, w := range listOfWords {
		if target == w {
			return i
		}
		i++
	}
	return -1
}

// CheckEmpty returns true if the file with path filePath is empty.
// And false otherwise, it returns true, err if an error was returned
// when getting the file information.
func CheckEmpty(filePath string) (bool, error) {
	fi, err := os.Stat(filePath)
	if err != nil {
		return true, err
	}

	return (fi.Size() == 0), nil
}

// MapStrings
func MapStrings(arr []string, f func(string) string) []string {
	newArr := make([]string, len(arr))
	for i, s := range arr {
		newArr[i] = f(s)
	}
	return newArr
}
