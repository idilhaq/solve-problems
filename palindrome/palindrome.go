package palindrome

func Palindrome(input string) bool {
	for int := range input {
		if int == len(input)-int {
			return true
		}
		if input[int] != input[len(input)-1-int] {
			return false
		}
	}

	return true
}
