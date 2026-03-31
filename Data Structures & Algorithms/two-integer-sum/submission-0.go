func twoSum(nums []int, target int) []int {
    LookupMap := make(map[int]int)
    for i,v := range nums {
		diff := target - v
		if n,ok := LookupMap[diff]; ok {
			return []int{n,i}
		}
		LookupMap[v] = i
	}
	return []int{}
}
