func longestCommonPrefix(strs []string) string {
    var answer string
    var ch string
    pos := -1
    if len(strs) == 0{
        return answer
    }
    if len(strs) <= 1{
        return strs[0]
    }
    first := strs[0]
    for i:=0; i<len(first); i++{
    found := false

        ch = string(first[i]) 

        for j := 1; j < len(strs); j++ {

            str := strs[j]
            for k:=0; k < len(str); k++{
                if string(str[k]) == ch {
                    if pos < 0 {
                        found = true
                        pos++
                        break
                    }else if k > 0 && str[k-1] == answer[pos]{
                        found = true
                        pos++
                     break
                    }
                } 
                   
            }
                found = false
            }
        }
        if found {
                answer = fmt.Sprintf("%s%s",answer,ch)
            }     
    }
    return answer
}