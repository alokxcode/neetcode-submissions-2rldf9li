func productExceptSelf(nums []int) []int {
	leftArray := make([]int,len(nums))
	rightArray := make([]int, len(nums))
	leftArray[0] = 1
	lastElement := len(nums) -1
	rightArray[lastElement] = 1
	for i := 1; i < len(nums); i++ {
		leftArray[i] = leftArray[i-1] * nums[i-1]
	}
	for i := len(nums) - 2; i >= 0 ; i-- {
		rightArray[i] = rightArray[i+1] * nums[i+1]
	}

	result := make([]int, len(nums))
	for i,v := range leftArray {
		result[i] = v * rightArray[i]
	}
	return result
}
