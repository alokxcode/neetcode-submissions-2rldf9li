func topKFrequent(nums []int, k int) []int {
	hashMap := make(map[int]int)
	for _,v := range nums {
		if n,ok := hashMap[v]; ok {
			hashMap[v] = n + 1
		} else {
			hashMap[v] = 1
		}
	}
	var intSlice [][2]int
	for i,v := range hashMap {
		intSlice = append(intSlice,[2]int{i,v})
	}

	sort.Slice(intSlice,func(i,j int) bool {
		return intSlice[i][1] > intSlice[j][1]
	})

	var return_Slice []int

	for i := range k {
		return_Slice = append(return_Slice, intSlice[i][0])
	}

	return return_Slice

}
