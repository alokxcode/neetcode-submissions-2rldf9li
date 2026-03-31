func isPalindrome(s string) bool {
	str_lower := strings.ToLower(s)
	rn := []rune(str_lower)
	var sb strings.Builder
	for _,v := range rn {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			sb.WriteRune(v)
		}
	}
	str := sb.String()

	leftPointer := 0
	rightPointer := len(str) - 1

	for leftPointer < rightPointer {
		if str[leftPointer] != str[rightPointer] {
			return false
		}
		leftPointer++
		rightPointer--

	}

	return true


}
