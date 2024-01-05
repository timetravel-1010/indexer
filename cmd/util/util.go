package util

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
