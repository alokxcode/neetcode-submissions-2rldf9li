func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    sMap := make(map[rune]int)
    count := 1
    for _,v := range s {
        if n,ok := sMap[v]; ok {
            sMap[v] = n+1
        } else {
            sMap[v] = count
        }

    }

    tMap := make(map[rune]int)
    for _,v := range t {
        if n,ok := tMap[v]; ok {

            tMap[v] = n+1
        } else {
            tMap[v] = count
        }

    }
    for k,v := range sMap {
        if n,ok := tMap[k]; ok {
            if n != v {
                return false
            }
        } else {
            return false
        }
    }
    return true
}
