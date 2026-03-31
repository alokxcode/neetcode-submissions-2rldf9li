func productExceptSelf(nums []int) []int {
	leftArray := make([]int,len(nums))
	rightArray := make([]int, len(nums))
	leftArray[0] = 1
	lastElement := len(nums) -1
	rightArray[lastElement] = 1
	for i := len(nums) - 2; i >= 0 ; i-- {
		rightArray[i] = rightArray[i+1] * nums[i+1]
	}

	result := make([]int, len(nums))
	previousProduct := 1
	result[0] = rightArray[0]
	for i := 1; i < len(rightArray); i++ {
		previousProduct = previousProduct * nums[i-1]
		result[i] = rightArray[i] * previousProduct
	}
	return result
}
