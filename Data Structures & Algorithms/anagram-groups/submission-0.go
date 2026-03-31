func groupAnagrams(strs []string) [][]string {
	hashMap := make(map[string][]string)
	for _,v := range strs {
		rn := []rune(v)
		sort.Slice(rn, func(i,j int)bool {
			return rn[i] < rn[j]
		})
		sString := string(rn)
		if _, ok := hashMap[sString]; ok {
			hashMap[sString] = append(hashMap[sString],v)
		} else {
			hashMap[sString] = []string{v}
		}
	}
	var return_Slice [][]string
	for _,v := range hashMap {
		return_Slice = append(return_Slice,v)
	}

	return return_Slice
}
