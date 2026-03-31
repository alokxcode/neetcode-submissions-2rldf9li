type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder
	for _,v := range strs {
		sb.WriteString(fmt.Sprintf("%d#%s",len(v),v))	
	}
	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	var result []string
	i := 0
	for i < len(encoded) {
		j := i
		for encoded[j] != '#' {
			j++
		}
		length,_ := strconv.Atoi(encoded[i:j])
		result = append(result, encoded[j+1:j+1+length])
		i = j + 1 + length
	}
	return result
}
