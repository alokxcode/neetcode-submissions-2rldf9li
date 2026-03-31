func hasDuplicate(nums []int) bool {
    numsMap := make(map[int]int)
    for i,v := range nums {
        if _,ok := numsMap[v]; ok {
            return true
        }

        numsMap[v] = i
    }
    return false
}
